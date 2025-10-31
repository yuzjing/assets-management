// models/asset.go

package models

import "time"

// AssetBase 包含主表和附表共有的所有字段
type AssetBase struct {
	BH     string    `gorm:"type:varchar(50);column:gdzc_BH;not null" json:"gdzc_bh"`
	LB     string    `gorm:"type:varchar(50);column:gdzc_LB;not null" json:"gdzc_lb"`
	GGXH   string    `gorm:"type:varchar(255);column:gdzc_GGXH" json:"gdzc_ggxh"`
	SCCJ   string    `gorm:"type:varchar(255);column:gdzc_SCCJ" json:"gdzc_sccj"`
	ZJFS   string    `gorm:"type:varchar(50);column:gdzc_ZJFS;not null" json:"gdzc_zjfs"`
	USERBM string    `gorm:"type:varchar(50);column:gdzc_USERBM" json:"gdzc_userbm"`
	ZT     string    `gorm:"type:varchar(50);column:gdzc_ZT;not null" json:"gdzc_zt"`
	GJXX   string    `gorm:"type:varchar(255);column:gdzc_GJXX" json:"gdzc_gjxx"`
	USER   string    `gorm:"type:varchar(50);column:gdzc_USER" json:"gdzc_user"`
	RZDATE time.Time `gorm:"column:gdzc_RZDATE;not null" json:"gdzc_rzdate"`
	LYDATE time.Time `gorm:"column:gdzc_LYDATE;not null" json:"gdzc_lydate"`
	GSNAME string    `gorm:"type:varchar(50);column:gdzc_GSNAME" json:"gdzc_gsname"`
	DW     string    `gorm:"type:varchar(50);column:gdzc_DW;not null" json:"gdzc_dw"`
	JE     int       `gorm:"column:gdzc_JE" json:"gdzc_je"`
	NX     int       `gorm:"column:gdzc_NX" json:"gdzc_nx"`
	BEIZHU string    `gorm:"type:varchar(50);column:gdzc_BEIZHU" json:"gdzc_beizhu"`
	DJDATE time.Time `gorm:"column:gdzc_DJDATE;not null" json:"gdzc_djdate"`
	DJUSER string    `gorm:"type:varchar(50);column:gdzc_DJUSER" json:"gdzc_djuser"`
}

// Asset 对应主表 gdzc_tz
type Asset struct {
	ID        uint `gorm:"primaryKey;column:gdzc_ID;autoIncrement" json:"gdzc_id"`
	AssetBase      // 嵌入基础结构体
}

func (Asset) TableName() string {
	return "gdzc_tz"
}

// AssetChangeLog 对应附表 BGgdzc_jl
type AssetChangeLog struct {
	LogID     uint      `gorm:"primaryKey;column:gdzc_ID;autoIncrement" json:"log_id"` // 主键
	AssetBase           // 嵌入基础结构体
	BGtime    time.Time `gorm:"column:BGtime;not null" json:"bg_time"`
}

func (AssetChangeLog) TableName() string {
	return "BGgdzc_jl"
}
