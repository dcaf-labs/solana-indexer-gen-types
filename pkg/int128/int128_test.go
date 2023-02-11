package int128

import (
	"github.com/dcaf-labs/solana-indexer-gen-types/pkg/internal"
	bin "github.com/gagliardetto/binary"
	"gorm.io/gorm"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Table struct {
	gorm.Model
	Int128
}

func TestInt128(t *testing.T) {
	// open a database
	db, cleanup := internal.SetupEmbeddedDb(t)
	defer cleanup()
	err := db.AutoMigrate(&Table{})
	assert.NoError(t, err)
	// Insert
	val := Int128{Int128: bin.Int128(bin.Uint128{
		Lo: 10,
	})}
	insert := Table{Int128: val}
	assert.Equal(t, "10", insert.Int128.String())
	assert.NoError(t, db.Create(&insert).Error)
	// Read
	var find Table
	assert.NoError(t, db.First(&find, 1).Error) // find by gorm ID
	assert.Equal(t, insert.Int128.String(), find.Int128.String())
	assert.Equal(t, "10", find.Int128.String())
}
