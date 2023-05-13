package model

import (
	"firstapp/database"

	"gorm.io/gorm"
)

type Detail struct {
	gorm.Model
	UserID     uint     `gorm:"not null" json:"-"`
	User       User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	CategoryID uint     `gorm:"not null" json:"-"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
	Cost       int      `gorm:"not null" json:"cost"`
	Date       string   `gorm:"size:255;not null" json:"date"`
}

func (detail *Detail) Save() (*Detail, error) {
	err := database.Database.Create(&detail).Error
	if err != nil {
		return &Detail{}, err
	}
	return detail, nil
}

func GetAllDetailsOfUser(userID uint) (*[]Detail, error) {
	var details []Detail
	err := database.Database.Preload("User").Preload("Category").Where("user_id=?", userID).Find(&details).Error
	if err != nil {
		return &[]Detail{}, err
	}
	return &details, nil
}

func GetDetail(id string) (Detail, error) {
	var detail Detail
	err := database.Database.Preload("User").Preload("Category").Where("id=?", id).First(&detail).Error
	if err != nil {
		return Detail{}, err
	} else {
		return detail, nil
	}
}

func DeleteDetail(id string) error {
	err := database.Database.Delete(&Detail{}, id).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
