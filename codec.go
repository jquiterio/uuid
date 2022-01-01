/*
 * @file: handle.go
 * @author: Jorge Quitério
 * @copyright (c) 2021 Jorge Quitério
 * @license: MIT
 */

package uuid

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
)

// NullUUID is a nullable UUID.
type NullUUID struct {
	UUID  UUID
	Valid bool
}

// MarshalJSON implements the json.Marshaler interface.
func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *UUID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	return u.UnmarshalText(s)
}

// String returns the UUID in it's canonical form, a 32
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

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (u *UUID) UnmarshalText(text string) error {
	d := Parse(text)
	if d == Nil {
		return errors.New("uuid.UnmarshalText: Invalid UUID")
	}
	copy(u[:], d[:])
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (u *UUID) UnmarshalBinary(data []byte) error {
	d := Parse(data)
	if d == Nil {
		return errors.New("uuid.UnmarshalBinary: Invalid UUID")
	}
	copy(u[:], d[:])
	return nil
}

// Bytes returns the raw UUID bytes.
func (u UUID) Bytes() []byte {
	return u[:]
}

// MarshalJSON implements the json.Marshaler interface.
func (nu NullUUID) MarshalJSON() ([]byte, error) {
	if !nu.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nu.UUID)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
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
