package model

import "encoding/json"

// DomainObject defines the common interface for domain models.
type DomainObject interface {
	Account | Authority | Book | Category | Format
}

// toString returns the JSON data of the domain models.
func toString[T DomainObject](o *T) string {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(o); err != nil {
		return ""
	}
	return string(bytes)
}
