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
	"app/backend/utils"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/zalando/go-keyring"
)

type System struct {
}

func (obj *System) CheckUpdate() *types.Tag {
	client := resty.New()
	tag := &types.Tag{}
	resp, err := client.R().SetResult(tag).Get(common.UpdateUrl)
	if err != nil || resp.StatusCode() != 200 {
		return nil
	}
	tag.TagName = strings.TrimSpace(tag.TagName)
	return tag
}

// SysGetPwd 系统安全秘钥：获取
func (obj *System) SysGetPwd(user string) (string, string) {
	// 读取密钥
	secret, err := keyring.Get(common.AppName, user)
	if err != nil {
		return "", err.Error()
	}
	return secret, ""

}

// SysSetPwd 系统安全秘钥：设置
func (obj *System) SysSetPwd(user, pwd string) string {
	err := keyring.Set(common.AppName, user, pwd)
	if err != nil {
		return err.Error()
	}
	return ""
}

// SysDeletePwd 系统安全秘钥：删除
func (obj *System) SysDeletePwd(user string) string {
	err := keyring.Delete(common.AppName, user)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (obj *System) GetProcessInfo() string {
	// 获取内存统计信息
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// 获取构建信息
	var goVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goVersion = buildInfo.GoVersion
	} else {
		goVersion = runtime.Version()
	}

	// 格式化输出详细信息
	info := fmt.Sprintf(
		"基本信息:\n"+
			"- Go版本: %s\n"+
			"- 操作系统: %s\n"+
			"- 体系结构: %s\n"+
			"- CPU数量: %d\n"+
			"- 协程数量: %d\n"+
			"- 当前时间戳: %s\n\n"+
			"内存统计:\n"+
			"- 已分配内存: %.2f MB\n"+
			"- 总分配内存: %.2f MB\n"+
			"- 系统内存: %.2f MB\n"+
			"- 堆分配: %.2f MB\n"+
			"- 堆系统内存: %.2f MB\n"+
			"- 堆空闲: %.2f MB\n"+
			"- 堆使用中: %.2f MB\n"+
			"- 栈使用中: %.2f MB\n"+
			"- 堆对象数量: %d\n"+
			"- 内存分配次数: %d\n"+
			"- 内存释放次数: %d\n\n"+
			"垃圾回收统计:\n"+
			"- 垃圾回收运行次数: %d\n"+
			"- 上次垃圾回收时间: %s\n"+
			"- 下次垃圾回收限制: %.2f MB\n"+
			"- 垃圾回收CPU占比: %.4f%%\n"+
			"- 垃圾回收总暂停时间: %v\n",
		goVersion,                        // Go版本
		runtime.GOOS,                     // 操作系统
		runtime.GOARCH,                   // 体系结构
		runtime.NumCPU(),                 // CPU数量
		runtime.NumGoroutine(),           // 协程数量
		time.Now().Format(time.DateTime), // 当前时间戳

		float64(memStats.Alloc)/1024/1024,      // 已分配内存
		float64(memStats.TotalAlloc)/1024/1024, // 总分配内存
		float64(memStats.Sys)/1024/1024,        // 系统内存
		float64(memStats.HeapAlloc)/1024/1024,  // 堆分配
		float64(memStats.HeapSys)/1024/1024,    // 堆系统内存
		float64(memStats.HeapIdle)/1024/1024,   // 堆空闲
		float64(memStats.HeapInuse)/1024/1024,  // 堆使用中
		float64(memStats.StackInuse)/1024/1024, // 栈使用中
		memStats.HeapObjects,                   // 堆对象数量
		memStats.Mallocs,                       // 内存分配次数
		memStats.Frees,                         // 内存释放次数

		memStats.NumGC, // 垃圾回收运行次数
		time.Unix(0, int64(memStats.LastGC)).Format(time.DateTime), // 上次垃圾回收时间
		float64(memStats.NextGC)/1024/1024,                         // 下次垃圾回收限制
		memStats.GCCPUFraction*100,                                 // 垃圾回收CPU占比
		time.Duration(memStats.PauseTotalNs),                       // 垃圾回收总暂停时间
	)

	return info
}

func (obj *System) Migrate() {
	st := time.Now()
	defer func() {
		utils.Log.Infof("Migrate完成，耗时 %s", time.Since(st))
	}()
	err := DB.AutoMigrate(
		&model.Todo{},
		&model.AppConfigModel{},
		&model.CronJob{},
		&model.Anime{},
		&model.ApiRequest{},
		&model.Sentences{},
	)
	if err != nil {
		utils.Log.Error(err.Error())
	}
}

func (obj *System) Checkpoint() {
	DB.Exec("PRAGMA wal_checkpoint;")
	// 手动触发VACUUM
	DB.Exec("VACUUM;")
}

// BackupDB 按照日期备份数据库文件，仅当同名文件不存在时执行备份，最多保留 maxBackups 份
func (obj *System) BackupDB() error {
	// 获取当前时间并格式化为 YYYYMMDD
	now := time.Now().Format("20060102")
	maxBackups := 14

	// 构造备份文件夹路径
	backupDir := filepath.Join(filepath.Dir(DBFilePath), "backups")

	// 创建备份文件夹（如果不存在）
	if err := os.MkdirAll(backupDir, os.ModePerm); err != nil {
		return fmt.Errorf("创建备份目录失败: %v", err)
	}

	// 构造新的备份文件路径
	dbFileName := filepath.Base(DBFilePath)
	backupFileName := fmt.Sprintf("%s_%s", now, dbFileName)
	backupFilePath := filepath.Join(backupDir, backupFileName)

	// 如果同名文件不存在则备份
	if _, err := os.Stat(backupFilePath); err != nil {
		utils.Log.Infof("开始备份：%s", backupFilePath)
		// 执行VACUUM INTO命令，进行备份
		cmd := fmt.Sprintf("VACUUM INTO '%s'", backupFilePath)
		if err := DB.Exec(cmd).Error; err != nil {
			return err
		}
		utils.Log.Infof("备份完成")
	} else {
		utils.Log.Infof("备份文件已存，跳过备份")
	}

	// 获取所有旧的备份文件并排序
	files, err := os.ReadDir(backupDir)
	if err != nil {
		return fmt.Errorf("读取备份目录失败: %v", err)
	}

	var oldBackups []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".db" {
			oldBackups = append(oldBackups, filepath.Join(backupDir, file.Name()))
		}
	}

	// 如果超过最大保留数，删除最早的备份
	if len(oldBackups) > maxBackups {
		sort.Strings(oldBackups)
		for i := 0; i < len(oldBackups)-maxBackups; i++ {
			if err := os.Remove(oldBackups[i]); err != nil {
				return fmt.Errorf("删除旧备份文件失败: %v", err)
			}
		}
	}

	return nil
}
