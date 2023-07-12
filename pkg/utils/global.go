package utils

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CRUD struct {
	UUIDPrimaryKey
	ModelUUID
	ModelTimestamp
}

type ModelUUID struct {
	CreatedID uuid.UUID `gorm:"type:uuid" json:"createdId,omitempty"`
	UpdatedID uuid.UUID `gorm:"type:uuid" json:"updatedId,omitempty"`
}

type ModelTimestamp struct {
	CreatedAt int64 `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (mt *ModelTimestamp) BeforeCreate(*gorm.DB) (err error) {
	mt.CreatedAt = time.Now().Unix()
	mt.UpdatedAt = time.Now().Unix()
	return
}

type IDDecimalPrimaryKey struct {
	ID uint `gorm:"autoIncrement;primaryKey;unique" json:"id"`
}

type UUIDPrimaryKey struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;unique" json:"id"`
}

func (u *UUIDPrimaryKey) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type GetGlobalResponsePaginationDto struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
	TotalRows int64       `json:"total_rows"`
}

type SetPaginationDto struct {
	PageIndex    int    `json:"page_index"`
	PageSize     int    `json:"page_size"`
	GlobalSearch string `json:"search"`
	SortBy       string `json:"sort_by"`
}
