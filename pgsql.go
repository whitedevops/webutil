package webutil

import (
	"fmt"
	"strings"
)

// PGPoint returns a PostgreSQL formatted point from longitude and latitude (in this order).
func PGPoint(lat, lon float64) string {
	return fmt.Sprintf("POINT(%f %f)", lon, lat)
}

// PGArrayMarshal returns a PostgreSQL formatted array from a slice, or nil if the slice is empty.
func PGArrayMarshal(a []string) interface{} {
	if len(a) == 0 {
		return nil
	}
	return "{" + strings.Join(a, ",") + "}"
}

// PGArrayUnmarshal returns a slice from a PostgreSQL formatted array.
func PGArrayUnmarshal(s string) (a []string) {
	if s == "" {
		return
	}
	ss := strings.Split(strings.Trim(s, "{}"), ",")
	if len(ss) == 1 && ss[0] == "" {
		return
	}
	for _, e := range ss {
		a = append(a, strings.Trim(e, "\"")) // The string might be wrapped in quotes, so trim them.
	}
	return
}
