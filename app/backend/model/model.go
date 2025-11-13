package model

import (
	"time"
)

type AppConfigModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Key       string    `gorm:"uniqueIndex" json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"` // 创建时间
}

func (a *AppConfigModel) TableName() string {
	return "config"
}

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content"`                         // 待办内容
	Tags      string    `gorm:"index" json:"tags"`               // 标签，逗号分隔
	DueDate   time.Time `json:"dueDate"`                         // 截止日期
	Completed int       `gorm:"index" json:"completed"`          // 是否完成
	Level     int       `gorm:"default:0" json:"level"`          // 重要级
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"` // 创建时间
}

type CronJob struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string    `gorm:"type:varchar(256)"`
	Crontab      string    `gorm:"type:varchar(256)"`
	Func         string    `gorm:"type:varchar(1024)"`
	LastRunStart time.Time `gorm:"type:timestamp"`
	LastRunEnd   time.Time `gorm:"type:timestamp"`
	RunCount     int64
	Success      bool      // 执行是否成功
	Error        string    `gorm:"type:text"` // 错误信息
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (m *CronJob) TableName() string {
	return "cron_job" // return "schema.table"
}

type Anime struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"type:text;uniqueIndex"`
	Tags string `gorm:"type:text;"`
	//Desc   string  `gorm:"type:text"`
	Score  float64 `gorm:"type:real;"`
	Hot    int32   `gorm:"type:int;"`
	Number int32   `gorm:"type:int;"`
	// 完结、未完结、是否观看等
	State int32 `gorm:"type:int;default:0"`
	// 发布时间
	PublicDate time.Time `gorm:"colum:public_date;type:timestamp"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (m *Anime) TableName() string {
	return "anime" // return "schema.table"
}

// ApiRequest 表示一个接口请求记录
type ApiRequest struct {
	ID        uint      `gorm:"primaryKey"`
	GroupName string    `gorm:"column:group_name;uniqueIndex:idx_uni"` // 分组名称
	ApiName   string    `gorm:"column:api_name;uniqueIndex:idx_uni"`   // 接口名称
	Method    string    `gorm:"column:method;uniqueIndex:idx_uni"`     // 请求方法
	URL       string    `gorm:"column:url;uniqueIndex:idx_uni"`        // 请求 URL
	Headers   string    `gorm:"column:headers"`                        // 请求头（JSON 字符串）
	Params    string    `gorm:"column:params"`                         //
	Type      string    `gorm:"column:type"`                           //
	Body      string    `gorm:"column:body"`                           //
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;"`     // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;"`     // 更新时间（每次使用时更新）
}

func (m *ApiRequest) TableName() string {
	return "api" // return "schema.table"
}

// Sentences 一言
type Sentences struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"column:content;uniqueIndex"`
	From      string    `gorm:"column:from;"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;"` // 更新时间（每次使用时更新）
}

func (m *Sentences) TableName() string {
	return "sentences" // return "schema.table"
}
