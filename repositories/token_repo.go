package repositories

import (
	"nganterin-go/helpers"
	"nganterin-go/models"
	"time"
)

func (r *compRepository) RegisterEmailVerificationToken(data models.Users) (*string, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	delete_result := tx.Where("user_id = ? AND category = ?", data.ID, "email_verification").Delete(&models.UserTokens{})
	if delete_result.Error != nil {
		tx.Rollback()
		return nil, delete_result.Error
	}

	token, err := helpers.GenerateToken(32)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	token_data := models.UserTokens{
		UserID:    data.ID,
		Token:     token,
		Category:  "email_verification",
		ExpiredAt: time.Now().Add(time.Hour * 24 * 3),
	}

	create_result := tx.Create(&token_data)
	if create_result.Error != nil {
		tx.Rollback()
		return nil, create_result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *compRepository) RegisterPartnerEmailVerificationToken(data models.Partners) (*string, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	delete_result := tx.Where("user_id = ? AND category = ?", data.ID, "email_verification").Delete(&models.PartnerTokens{})
	if delete_result.Error != nil {
		tx.Rollback()
		return nil, delete_result.Error
	}

	token, err := helpers.GenerateToken(32)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	token_data := models.PartnerTokens{
		PartnerID: data.ID,
		Token:     token,
		Category:  "email_verification",
		ExpiredAt: time.Now().Add(time.Hour * 24 * 3),
	}

	create_result := tx.Create(&token_data)
	if create_result.Error != nil {
		tx.Rollback()
		return nil, create_result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &token, nil
}
