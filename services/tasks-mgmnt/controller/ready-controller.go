package controller

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

const (
	valid  = "valid"
	exists = "exists"
)

type ReadyController struct {
	log     *zerolog.Logger
	db      *gorm.DB
	appConf *config.AppConfig
}

func NewReadyController(
	log *zerolog.Logger,
	db *gorm.DB,
	appConf *config.Appconfig,
) *ReadyController {
	return &ReadyController{
		log:     log,
		db:      db,
		appConf: appConf,
	}
}
