package basemodel

import "time"

type OrmModel struct {
	ID        uint64    `json:"id" gorm:"column:id;primary_key;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	//CreatedBy uint64     `json:"created_by" gorm:"column:created_by;default:0;"`     // 创建人
	//UpdatedBy uint64     `json:"updated_by" gorm:"column:updated_by;default:0;"`
}
