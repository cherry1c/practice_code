package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Shops struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 商铺 id
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"` // 商铺名称
	//[ 2] type                                           ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	Type uint64 `gorm:"column:type;type:ubigint;default:0;" json:"type"` // 商铺类型 grid_option主键
	//[ 3] phone                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Phone string `gorm:"column:phone;type:varchar;size:255;" json:"phone"` // 电话号码
	//[ 4] address                                        varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	Address string `gorm:"column:address;type:varchar;size:1024;" json:"address"` // 商铺地址
	//[ 5] business_hours                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BusinessHours string `gorm:"column:business_hours;type:varchar;size:255;" json:"business_hours"` // 营业时间
	//[ 6] url                                            varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	URL string `gorm:"column:url;type:varchar;size:1024;" json:"url"` // 图片地址
	//[ 7] is_delete                                      tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsDelete sql.NullInt64 `gorm:"column:is_delete;type:tinyint;default:0;" json:"is_delete"` // 是否删除(0:未删除, 1:已删除)
	//[ 8] status                                         tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Status sql.NullInt64 `gorm:"column:status;type:tinyint;default:0;" json:"status"` // 0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过
	//[ 9] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	//[10] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间

}

func demo01(db *gorm.DB) {
	var shopList []Shops
	result := db.Model(&Shops{}).Find(&shopList)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	for i := range shopList {
		shopList[i].Type = 6
		shopList[i].ID = 0
	}
	result = db.Create(&shopList)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
}

func localhost_() {
	dsn := "root:123456@tcp(localhost:3306)/local_life?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("open db failed")
		return
	}

	demo01(db)
}
