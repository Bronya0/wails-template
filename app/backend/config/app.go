package config

import (
	"app/backend/common"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
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
	Theme    string `json:"theme"`
}

func (a *AppConfig) Start(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppConfig) GetConfig() *Config {
	configPath := a.getConfigPath()
	defaultConfig := &Config{
		Width:    common.Width,
		Height:   common.Height,
		Language: common.Language,
		Theme:    common.Theme,
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return defaultConfig
	}
	err = yaml.Unmarshal(data, &defaultConfig)
	if err != nil {
		return defaultConfig
	}
	return defaultConfig
}

func (a *AppConfig) SaveConfig(config *Config) string {
	configPath := a.getConfigPath()
	fmt.Println(configPath)

	data, err := yaml.Marshal(config)
	if err != nil {
		return err.Error()
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *AppConfig) getConfigPath() string {

	exePath, err := os.Executable()
	if err != nil {
		return common.ConfigPath
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return filepath.Join(res, common.ConfigPath)

}
