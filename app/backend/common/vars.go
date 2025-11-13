/*
 * Copyright 2025 Bronya0 <tangssst@163.com>.
 * Author Github: https://github.com/Bronya0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	// Version 会在编译时注入 -ldflags="-X 'app/backend/common.Version=${{ github.event.release.tag_name }}'"
	Version = ""
)

const (
	AppName     = "wails"
	Desc        = ""
	Width       = 1600
	Height      = 862
	Theme       = "dark"
	configDir   = ".wails"
	configFile  = "config.yaml"
	LocalDBFile = "wails.db"
	LogPath     = "wails.log"
	Language    = "zh-CN"
	PingUrl     = "https://ysboke.cn/api/kingTool/ping"
)

const (
	Token  = "ewewewewe"
	SK     = "wailsKey"
	AESKey = "wailsAESKeyTT"
)

var (
	HomeDir    = getHomeDir()
	ConfigDir  = filepath.Join(HomeDir, configDir)
	WordDir    = getWordDir()
	ConfigFile = InitConfigDir()
	HttpClient = initHttpClient()
)

var (
	Project   = "Bronya0/ES-King"
	UpdateUrl = fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", Project)
)

var (
	BaseHeader = map[string]string{
		"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"cache-control":   "no-cache",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
	}
)

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir
}

func getWordDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

func InitConfigDir() string {
	_, err := os.Stat(ConfigDir)
	if os.IsNotExist(err) {
		_ = os.Mkdir(ConfigDir, os.ModePerm)
	}
	return filepath.Join(ConfigDir, configFile)
}

func initHttpClient() *resty.Client {
	// 创建一个新的resty客户端实例
	return resty.New().
		// 设置HTTP传输配置
		SetTransport(&http.Transport{
			MaxIdleConns:        10,               // 最大空闲连接数
			MaxIdleConnsPerHost: 10,               // 每个主机的最大空闲连接数
			MaxConnsPerHost:     20,               // 每个主机的最大连接数
			IdleConnTimeout:     60 * time.Second, // 空闲连接超时
			DisableKeepAlives:   false,            // 启用 Keep-Alive
		}).
		// 设置客户端超时时间
		SetTimeout(60 * time.Second).
		// 设置TLS客户端配置，跳过TLS验证
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
