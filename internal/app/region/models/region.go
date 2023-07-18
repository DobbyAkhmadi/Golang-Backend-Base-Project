package models

import (
	"gorm.io/gorm"
)

type Village struct {
	ID         int       `gorm:"type:int;primaryKey;unique" json:"id"`
	Name       string    `json:"name,omitempty"`
	DistrictID int       `gorm:"type:int"`
	District   *District `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DistrictID;references:ID"`
}

type Province struct {
	ID   int    `gorm:"type:int;primaryKey;unique" json:"id"`
	Name string `json:"name,omitempty"`
}

type District struct {
	ID        int      `gorm:"type:int;primaryKey;unique" json:"id"`
	Name      string   `json:"name,omitempty"`
	RegencyID int      `gorm:"type:int"`
	Regency   *Regency `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RegencyID;references:ID"`
}

type Regency struct {
	ID         int       `gorm:"type:int;primaryKey;unique" json:"id"`
	Name       string    `json:"name,omitempty"`
	ProvinceID int       `gorm:"type:int"`
	Province   *Province `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProvinceID;references:ID"`
}

func (v *Village) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

func (v *Province) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

func (v *District) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}

func (v *Regency) AutoMigrate(tx *gorm.DB) (err error) {
	if !tx.Migrator().HasTable(&v) {
		err = tx.AutoMigrate(&v)
		if err != nil {
			return err
		}
	}
	return
}
