# UUID - RFC4122

<p align="center">
    <p>
      <a href="https://github.com/jquiterio/uuid/actions/workflows/ci.yml?query=branch%3Amain">
        <img src="https://github.com/jquiterio/uuid/actions/workflows/ci.yml/badge.svg" alt="Build Status">
      </a>
      <a href="https://codecov.io/gh/jquiterio/uuid/branch/main">
        <img src="https://codecov.io/gh/jquiterio/uuid/branch/main/graph/badge.svg" alt="Code Coverage">
      </a> 
      <a href="https://goreportcard.com/report/jquiterio/uuid">
        <img src="https://goreportcard.com/badge/jquiterio/uuid" alt="Go Report Card">
      </a>
      <a href="https://github.com/jquiterio/uuid/releases/latest">
        <img src="https://img.shields.io/badge/version-2.7.0-blue.svg" alt="Release Version">
      </a> 
      <a href="https://pkg.go.dev/github.com/jquiterio/uuid">
        <img src="https://pkg.go.dev/badge/github.com/jquiterio/uuid" alt="GoDoc">
      </a> 
      <a href="LICENSE">
        <img src="https://img.shields.io/github/license/jquiterio/uuid.svg" alt="License">
      </a> 
    </p>
</p>

---

Another UUID generator and parser for Go.
Returns UUID or Nil

## Usage

`go get -u github.com/jquiterio/uuid`

```go
func main(){

  // new V4
  u := uuid.New() // or u := uuid.NewV4() generates a new UUID v4.
  u.String() // returns a string uuid
  u.Bytes() // retruns a byte slice

  u5 := uuid.NewV5() // generates a new UUID v5
  u5.String()
  u5.Bytes() // retruns a byte slice

  // Parse UUID
  // Get UUID from String
  ufs := Parse("c5302009-7ff6-47d2-9a1c-72601da3e3e5")
  ufs.String()
  // Get UUID from Bytes
  ufb := Parse(uuid.New().Bytes())
}
```

Go struct with GORM

```go
type User struct {
  UUID uuid.UUID `gorm:"type:uuid" json:"uuid"`
  Name string `json:"name"`
}
```
