package int128

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"math/big"
)

// TODO: Add this to the official https://github.com/gagliardetto/binary

type Int128 struct {
	bin.Int128 `gorm:"-:all"`
}

func (p *Int128) FromBigInt(value *big.Int) error {
	if len(value.Bytes()) > 16 {
		return fmt.Errorf("val contains %d extra bytes for int128", 16-len(value.Bytes()))
	}
	bytes := value.FillBytes(make([]byte, 16))
	p.Int128 = bin.Int128{
		Lo:         binary.BigEndian.Uint64(bytes[8:]),
		Hi:         binary.BigEndian.Uint64(bytes[:8]),
		Endianness: binary.BigEndian,
	}
	// https://stackoverflow.com/questions/74709367/converting-big-int-to-2int64-vice-versa-and-twos-complement
	if value.Sign() < 0 {
		// convert the unsigned value to two's compliment
		p.Int128.Lo = ^p.Int128.Lo
		p.Int128.Hi = ^p.Int128.Hi

		p.Int128.Lo++
		// check for carry
		if p.Int128.Lo == 0 {
			p.Int128.Hi++
		}
	}
	return nil
}

// Scan assigns a value from a database driver.
func (p *Int128) Scan(value any) error {
	if value == nil {
		return nil
	}
	switch s := value.(type) {
	case string:
		val, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return fmt.Errorf("failed to convert string to bigint when scanning Int128")
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
		p.Int128 = bin.Int128{
			Lo:         binary.BigEndian.Uint64(s[8:]),
			Hi:         binary.BigEndian.Uint64(s[:8]),
			Endianness: binary.BigEndian,
		}
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, p)
}

// Value returns a driver Value.
func (p *Int128) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return p.String(), nil
}
