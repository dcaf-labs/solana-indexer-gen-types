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
	PublicKey
}

func TestPublicKeyScanValue(t *testing.T) {
	// open a database
	db, cleanup := internal.SetupEmbeddedDb(t)
	defer cleanup()
	err := db.AutoMigrate(&Table{})
	assert.NoError(t, err)
	// Insert
	pubInsert := Table{PublicKey: PublicKey{solana.MustPublicKeyFromBase58("HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ")}}
	assert.Equal(t, "HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ", pubInsert.PublicKey.String())
	assert.NoError(t, db.Create(&pubInsert).Error)
	// Read
	var pubFind Table
	assert.NoError(t, db.First(&pubFind, 1).Error) // find by gorm ID
	assert.Equal(t, pubInsert.PublicKey.String(), pubFind.PublicKey.String())
	assert.Equal(t, "HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ", pubFind.PublicKey.String())
}
