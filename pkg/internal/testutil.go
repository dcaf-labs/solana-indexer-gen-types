package internal

import (
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func SetupEmbeddedDb(t *testing.T) (conn *gorm.DB, cleanup func()) {
	database := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Logger(nil))
	if err := database.Start(); err != nil {
		t.Fatal(err)
	}
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	//return db, err
	cleanup = func() {
		if err := database.Stop(); err != nil {
			t.Fatal(err)
		}
	}
	return conn, cleanup
}
