package userprofile

import (
	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

var userprofileDB = sqldb.NewDatabase("userprofile", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

//lint:ignore U1000 encore spec
func initService() (*Service, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: userprofileDB.Stdlib(),
	}))
	if err != nil {
		return nil, err
	}
	return &Service{db: db}, nil
}
