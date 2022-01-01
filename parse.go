/*
 * @file: parse.go
 * @author: Jorge Quitério
 * @copyright (c) 2022 Jorge Quitério
 * @license: MIT
 */

package uuid

import (
	"encoding/hex"
	"regexp"
	"strings"
)

// Parse decodes s string into a UUID or returns NilUUID if the string is invalid.
// formats:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
func Parse(in interface{}) UUID {
	u := UUID{}
	switch in := in.(type) {
	case string:
		r, _ := regexp.Compile(`^(|{|urn:uuid:)([0-9a-fA-F]{8})\-([0-9a-fA-F]{4})\-([0-9a-fA-F]{4})\-([0-9a-fA-F]{4})\-([0-9a-fA-F]{12})(|})$`)
		if r.MatchString(in) {
			in = strings.Replace(in, "urn:uuid:", "", 1)
			in = strings.Replace(in, "{", "", 1)
			in = strings.Replace(in, "}", "", 1)
			in = strings.Replace(in, "-", "", 4)
		} else {
			return Nil
		}
		b, err := hex.DecodeString(in)
		if err != nil {
			return Nil
		}
		if len(b) != 16 {
			return Nil
		}
		copy(u[:], b)
		return u
	case []byte:
		if len(in) != 16 {
			return Nil
		}
		copy(u[:], in)
		return u
	case UUID:
		return in
	case nil:
		return Nil
	default:
		return Nil
	}

}
