package laracom_user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/logger"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	logger.Info("Received CreateUser request:BeforeCreate",scope.Value)

	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}