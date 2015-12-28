package webutil

import (
	"crypto/md5"
	"fmt"
	"io"
	"reflect"
)

// StringSliceContains checks that the string slice s contains the element e.
func StringSliceContains(s []string, e string) bool {
	for _, se := range s {
		if e == se {
			return true
		}
	}
	return false
}

// ZeroToNil returns nil when v has its zero value.
func ZeroToNil(v interface{}) interface{} {
	if IsZero(v) {
		return nil
	}
	return v
}

// IsZero checks if the v value has its zero value according to its type.
//
// BUG(arthurwhite): Panics when struct have unexported fields.
func IsZero(v interface{}) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Ptr:
		return IsZero(rv.Elem().Interface())
	case reflect.Func:
		return rv.IsNil()
	case reflect.Map:
		return rv.Len() == 0
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			if !IsZero(rv.Index(i).Interface()) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			if !IsZero(rv.Field(i).Interface()) {
				return false
			}
		}
		return true
	}

	return rv.Interface() == reflect.Zero(rv.Type()).Interface()
}

// ChecksumMD5 returns the MD5 checksum of r content.
func ChecksumMD5(r io.Reader) (sum string, err error) {
	hash := md5.New()
	if _, err = io.Copy(hash, r); err != nil {
		return
	}
	sum = fmt.Sprintf("%x", hash.Sum(nil))
	return
}
