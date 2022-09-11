package model

import "encoding/json"

type DomainObject interface {
	Account | Authority | Book | Category | Format
}

func toString[T DomainObject](o *T) string {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(o); err != nil {
		return ""
	}
	return string(bytes)
}
