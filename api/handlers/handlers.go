package handlers

import (
	"inventory/database"
	"inventory/models"
)

func CreatePart(part *models.Part) error {
	db := database.DB
	return db.Create(part).Error
}

func GetPartByID(id uint64) (*models.Part, error) {
	db := database.DB
	var part models.Part
	err := db.First(&part, id).Error
	return &part, err
}

func UpdatePart(part *models.Part) error {
	db := database.DB
	return db.Save(part).Error
}

func DeletePart(part *models.Part) error {
	db := database.DB
	return db.Delete(part).Error
}

func GetPartByVersion(id uint, version int) (*models.Part, error) {
	db := database.DB

	var part models.Part
	if err := db.Where("id = ? AND version = ?", id, version).First(&part).Error; err != nil {
		return nil, err
	}

	return &part, nil
}
