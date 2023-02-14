package uint128

import (
	"github.com/dcaf-labs/solana-indexer-gen-types/pkg/internal"
	bin "github.com/gagliardetto/binary"
	"gorm.io/gorm"
	"math/big"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Table struct {
	gorm.Model
	Uint128 `gorm:"type:numeric"`
}

func TestInt128ScanValue(t *testing.T) {
	// open a database
	db, cleanup := internal.SetupEmbeddedDb(t)
	defer cleanup()
	err := db.AutoMigrate(&Table{})
	assert.NoError(t, err)
	// Insert
	val := Uint128{Uint128: bin.Uint128{Lo: 10}}
	insert := Table{Uint128: val}
	assert.Equal(t, "10", insert.Uint128.String())
	assert.NoError(t, db.Create(&insert).Error)
	// Read
	find := Table{}
	assert.NoError(t, db.First(&find, 1).Error) // find by gorm ID
	assert.Equal(t, insert.Uint128.String(), find.Uint128.String())
	assert.Equal(t, "10", find.Uint128.String())
}

func TestInt128_FromBigInt(t *testing.T) {
	testCases := []*big.Int{
		big.NewInt(10),
		big.NewInt(100000000000000000),
	}
	for _, testCase := range testCases {
		val := Uint128{}
		assert.NoError(t, val.FromBigInt(testCase))
		assert.Equal(t, 0, testCase.Cmp(val.BigInt()))
	}

	testCases = []*big.Int{
		big.NewInt(-10),
		big.NewInt(-94),
		big.NewInt(-100000000000000000),
	}
	for _, testCase := range testCases {
		val := Uint128{}
		assert.Errorf(t, val.FromBigInt(testCase), "negative value cannot be assigned to uint128")
	}
}
