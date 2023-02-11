package publickey

import (
	"database/sql/driver"
	"fmt"
	"github.com/gagliardetto/solana-go"
)

// TODO: Add this to the official https://github.com/gagliardetto/solana-go

type PublicKey struct {
	solana.PublicKey `gorm:"-:all"`
}

// Scan assigns a value from a database driver.
func (p *PublicKey) Scan(value any) error {
	if value == nil {
		return p.PublicKey.Set("")
	}
	if convertValue, err := driver.String.ConvertValue(value); err == nil {
		switch s := convertValue.(type) {
		case string:
			if err := p.PublicKey.Set(s); err == nil {
				return nil
			}
		case []byte:
			if err := p.PublicKey.Set(string(s)); err == nil {
				return nil
			}
		}
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, p)
}

// Value returns a driver Value.
func (p *PublicKey) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return p.PublicKey.String(), nil
}
