package service

import (
	"fmt"
	"net"
	"regexp"

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

	hostname := net.ParseIP(s.Hostname)
	validDomain, _ := regexp.MatchString(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`, s.Hostname)
	if (!validDomain) {
		return fmt.Errorf("invalid hostname: %s", s.Hostname)
	}
	if hostname == nil && !validDomain {
			return fmt.Errorf("invalid hostname: %s", s.Hostname)
	}

	if s.Port < 0 || s.Port > 65535 {
			return fmt.Errorf("invalid port: %d", s.Port)
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
	var service *Service

	if err := db.Where("ID = ?", id).First(&service).Error; err != nil {
		return nil, err
	}

	if err := db.Preload("User").First(service, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if service.ApiKey != "" {
		var DecryptedApiKey, err = encryption.Decrypt(service.ApiKey, service.User.Password, service.User.Salt)
		if err != nil {
			return nil, err
		}

		service.ApiKey = DecryptedApiKey
	}

	if service.Password != "" {
		var DecryptedPassword, err = encryption.Decrypt(service.Password, service.User.Password, service.User.Salt)
		if err != nil {
			return nil, err
		}

		service.Password = DecryptedPassword
	}

	return service, nil
}
