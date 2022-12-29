package postgres

import (
	"project/config/keys"
	model "project/models"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	connection := viper.GetString(keys.DatabaseURL)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}

	db.AutoMigrate(&model.User{}, &model.Photo{})

	return db
}
