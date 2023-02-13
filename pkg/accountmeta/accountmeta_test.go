package accountmeta

import (
	"github.com/dcaf-labs/solana-indexer-gen-types/pkg/internal"
	"github.com/dcaf-labs/solana-indexer-gen-types/pkg/publickey"
	"github.com/gagliardetto/solana-go"
	"gorm.io/gorm"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Table struct {
	gorm.Model
	AccountMeta
}

func TestAccountMeta(t *testing.T) {
	// open a database
	db, cleanup := internal.SetupEmbeddedDb(t)
	defer cleanup()
	err := db.AutoMigrate(&Table{})
	assert.NoError(t, err)
	// Insert
	insert := Table{AccountMeta: AccountMeta{
		PublicKey:  publickey.PublicKey{PublicKey: solana.MustPublicKeyFromBase58("HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ")},
		IsWritable: true,
		IsSigner:   true,
	}}
	assert.Equal(t, true, insert.AccountMeta.IsSigner)
	assert.Equal(t, true, insert.AccountMeta.IsWritable)
	assert.Equal(t, "HPbhpFTrR2qgwrWM2brB6WcwDnn7FJ83fPsVe3Deo1EJ", insert.AccountMeta.PublicKey.String())
	assert.NoError(t, db.Create(&insert).Error)
	// Read
	var find Table
	assert.NoError(t, db.First(&find, 1).Error) // find by gorm ID
	assert.Equal(t, insert.AccountMeta.IsWritable, find.AccountMeta.IsWritable)
	assert.Equal(t, insert.AccountMeta.IsSigner, find.AccountMeta.IsSigner)
}
