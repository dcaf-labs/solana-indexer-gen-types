package publickey

import (
	"github.com/dcaf-labs/solana-indexer-gen-types/pkg/internal"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type Table struct {
	gorm.Model
	GormPublicKey
}

func TestPublicKey(t *testing.T) {
	// open a database
	db, cleanup := internal.SetupEmbeddedDb(t)
	defer cleanup()
	err := db.AutoMigrate(&Table{})
	assert.NoError(t, err)
	// Insert
	pubInsert := Table{GormPublicKey: GormPublicKey{solana.MustPublicKeyFromBase58("HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ")}}
	assert.Equal(t, "HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ", pubInsert.GormPublicKey.String())
	assert.NoError(t, db.Create(&pubInsert).Error)
	// Read
	var pubFind Table
	assert.NoError(t, db.First(&pubFind, 1).Error) // find by gorm ID
	assert.Equal(t, pubInsert.GormPublicKey.String(), pubFind.GormPublicKey.String())
	assert.Equal(t, "HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ", pubFind.GormPublicKey.String())
}
