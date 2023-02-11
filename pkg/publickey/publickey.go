package publickey

import (
	"database/sql/driver"
	"fmt"
	"github.com/gagliardetto/solana-go"
)

type GormPublicKey struct {
	solana.PublicKey `gorm:"-:all"`
}

// Scan assigns a value from a database driver.
//
// The src value will be of one of the following types:
//
//    int64
//    float64
//    bool
//    []byte
//    string
//    time.Time
//    nil - for NULL values
//
// An error should be returned if the value cannot be stored
// without loss of information.
//
// Reference types such as []byte are only valid until the next call to Scan
// and should not be retained. Their underlying memory is owned by the driver.
// If retention is necessary, copy their values before the next call to Scan.
func (p *GormPublicKey) Scan(value any) error {
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
// Value must not panic.
func (p *GormPublicKey) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return p.PublicKey.String(), nil
}
