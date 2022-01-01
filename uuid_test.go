/*
 * @file: uuid_test.go
 * @author: Jorge Quitério
 * @copyright (c) 2021 Jorge Quitério
 * @license: MIT
 */

package uuid

import (
	"testing"
)

var (
	v4uuid string = "c5302009-7ff6-47d2-9a1c-72601da3e3e5"
	v5uuid string = "0b5bcdbf-1feb-5813-943d-8c325c7fe5bb"
)

func TestV4ParseString(t *testing.T) {
	uuid := Parse(v4uuid)
	if uuid.String() != v4uuid {
		t.Errorf("Expected %s, got %s", v4uuid, uuid.String())
	} else {
		t.Logf("uuid: %s", uuid.String())
	}
}

func TestV4ParseBytes(t *testing.T) {
	u := NewV4()
	b := u.Bytes()
	v := Parse(b)
	if u.String() != v.String() {
		t.Errorf("expected value %v, got %v", v.String(), u.String())
	}
}

func TestV5ParseString(t *testing.T) {
	uuid := Parse(v5uuid)
	if uuid.String() != v5uuid {
		t.Errorf("Expected %s, got %s", v5uuid, uuid.String())
	} else {
		t.Logf("uuid: %s", uuid.String())
	}
}

func TestV4SQLValue(t *testing.T) {
	u := NewV4()
	v, err := u.Value()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	s := u.String()
	if v != s {
		t.Errorf("expected value %v, got %v", s, v)
	}
}

func TestV5SQLValue(t *testing.T) {
	u := NewV5(NSDNS, "example.com")
	v, err := u.Value()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	s := u.String()
	if v != s {
		t.Errorf("expected value %v, got %v", s, v)
	}
}

func TestV4SQLScanText(t *testing.T) {
	u := NewV4()
	v := u.String()
	u.Scan(v)
	if u.String() != v {
		t.Errorf("expected value %v, got %v", v, u.String())
	}
}

func TestV5SQLScanText(t *testing.T) {
	u := NewV5(NSDNS, "example.com")
	v := u.String()
	u.Scan(v)
	if u.String() != v {
		t.Errorf("expected value %v, got %v", v, u.String())
	}
}

func TestV4SQLScanBinary(t *testing.T) {
	u := NewV4()
	b := u.Bytes()
	v := Parse(b)
	u.Scan(v)
	if u.String() != v.String() {
		t.Errorf("expected value %v, got %v", v.String(), u.String())
	}
}
