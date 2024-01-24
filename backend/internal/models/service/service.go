package service

import (
	encryption "github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/user"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string    `json:"name" gorm:"not null"`
	Type        string    `gorm:"not null" json:"type"`
	Description *string   `json:"description"`
	Hostname    string    `gorm:"uniqueIndex:idx_hostname_port" json:"hostname"`
	Port        int       `gorm:"type:int(5);uniqueIndex:idx_hostname_port" json:"port"`
	ApiKey      string    `json:"-"`
	Username    string    `json:"-"`
	Password    string    `json:"-"`
	UserID      int       `gorm:"not null" json:"user_id"`
	User        user.User `gorm:"constraint:OnUpdate:CASCADE;foreignKey:UserID" json:"-"`
	Icon        string    `json:"icon"`
}

type ServiceDTO struct {
	gorm.Model
	Name        string    `json:"name" gorm:"not null"`
	Type        string    `gorm:"not null" json:"type"`
	Description *string   `json:"description"`
	Hostname    string    `gorm:"uniqueIndex:idx_hostname_port" json:"hostname"`
	Port        int       `gorm:"type:int(5);uniqueIndex:idx_hostname_port" json:"port"`
	ApiKey      string    `json:"api_key"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	UserID      int       `gorm:"not null" json:"user_id"`
	User        user.User `gorm:"constraint:OnUpdate:CASCADE;foreignKey:UserID" json:"-"`
	Icon        string    `json:"icon"`
}

func (s *ServiceDTO) BeforeCreate(db *gorm.DB) (err error) {
	if err := db.Where("ID = ?", s.UserID).First(&s.User).Error; err != nil {
		return err
	}

	if len(s.Password) > 0 {
		var EncryptedPassword, err = encryption.Encrypt(s.Password, s.User.Password, s.User.Salt)
		if err != nil {
			return err
		}

		s.Password = EncryptedPassword
	}

	if len(s.ApiKey) > 0 {
		var EncryptedApiKey, err = encryption.Encrypt(s.ApiKey, s.User.Password, s.User.Salt)
		if err != nil {
			return err
		}

		s.ApiKey = EncryptedApiKey
	}

	return nil
}

func FetchService(id int, db *gorm.DB) (*Service, error) {
	var service Service

	if err := db.Where("ID = ?", id).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func FetchServiceWithUser(id int, db *gorm.DB) (*Service, error) {
	s, err := FetchService(id, db)
	if err != nil {
		return nil, err
	}

	if err := db.Preload("User").First(s, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return s, nil
}
