package uint128

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"math/big"
)

type Uint128 struct {
	bin.Uint128 `gorm:"-:all"`
}

func (p *Uint128) FromBigInt(value *big.Int) error {
	if value.Sign() < 0 {
		return fmt.Errorf("negative value cannot be assigned to uint128")
	}
	if len(value.Bytes()) > 16 {
		return fmt.Errorf("val contains %d extra bytes for int128", 16-len(value.Bytes()))
	}
	bytes := value.FillBytes(make([]byte, 16))
	p.Lo = binary.BigEndian.Uint64(bytes[8:])
	p.Hi = binary.BigEndian.Uint64(bytes[:8])
	p.Endianness = binary.BigEndian
	return nil
}

// Scan assigns a value from a database driver.
func (p *Uint128) Scan(value any) error {
	if value == nil {
		return nil
	}
	switch s := value.(type) {
	case string:
		val, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return fmt.Errorf("failed to convert string to bigint when scanning Int128")
		}
		if val.Sign() < 0 {
			return fmt.Errorf("negative value cannot be assigned to uint128")
		}
		bytes := val.FillBytes(make([]byte, 16))
		return p.Scan(bytes)
	case int64:
		bytes := make([]byte, 16)
		binary.BigEndian.PutUint64(bytes, uint64(s))
		return p.Scan(bytes)
	case []byte:
		if len(s) > 16 {
			return fmt.Errorf("failed to scan Int128, value larger than 16 bytes")
		}
		p.Lo = binary.BigEndian.Uint64(s[8:])
		p.Hi = binary.BigEndian.Uint64(s[:8])
		p.Endianness = binary.BigEndian
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, p)
}

// Value returns a driver Value.
func (p *Uint128) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return p.String(), nil
}
