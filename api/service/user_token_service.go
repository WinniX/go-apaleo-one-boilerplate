package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/winnix/go-apaleo-one-boilerplate/api/domain"
	"gorm.io/gorm"
)

type UserTokenService interface {
	GetCurrent(userId uuid.UUID) (*domain.UserToken, error)
	SaveUserToken(ut *domain.UserToken) error
}

type userTokenService struct {
	db          *gorm.DB
	accountCode string
}

func NewUserTokenService(db *gorm.DB, accountCode string) UserTokenService {
	return &userTokenService{
		db:          db,
		accountCode: accountCode,
	}
}

func (s *userTokenService) GetCurrent(userId uuid.UUID) (*domain.UserToken, error) {
	res := &domain.UserToken{}
	err := s.db.
		Where(&domain.UserToken{AccountCode: s.accountCode, UserUID: userId}).
		First(res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userTokenService) SaveUserToken(ut *domain.UserToken) error {
	err := s.db.
		Where(&domain.UserToken{AccountCode: s.accountCode, UserUID: ut.UserUID}).
		Save(&ut).Error
	if err != nil {
		return fmt.Errorf("failed saving user token: %s", err.Error())
	}

	return nil
}
