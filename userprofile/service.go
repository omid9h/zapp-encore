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

// initService initializes the site service.
// It is automatically called by Encore on service startup.
func initService() (*Service, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: userprofileDB.Stdlib(),
	}))
	if err != nil {
		return nil, err
	}
	return &Service{db: db}, nil
}
