package internal

import (
	"fmt"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/phayes/freeport"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
)

func getConnectionString(host, user, password, dbName string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName)
}

func SetupEmbeddedDb(t *testing.T) (conn *gorm.DB, cleanup func()) {
	dbName := "postgres"
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}
	user, password := "postgres", "postgres"
	database := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Username(user).
		Password(password).
		Port(uint32(port)).
		Version(embeddedpostgres.V15).
		RuntimePath(fmt.Sprintf("./.tmp/%s/extracted", dbName)).
		DataPath(fmt.Sprintf("./.tmp/%s/extracted/data", dbName)).
		BinariesPath(fmt.Sprintf("./.tmp/%s/extracted", dbName)).
		Logger(nil))
	if err := database.Start(); err != nil {
		t.Fatal(err)
	}
	connString := getConnectionString("localhost", "postgres", "postgres", dbName, port)
	conn, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	cleanup = func() {
		if err := database.Stop(); err != nil {
			t.Fatal(err)
		}
		if err := os.RemoveAll(fmt.Sprintf("./.tmp/%s/", dbName)); err != nil {

		}
	}
	return conn, cleanup
}
