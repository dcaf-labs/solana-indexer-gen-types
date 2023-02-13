package int128

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
	Int128
}

func TestInt128ScanValue(t *testing.T) {
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
	find := Table{}
	assert.NoError(t, db.First(&find, 1).Error) // find by gorm ID
	assert.Equal(t, insert.Int128.String(), find.Int128.String())
	assert.Equal(t, "10", find.Int128.String())

	// Insert
	val = Int128{}
	assert.NoError(t, val.FromBigInt(big.NewInt(-10)))
	insert = Table{Int128: val}
	assert.Equal(t, "-10", insert.Int128.DecimalString())
	assert.NoError(t, db.Create(&insert).Error)
	// Read
	find = Table{}
	assert.NoError(t, db.First(&find, 2).Error) // find by gorm ID
	assert.Equal(t, insert.Int128.DecimalString(), find.Int128.DecimalString())
	assert.Equal(t, "-10", find.Int128.DecimalString())
}

func TestInt128_FromBigInt(t *testing.T) {
	testCases := []*big.Int{
		big.NewInt(-9999999999999999),
		big.NewInt(-94),
		big.NewInt(-90),
		big.NewInt(10),
		big.NewInt(100000000000000000),
	}
	for _, testCase := range testCases {
		val := Int128{}
		assert.NoError(t, val.FromBigInt(testCase))
		assert.Equal(t, 0, testCase.Cmp(val.BigInt()))
	}
}
