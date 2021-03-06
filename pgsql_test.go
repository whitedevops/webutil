package webutil

import (
	"reflect"
	"testing"
)

func TestPGPoint(t *testing.T) {
	w := "POINT(20.000000 10.000000)"
	g := PGPoint(10, 20)
	if w != g {
		t.Errorf("want %#v, got %#v", w, g)
	}
}

func TestPGArrayMarshal(t *testing.T) {
	for _, c := range [][2]interface{}{
		{[]string(nil), nil},
		{[]string{}, nil},
		{[]string{"foo", "bar"}, "{foo,bar}"},
	} {
		if g := PGArrayMarshal(c[0].([]string)); g != c[1] {
			t.Errorf("%#v: want %q, got %q", c[0], c[1], g)
		}
	}
}

func TestPGArrayUnmarshal(t *testing.T) {
	for _, c := range [][2]interface{}{
		{"", []string(nil)},
		{"{foo,bar}", []string{"foo", "bar"}},
	} {
		if g := PGArrayUnmarshal(c[0].(string)); !reflect.DeepEqual(g, c[1]) {
			t.Errorf("%#v: want %#v, got %#v", c[0], c[1], g)
		}
	}
}
