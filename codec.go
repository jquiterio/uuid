/*
 * @file: handle.go
 * @author: Jorge Quitério
 * @copyright (c) 2021 Jorge Quitério
 * @license: MIT
 */

package uuid

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
)

type NullUUID struct {
	UUID  UUID
	Valid bool
}

func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *UUID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	return u.UnmarshalText(s)
}

func (u UUID) String() string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}

func (u *UUID) UnmarshalText(text string) error {
	d := Parse([]byte(strings.ToLower(text)))
	if d == Nil {
		return errors.New("uuid.UnmarshalText: Invalid UUID")
	}
	return u.UnmarshalBinary(d[:])
}

func (u *UUID) UnmarshalBinary(data []byte) error {
	d := Parse(data)
	if d == Nil {
		return errors.New("uuid.UnmarshalBinary: Invalid UUID")
	}
	copy(u[:], d[:])
	return nil
}

func (u UUID) Bytes() []byte {
	return u[:]
}

// Value implements the driver.Valuer interface.
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case UUID:
		*u = src
		return nil
	case []byte:
		if len(src) == 16 {
			return u.UnmarshalBinary(src)
		}
	case string:
		return u.UnmarshalText(src)
	}
	return errors.New("uuid.Value: invalid UUID")
}

func (nu NullUUID) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}
	return nu.UUID.Value()
}

func (nu *NullUUID) Scan(src interface{}) error {
	if src == nil {
		nu.UUID, nu.Valid = Nil, false
		return nil
	}
	nu.Valid = true
	return nu.UUID.Scan(src)
}

func (nu NullUUID) MarshalJSON() ([]byte, error) {
	if !nu.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nu.UUID)
}

func (nu *NullUUID) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		nu.UUID, nu.Valid = Nil, false
		return nil
	}
	err := json.Unmarshal(b, &nu.UUID)
	if err != nil {
		return err
	}
	nu.Valid = true
	return nil
}
