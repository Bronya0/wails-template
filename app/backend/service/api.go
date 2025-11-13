package service

import (
	"app/backend/common"
	"app/backend/utils"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Bronya0/go-toast"
	"github.com/Bronya0/go-utils/fileutil"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/pelletier/go-toml"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
)

type Api struct {
	ctx  context.Context
	mu   sync.Mutex
	icon []byte
}

func NewApi() *Api {
	return &Api{}
}
func (a *Api) SetIcon(icon []byte) {
	a.icon = icon
}

func (a *Api) Start(ctx context.Context) {
	a.ctx = ctx
}

// Proxy 发送代理请求并返回 JSON 字符串
func (a *Api) Proxy(method, _url string, headers, queryParams map[string]string, json string, form, files map[string]string, useToken bool) string {
	resp, err := a.ProxyHandler(method, _url, headers, queryParams, json, form, files, useToken)
	if err != nil {
		return err.Error()
	}
	return resp.String()
}
func (a *Api) ProxyWithInfo(method, _url string, headers, queryParams map[string]string, json string, form, files map[string]string, useToken bool) map[string]any {
	resp, err := a.ProxyHandler(method, _url, headers, queryParams, json, form, files, useToken)
	if err != nil {
		return map[string]any{
			"err": err.Error(),
		}
	}
	return map[string]any{
		"status": resp.Status(),
		"body":   resp.String(),
		"time":   resp.Time(),
		"size":   resp.Size(),
		"header": resp.Header(),
		"err":    nil,
	}
}

func (a *Api) ProxyHandler(method, _url string, headers, queryParams map[string]string, json string, form, files map[string]string, useToken bool) (*resty.Response, error) {
	// 设置请求方法
	if method == "" {
		return nil, errors.New("method is required")
	}
	if _url == "" {
		return nil, errors.New("url is required")
	}
	req := a.prepare(common.HttpClient.R(), headers, queryParams, json, form, files, useToken)
	// 发送请求
	resp, err := req.Execute(strings.ToUpper(method), _url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *Api) prepare(req *resty.Request, headers, queryParams map[string]string, json string, form, files map[string]string, useToken bool) *resty.Request {

	// 设置 Headers
	if len(headers) > 0 {
		req.SetHeaders(headers)
	}

	// 处理请求体，只处理一个
	switch {
	case len(json) > 0:
		// JSON 数据
		req.SetHeader("Content-Type", "application/json")
		req.SetBody(json)
	case len(form) > 0 && len(files) > 0:
		// 表单和文件混合 (multipart/form-data)
		req.SetFormData(form)
		req.SetFiles(files)
	case len(form) > 0:
		// 仅表单数据
		req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		req.SetFormData(form)
	case len(files) > 0:
		// 仅文件上传
		req.SetFiles(files)
	}

	if useToken {
		queryParams["token"] = common.Token
	}

	if len(queryParams) > 0 {
		// 冲突时会覆盖url里的
		req.SetQueryParams(queryParams)
	}
	return req
}

// ProxyQuery 使用 GORM 执行通用 SQLite 查询，返回 []map[string]any
func (a *Api) ProxyQuery(query string, args []any) []map[string]any {
	// 执行原生 SQL 查询
	var results []map[string]any
	err := DB.Raw(query, args...).Scan(&results).Error
	if err != nil {
		utils.Log.Error(err.Error())
		return nil
	}
	return results
}

// ProxyInsert 使用 GORM 执行批量插入，返回插入成功的行数和错误信息
// 当 lookupFields 为空时执行普通插入，否则执行带冲突处理的插入
// batchSize 参数控制每批插入的数量，默认为 1000
func (a *Api) ProxyInsert(tableName string, data []map[string]any, lookupFields, updateFields []string) (int64, string) {
	if tableName == "" {
		return 0, "表名不能为空"
	}

	if len(data) == 0 {
		return 0, "插入数据为空"
	}

	// 设置默认批次大小
	batch := 1000

	// 提取字段名
	columns := make([]string, 0)
	for key := range data[0] {
		columns = append(columns, key)
	}

	var totalAffected int64
	// 分批处理
	for i := 0; i < len(data); i += batch {
		end := i + batch
		if end > len(data) {
			end = len(data)
		}
		batchData := data[i:end]

		// 构造占位符和参数
		placeholders := make([]string, len(batchData))
		values := make([]any, 0, len(batchData)*len(columns))
		for j, row := range batchData {
			place := make([]string, len(columns))
			for k, col := range columns {
				place[k] = "?"
				values = append(values, row[col])
			}
			placeholders[j] = fmt.Sprintf("(%s)", strings.Join(place, ","))
		}

		// 构造基本 INSERT 查询
		query := fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES %s",
			tableName,
			strings.Join(columns, ","),
			strings.Join(placeholders, ","),
		)

		// 如果提供了 lookupFields，则添加冲突处理逻辑
		if len(lookupFields) > 0 {
			conflictClause := fmt.Sprintf("ON CONFLICT (%s)", strings.Join(lookupFields, ","))
			var updateClause string

			if len(updateFields) == 0 {
				// 默认更新所有非查找字段
				updateSet := make([]string, 0)
				for _, col := range columns {
					if !slices.Contains(lookupFields, col) {
						updateSet = append(updateSet, fmt.Sprintf("%s = excluded.%s", col, col))
					}
				}
				updateClause = fmt.Sprintf("DO UPDATE SET %s", strings.Join(updateSet, ","))
			} else {
				// 更新指定字段
				updateSet := make([]string, 0)
				for _, uf := range updateFields {
					if !slices.Contains(lookupFields, uf) {
						updateSet = append(updateSet, fmt.Sprintf("%s = excluded.%s", uf, uf))
					}
				}
				if len(updateSet) == 0 {
					return 0, "更新字段无效（不能仅包含查找字段）"
				}
				updateClause = fmt.Sprintf("DO UPDATE SET %s", strings.Join(updateSet, ","))
			}

			// 添加冲突处理子句
			query = fmt.Sprintf("%s %s %s", query, conflictClause, updateClause)

		}

		// 执行批量插入
		result := DB.Exec(query, values...)
		if result.Error != nil {
			return totalAffected, fmt.Sprintf("操作失败: %v", result.Error)
		}
		totalAffected += result.RowsAffected
	}

	return totalAffected, ""
}

// ProxySql 使用 GORM 执行sql操作，返回受影响的行数和错误信息
func (a *Api) ProxySql(sql string) (int64, string) {
	result := DB.Exec(sql)
	if result.Error != nil {
		return 0, fmt.Sprintf("操作失败: %v", result.Error)
	}
	return result.RowsAffected, ""
}

func (a *Api) CheckAnime() string {

	// get请求，解析返回值为json
	animeUrl := "https://www.agedm.io/update"
	resp, err := common.HttpClient.R().Get(animeUrl)
	if err != nil {
		return err.Error()
	}
	if resp.StatusCode() != 200 {
		return "请求失败"
	}
	respBody := resp.Body()

	// 使用 goquery 解析网页
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(respBody)))
	if err != nil {
		return err.Error()
	}
	type Item struct {
		Title    string `json:"title"`
		Episode  string `json:"episode"`
		Url      string `json:"url"`
		ImageUrl string `json:"image_url"`
	}
	type Res struct {
		Date  string  `json:"date"`
		Items []*Item `json:"items"`
	}
	var result []*Res
	// 遍历每个更新日期区块
	doc.Find("#recent_update_video_wrapper .video_list_box").Each(func(i int, s *goquery.Selection) {
		var items []*Item
		// 提取日期
		date := s.Find("button").Text()
		//result.WriteString(fmt.Sprintf("<h3>%s</h3>\n<ul>\n", date))
		// 提取动漫名称和集数
		s.Find(".video_item").Each(func(j int, item *goquery.Selection) {
			src, _ := item.Find(".video_item--image > img").Attr("src")
			title := item.Find(".video_item-title a").Text()
			href, _ := item.Find(".video_item-title a").Attr("href")
			episode := item.Find(".video_item--info").Text()
			//result.WriteString(fmt.Sprintf("  <li>%s | %s</li>\n", title, episode))
			items = append(items, &Item{
				Title:    title,
				Episode:  episode,
				Url:      href,
				ImageUrl: src,
			})
		})
		result = append(result, &Res{
			Date:  date,
			Items: items,
		})

	})
	bytes, err := json.Marshal(&result)
	return string(bytes)
}

func (a *Api) GetWeather(city string) string {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	res, err := common.HttpClient.R().SetHeaders(header).SetQueryParams(map[string]string{
		"token": common.Token,
		"city":  city,
	}).Get(fmt.Sprintf("https://v3.alapi.cn/api/tianqi"))

	if err != nil {
		return err.Error()
	}

	return res.String()
}

type Site struct {
	URLTemplate string `json:"urlTemplate"`
	SiteName    string `json:"siteName"`
}

// SearchResult 结构体表示搜索结果
type SearchResult struct {
	URL      string `json:"url"`
	SiteName string `json:"siteName"`
}

var bookKey = "book_sites"

func (a *Api) LoadSites() ([]Site, error) {
	ac := AppConfig{}
	sites := ac.GetValue(bookKey, false)
	if sites == "" {
		return []Site{}, nil
	}
	var _sites []Site
	_ = json.Unmarshal([]byte(sites), &_sites)
	return _sites, nil
}

// AddSite 添加新的网站
func (a *Api) AddSite(urlTemplate string, siteName string) string {
	if !strings.Contains(urlTemplate, "{%v}") {
		return "URL必须包含{%v}占位符"
	}
	sites, err := a.LoadSites()
	if err != nil {
		return err.Error()
	}
	// 添加新网站
	sites = append(sites, Site{
		URLTemplate: urlTemplate,
		SiteName:    siteName,
	})
	bs, _ := json.Marshal(sites)
	ac := AppConfig{}
	return ac.SetValue(bookKey, string(bs), false)
}

// RemoveSite 删除网站
func (a *Api) RemoveSite(index int) string {
	sites, err := a.LoadSites()
	if err != nil {
		return err.Error()
	}
	if index < 0 || index >= len(sites) {
		return "无效的索引"
	}
	// 删除指定位置的网站
	sites = append(sites[:index], sites[index+1:]...)
	ac := AppConfig{}
	bs, _ := json.Marshal(sites)
	return ac.SetValue(bookKey, string(bs), false)
}

// SearchBook 搜索书籍
func (a *Api) SearchBook(bookTitle string) ([]SearchResult, error) {
	if bookTitle == "" {
		return nil, fmt.Errorf("书名不能为空")
	}
	sites, err := a.LoadSites()
	if err != nil {
		return nil, err
	}
	var results []SearchResult
	for _, site := range sites {
		// 替换URL中的占位符
		encodedTitle := url.QueryEscape(bookTitle)
		searchURL := strings.Replace(site.URLTemplate, "{%v}", encodedTitle, -1)
		results = append(results, SearchResult{
			URL:      searchURL,
			SiteName: site.SiteName,
		})
	}
	return results, nil
}

func (a *Api) Yaml2json(yamlStr string) string {

	// 解析YAML字符串到一个any
	var yamlData any
	err := yaml.Unmarshal([]byte(yamlStr), &yamlData)
	if err != nil {
		return err.Error()
	}

	// 将解析得到的数据转换为JSON
	jsonBytes, err := json.MarshalIndent(yamlData, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(jsonBytes)
}
func (a *Api) Json2Yaml(jsonStr string) string {
	var jsonData any
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return err.Error()
	}
	Bytes, err := yaml.Marshal(jsonData)
	if err != nil {
		return err.Error()
	}
	return string(Bytes)
}
func (a *Api) Yaml2toml(yamlStr string) string {
	// 解析YAML字符串到一个any
	var yamlData any
	err := yaml.Unmarshal([]byte(yamlStr), &yamlData)
	if err != nil {
		return err.Error()
	}
	Bytes, err := toml.Marshal(yamlData)
	if err != nil {
		return err.Error()
	}
	return string(Bytes)
}
func (a *Api) Toml2yaml(yamlStr string) string {
	var Data any
	err := toml.Unmarshal([]byte(yamlStr), &Data)
	if err != nil {
		return err.Error()
	}
	Bytes, err := yaml.Marshal(Data)
	if err != nil {
		return err.Error()
	}
	return string(Bytes)
}
func (a *Api) Json2Toml(jsonStr string) string {
	var jsonData any
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return err.Error()
	}
	yamlBytes, err := toml.Marshal(jsonData)
	if err != nil {
		return err.Error()
	}
	return string(yamlBytes)
}

func (a *Api) Toml2json(yamlStr string) string {
	// 解析YAML字符串到一个any
	var yamlData any
	err := toml.Unmarshal([]byte(yamlStr), &yamlData)
	if err != nil {
		return err.Error()
	}
	jsonBytes, err := json.MarshalIndent(yamlData, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(jsonBytes)
}

// Benchmark 并发基准测试方法
func (a *Api) Benchmark(clientNum, durationSeconds, timeout int, method, _url string, headers, queryParams map[string]string, _json string, form, files map[string]string, useToken bool) string {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	if clientNum < 1 {
		return "clientNum must be positive"
	}
	if durationSeconds < 1 {
		return "durationSeconds must be positive"
	}
	method = strings.ToUpper(method)

	// ResponseTimeStats 响应时间统计
	type ResponseTimeStats struct {
		Mean int `json:"mean_ms"`
		Min  int `json:"min_ms"`
		Max  int `json:"max_ms"`
		P50  int `json:"p50_ms"`
		P90  int `json:"p90_ms"`
		P95  int `json:"p95_ms"`
		P99  int `json:"p99_ms"`
	}

	// BenchmarkResult 存储基准测试的结果
	type BenchmarkResult struct {
		QPS           float64           `json:"qps"`
		TotalRequests int               `json:"total_requests"`
		Duration      float64           `json:"duration_seconds"`
		ResponseTimes ResponseTimeStats `json:"response_times"`
		StatusCodes   map[int]int       `json:"status_codes"`
		Errors        int               `json:"errors"`
		Err           string            `json:"err"`
	}

	type OnceResult struct {
		StatusCode   int
		responseTime int
		err          error
	}

	result := BenchmarkResult{
		TotalRequests: 0,
		Duration:      float64(durationSeconds),
		StatusCodes:   make(map[int]int),
		ResponseTimes: ResponseTimeStats{},
	}

	var wg sync.WaitGroup
	var wgRecv sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(durationSeconds)*time.Second)
	defer cancel()

	ch := make(chan *OnceResult, 1000)
	responseTimes := make([]int, 0, 100*10000) // 预分配容量以优化性能
	var sum int

	wgRecv.Add(1)
	go func() {
		defer func() {
			recover()
			wgRecv.Done()
		}()
		for v := range ch {
			result.TotalRequests++
			result.StatusCodes[v.StatusCode]++
			if v.err != nil {
				result.Errors++
				if result.Err == "" {
					result.Err = v.err.Error()
				}
			}
			responseTimes = append(responseTimes, v.responseTime)
			sum += v.responseTime
		}
	}()

	// 并发客户端
	client := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTimeout(time.Duration(timeout) * time.Second). // 请求超时，防止连接挂起
		SetTransport(&http.Transport{
			MaxIdleConns:        100,              // tcp连接池，总空闲连接数（空闲连接是指已经建立但当前没有被使用的 TCP 连接，保存在连接池中以供后续请求复用）
			MaxIdleConnsPerHost: 100,              // tcp连接池，单个主机（ip+端口）的最大空闲连接数
			IdleConnTimeout:     30 * time.Second, // 空闲连接超时
			//MaxConnsPerHost:     50,               // 每个主机的最大连接数
			//DialContext: (&net.Dialer{
			//	Timeout:   2 * time.Second,  // 连接建立超时
			//	KeepAlive: 30 * time.Second, // Keep-Alive 探测间隔
			//}).DialContext,
			//TLSHandshakeTimeout: 2 * time.Second, // TLS 握手超时
			//ForceAttemptHTTP2: true, // 启用 HTTP/2
		})

	for i := 0; i < clientNum; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				recover()
				wg.Done()
			}()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					// Request对象必须是每个请求隔离的。
					req := a.prepare(client.R(), headers, queryParams, _json, form, files, useToken)
					st := time.Now()
					resp, err := req.Execute(method, _url)
					status := 0
					if resp != nil {
						status = resp.StatusCode()
					}
					ch <- &OnceResult{
						StatusCode:   status,
						responseTime: int(time.Now().Sub(st).Milliseconds()),
						err:          err,
					}

				}
			}
		}()
	}

	wg.Wait()
	// 必须关闭，否则第一个协程不会结束
	close(ch)
	wgRecv.Wait()

	// 计算统计数据
	if result.TotalRequests > 0 {
		result.QPS = float64(result.TotalRequests) / result.Duration

		result.ResponseTimes.Mean = sum / result.TotalRequests
		// 计算 Min, Max, P50, P90, P95, P99
		sort.Ints(responseTimes)
		result.ResponseTimes.Min = responseTimes[0]
		result.ResponseTimes.Max = responseTimes[len(responseTimes)-1]
		result.ResponseTimes.P50 = percentile(responseTimes, 0.50)
		result.ResponseTimes.P90 = percentile(responseTimes, 0.90)
		result.ResponseTimes.P95 = percentile(responseTimes, 0.95)
		result.ResponseTimes.P99 = percentile(responseTimes, 0.99)

	}

	// 返回 JSON 结果
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf("failed to marshal result: %v", err)
	}
	return string(jsonResult)
}

// percentile 计算给定百分位数的响应时间
func percentile(data []int, p float64) int {
	if len(data) == 0 {
		return 0
	}
	index := p * float64(len(data)-1)
	i := int(index)
	frac := index - float64(i)
	if i+1 < len(data) {
		return int(float64(data[i])*(1-frac) + float64(data[i+1])*frac)
	}
	return data[i]
}

func (a *Api) Notify(title, content string) {
	_ = toast.Push(content,
		toast.WithTitle(title),
		toast.WithAppID("主人~"),
		toast.WithAudio(toast.Reminder),
		//toast.WithLongDuration(),
		toast.WithIconRaw(a.icon),
	)
}

const OtherCfKey = "other_config"

func (a *Api) Hash(path string, alg string) string {
	hash, err := fileutil.HashFile(path, alg)
	if err != nil {
		return err.Error()
	}
	return hash
}

func (a *AppConfig) OpenFileDialog(options runtime.OpenDialogOptions) (string, error) {
	return runtime.OpenFileDialog(a.ctx, options)
}
