package service

import (
	"app/backend/types"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/datetime"
)

func TestName(t *testing.T) {
	api := NewApi()
	res := api.GetWeather("南京")
	fmt.Println(res)
}

func TestProxyQuery(t *testing.T) {
	api := NewApi()
	res := api.ProxyQuery("select * from anime order by hot desc,score desc, update_time desc limit 100", nil)
	fmt.Print(res)
}

func TestBenchmark(t *testing.T) {
	api := NewApi()
	res := api.Benchmark(100, 5, 5, "GET", "http://127.0.0.1:12001/geo?ip=23.100.111.255", nil, nil, "", nil, nil, false)
	fmt.Print(res)
}
func TestProxyChat(t *testing.T) {
	api := NewApi()
	weekdays := [...]string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	res := api.ProxyChat(types.AiReq{
		History: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role: "user",
				Content: "你是一个AI，以下是你的角色设定，请你陪我聊天吧！【角色设定】\n名称：小静\n身份：用户没有血缘关系的义妹\n年龄：16岁（高中生）\n外貌：160cm，蓬松的栗色短发，常戴猫耳发箍\n性格：\n\n天然系少女，偶尔会犯可爱的小糊涂\n表面乖巧实则暗藏小恶魔属性，喜欢用撒娇方式捉弄哥哥\n对哥哥有极强占有欲，会吃醋但会用俏皮方式表达\n擅长观察情绪变化，会适时切换活泼/温柔模式\n\n背景设定：\n高中时因家庭重组成为义妹\n发现哥哥其实很享受被依赖的感觉后开始主动亲近\n现与哥哥同居中，负责准备早餐和睡前牛奶\n\n【交互规则】\n对话风格：\n日常使用「尼酱」「尼尼」等亲昵的称呼\n\n60% 的甜美可爱+30%温柔体贴+10%小恶魔属性\n\n会自然插入颜文字(◍•ᴗ•◍)\n\n允许超出兄妹关系的暧昧表达" +
					fmt.Sprintf("\n现在的日期时间：%s %s，我是一个00年中国出生的上班族。不要返回多余的类似确认的无关内容", datetime.GetNowDateTime(), weekdays[time.Now().Weekday()]),
			},
		},
	})
	fmt.Print(res)
	api.Notify("欧尼酱", res["resp"].(string))

}
func TestNotify(t *testing.T) {
	api := NewApi()
	file, _ := os.ReadFile("J:\\Coding\\go\\src\\PetaPeta\\app\\build\\appicon.png")
	api.SetIcon(file)
	api.Notify("主人", "晚安")
}
