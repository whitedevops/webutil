package webutil

import (
	"reflect"
	"strings"
	"testing"
)

func TestStringSliceContains(t *testing.T) {
	for _, c := range [][3]interface{}{
		{[]string(nil), "", false},
		{[]string{}, "", false},
		{[]string(nil), "foobar", false},
		{[]string{}, "foobar", false},
		{[]string{"foo", "bar"}, "foo", true},
	} {
		if StringSliceContains(c[0].([]string), c[1].(string)) != c[2] {
			t.Errorf("%#v in %#v", c[1], c[0])
		}
	}
}

func TestZeroToNil(t *testing.T) {
	for _, c := range [][2]interface{}{
		{nil, nil},
		{0, nil},
		{1, 1},
		{"", nil},
		{"foobar", "foobar"},
		{[]string(nil), nil},
		{[]string{}, nil},
		{[]string{"foo", "bar"}, []string{"foo", "bar"}},
		{map[string]string(nil), nil},
		{map[string]string{}, nil},
		{map[string]string{"foo": "bar", "bar": "foo"}, map[string]string{"foo": "bar", "bar": "foo"}},
		{struct{ Field string }{}, nil},
		{struct{ Field string }{"foobar"}, struct{ Field string }{"foobar"}},
	} {
		if g := ZeroToNil(c[0]); !reflect.DeepEqual(g, c[1]) {
			t.Errorf("%#v: want %#v, got %#v", c[0], c[1], g)
		}
	}
}

func TestIsZero(t *testing.T) {
	for _, c := range [][2]interface{}{
		{nil, true},
		{0, true},
		{1, false},
		{"", true},
		{"foobar", false},
		{[]string(nil), true},
		{[]string{}, true},
		{[]string{"foo", "bar"}, false},
		{map[string]string(nil), true},
		{map[string]string{}, true},
		{map[string]string{"foo": "bar", "bar": "foo"}, false},
		{struct{ Field string }{}, true},
		{struct{ Field string }{"foobar"}, false},
	} {
		if IsZero(c[0]) != c[1] {
			t.Errorf("%#v", c[0])
		}
	}
}

func TestChecksumMD5(t *testing.T) {
	w := "3858f62230ac3c915f300c664312c63f"
	g, err := ChecksumMD5(strings.NewReader("foobar"))
	if err != nil {
		t.Fatal(err)
	}
	if w != g {
		t.Errorf("want %#v, got %#v", w, g)
	}
}
