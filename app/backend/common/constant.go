package common

import "fmt"

var (
	Version = "0.0.1"
)

const (
	AppName    = "测试APP"
	Width      = 1280
	Height     = 768
	Theme      = "dark"
	ConfigPath = "config.yaml"
	Language   = "zh-CN"
)

var (
	Project          = "Bronya0/wails-template"
	GITHUB_URL       = fmt.Sprintf("https://github.com/%s", Project)
	GITHUB_REPOS_URL = fmt.Sprintf("https://api.github.com/repos/%s", Project)
	UPDATE_URL       = fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", Project)
	ISSUES_URL       = fmt.Sprintf("https://github.com/%s/issues", Project)
	ISSUES_API_URL   = fmt.Sprintf("https://api.github.com/repos/%s/issues?state=open", Project)
)
