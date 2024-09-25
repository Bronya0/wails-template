package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path/filepath"
)

type AppConfig struct {
	ctx context.Context
}

type Config struct {
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Language string `json:"language"`
}

func (a *AppConfig) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppConfig) LoadConfig() error {
	config, err := a.GetConfig()
	if err != nil {
		return err
	}

	// 设置窗口大小
	runtime.WindowSetSize(a.ctx, config.Width, config.Height)

	return nil
}

func (a *AppConfig) GetConfig() (*Config, error) {
	configPath, err := a.getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果配置文件不存在，返回默认配置
			return &Config{Width: 1024, Height: 768, Language: "zh-CN"}, nil
		}
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (a *AppConfig) SaveConfig(config *Config) error {
	configPath, err := a.getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func (a *AppConfig) getConfigPath() (string, error) {
	//homeDir, err := os.UserHomeDir()
	//if err != nil {
	//	return "", err
	//}
	//获取当前路径
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	fmt.Println(filepath.Join(res, "config.json"))
	return filepath.Join(res, "config.json"), nil

}
