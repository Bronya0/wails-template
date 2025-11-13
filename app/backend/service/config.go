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

package service

import (
	"app/backend/common"
	"app/backend/model"
	"app/backend/types"
	"context"
	"encoding/base64"
	"os"
	"sync"

	"github.com/duke-git/lancet/v2/cryptor"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm/clause"
)

type AppConfig struct {
	ctx context.Context
	mu  sync.Mutex
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (a *AppConfig) Start(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppConfig) GetConfig() *types.Config {
	var defaultConfig = &types.Config{
		Width:    common.Width,
		Height:   common.Height,
		Language: common.Language,
		Theme:    common.Theme,
		DBDir:    common.ConfigDir,
	}
	configPath := a.getConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		return defaultConfig
	}
	err = yaml.Unmarshal(data, defaultConfig)
	if err != nil {
		return defaultConfig
	}
	return defaultConfig
}

func (a *AppConfig) SaveConfig(config *types.Config) string {
	a.mu.Lock()
	defer a.mu.Unlock()
	configPath := a.getConfigPath()
	//校验
	if config.DBDir != "" {
		_, err := os.Stat(config.DBDir)
		if os.IsNotExist(err) {
			return "数据库目录不存在"
		}
	}

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
func (a *AppConfig) SaveTheme(theme string) string {
	a.mu.Lock()
	defer a.mu.Unlock()
	config := a.GetConfig()
	config.Theme = theme
	data, err := yaml.Marshal(config)
	if err != nil {
		return err.Error()
	}
	configPath := a.getConfigPath()
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *AppConfig) getConfigPath() string {
	return common.ConfigFile
}

// GetVersion returns the application version
func (a *AppConfig) GetVersion() string {
	return common.Version
}

func (a *AppConfig) GetAppName() string {
	return common.AppName
}
func (a *AppConfig) GetDesc() string {
	return common.Desc
}

func (a *AppConfig) GetValue(key string, decrypt bool) string {
	var c model.AppConfigModel
	err := DB.Where("key = ?", key).First(&c).Error
	if err != nil {
		return ""
	}
	if decrypt {
		decodeString, err := base64.StdEncoding.DecodeString(c.Value)
		if err != nil {
			return ""
		}
		cbcDecrypt := cryptor.AesGcmDecrypt(decodeString, []byte(common.AESKey))
		return string(cbcDecrypt)
	}
	return c.Value
}
func (a *AppConfig) SetValue(k string, value string, encrypt bool) string {
	if encrypt {
		value = base64.StdEncoding.EncodeToString(cryptor.AesGcmEncrypt([]byte(value), []byte(common.AESKey)))
	}
	// 更新或创建
	result := DB.Model(&model.AppConfigModel{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},              // 冲突列，要提前添加唯一约束
		DoUpdates: clause.AssignmentColumns([]string{"value"}), // 更新的列
	}).Create(&model.AppConfigModel{
		Key:   k,
		Value: value,
	})
	if result.Error != nil {
		return result.Error.Error()
	}
	return ""
}

func (a *AppConfig) DeleteValue(k string) {
	DB.Where("key = ?", k).Delete(&model.AppConfigModel{})
}
