/*
 * @file: sql.go
 * @author: Jorge Quitério
 * @copyright (c) 2022 Jorge Quitério
 * @license: MIT
 */

package uuid

import (
	"database/sql/driver"
	"fmt"
)

// Value implements the driver.Valuer interface.
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}

// Scan implements the sql.Scanner interface.
func (u *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case UUID:
		*u = src
		return nil
	case nil:
		return nil
	case []byte:
		if len(src) == 0 {
			return nil
		}
		copy((*u)[:], src)
	case string:
		if src == "" {
			return nil
		}
		uuid := Parse(src)
		if uuid == Nil {
			return fmt.Errorf("uuid.Scan: invalid UUID: %s", src)
		}
		*u = uuid
	default:
		return fmt.Errorf("uuid.Scan: invalid type %T", src)
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (nu NullUUID) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}
	return nu.UUID.Value()
}

// Scan implements the sql.Scanner interface.
func (nu *NullUUID) Scan(src interface{}) error {
	if src == nil {
		nu.UUID, nu.Valid = Nil, false
		return nil
	}
	nu.Valid = true
	return nu.UUID.Scan(src)
}
