/*
 * @file: uuid.go
 * @author: Jorge Quitério
 * @copyright (c) 2021 Jorge Quitério
 * @license: MIT
 */

package uuid

import (
	"crypto/rand"
	"crypto/sha1"
	"hash"
	"io"
)

type UUID [16]byte

var Nil = UUID{}

var (
	NSDNS  UUID = Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	NSURL  UUID = Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	NSOID  UUID = Parse("6ba7b812-9dad-11d1-80b4-00c04fd430c8")
	NSX500 UUID = Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
)

const (
	V4 byte = 0x04
	V5 byte = 0x05
)

const (
	_ byte = iota
	VariantNCS
	VariantRFC4122
	VariantMicrosoft
	VariantFuture
)

func (u UUID) Version() byte {
	return u[6] >> 4
}

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

func NewV4() UUID {
	var u UUID
	if _, err := io.ReadFull(rand.Reader, u[:]); err != nil {
		return Nil
	}
	u.setVariant()
	u.setVersion(V4)
	return u
}

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

func (u UUID) ToBytes() []byte {
	return u[:]
}
