package unixtimeseconds

import (
	"github.com/dcaf-labs/solana-indexer-gen-types/pkg/internal"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type Table struct {
	gorm.Model
	solana.UnixTimeSeconds `gorm:"type:bigint"`
}

func TestInt128ScanValue(t *testing.T) {
	// open a database
	db, cleanup := internal.SetupEmbeddedDb(t)
	defer cleanup()
	err := db.AutoMigrate(&Table{})
	assert.NoError(t, err)
	// Insert
	val := solana.UnixTimeSeconds(10)
	insert := Table{UnixTimeSeconds: val}
	assert.Equal(t, int64(10), insert.UnixTimeSeconds.Time().Unix())
	assert.NoError(t, db.Create(&insert).Error)
	// Read
	find := Table{}
	assert.NoError(t, db.First(&find, 1).Error) // find by gorm ID
	assert.Equal(t, insert.UnixTimeSeconds.String(), find.UnixTimeSeconds.String())
	assert.Equal(t, int64(10), find.UnixTimeSeconds.Time().Unix())
}
