package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Model struct {
	ID string `gorm:"type:varchar(32);primary_key"`
	Created  uint   `gorm:"type:int(10);autoCreateTime;not null"`
	Updated  uint   `gorm:"type:int(10);autoUpdateTime;default:0"`
	Deleted  uint   `gorm:"type:int(10);default:0"`
	Handlers string `gorm:"type:varchar(32);default:''"`
	State    uint8  `gorm:"type:tinyint(1);default:1"` // 1 正常 2 被删除
}


func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
	fmt.Printf("auto generate uuid")
	return nil
}

// TODO: 更新时间
// func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
// 	m.Updated = uint(time.Now().Unix())
// 	return nil
// }


