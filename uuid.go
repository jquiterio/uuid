/*
 * @file: uuid.go
 * @author: Jorge Quitério
 * @copyright (c) 2021 Jorge Quitério
 * @license: MIT
 */

package uuid

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"hash"
	"io"
)

// UUID is a 128 bit (16 byte) Universal Unique IDentifier.
type UUID [16]byte

// Nil is the Nil UUID.
var Nil = UUID{}

var (
	// NSDNS is the IETF namespace for the DNS.
	NSDNS UUID = Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	// NSURL is the IETF namespace for the URL.
	NSURL UUID = Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	// NSOID is the IETF namespace for the OID.
	NSOID UUID = Parse("6ba7b812-9dad-11d1-80b4-00c04fd430c8")
	// NSX500 is the IETF namespace for the X500.
	NSX500 UUID = Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
)

const (
	// V4 Represents the UUID version 4
	V4 byte = 0x04
	// V5 Represents the UUID version 5
	V5 byte = 0x05
)

const (
	_ byte = iota
	// VariantNCS Represents the UUID variant NCS
	VariantNCS
	// VariantRFC4122 Represents the UUID variant RFC4122
	VariantRFC4122
	// VariantMicrosoft Represents the UUID variant Microsoft
	VariantMicrosoft
	// VariantFuture Represents the UUID variant Future
	VariantFuture
)

// Version returns the version of the UUID
func (u UUID) Version() byte {
	return u[6] >> 4
}

// Variant returns the variant of the UUID
func (u UUID) Variant() byte {
	switch {
	case (u[8] >> 7) == 0x00:
		return VariantNCS
	case (u[8] >> 6) == 0x02:
		return VariantRFC4122
	case (u[8] >> 5) == 0x06:
		return VariantMicrosoft
	case (u[8] >> 5) == 0x07:
		fallthrough
	default:
		return VariantFuture
	}
}

func (u *UUID) setVersion(v byte) {
	u[6] = (u[6] & 0x0F) | (v << 4)
}

func (u *UUID) setVariant() {
	u[8] = (u[8]&(0xff>>2) | (0x02 << 6))
}

// New returns a new UUID based on the string input.
// It returns Nil if the input is not a valid UUID.
func New() UUID {
	return NewV4()
}

func NewString() string {
	return New().String()
}

// NewV4 returns a new random UUID.
// It returns Nil if the input is not a valid UUID.
func NewV4() UUID {
	var u UUID
	if _, err := io.ReadFull(rand.Reader, u[:]); err != nil {
		return Nil
	}
	u.setVariant()
	u.setVersion(V4)
	return u
}

// NewV5 returns a new UUID based on the SHA-1 hash of the namespace UUID and name.
// It returns Nil if the input is not a valid UUID.
func NewV5(ns UUID, name string) UUID {
	u := getHash(sha1.New(), ns, []byte(name))
	u.setVariant()
	u.setVersion(V5)
	return u
}

func getHash(h hash.Hash, ns UUID, name []byte) UUID {
	var u UUID
	h.Write(ns[:])
	h.Write(name)
	copy(u[:], h.Sum(nil))
	return u
}

// ToBytes returns the UUID as a byte slice.
func (u UUID) ToBytes() []byte {
	return u[:]
}

func Equal(u1, u2 UUID) bool {
	return bytes.Equal(u1[:], u2[:])
}

func (u *UUID) Equal(u2 UUID) bool {
	return Equal(*u, u2)
}
