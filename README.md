# UUID - RFC4122

---

Another UUID generator and parser for Go.
It return Nil or UUID.

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
