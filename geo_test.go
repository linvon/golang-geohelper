/*
 * Author Linvon
 */

package geohelper

import (
	_ "embed"
	"testing"
)

//go:embed testdata/China.json
var jsonChina []byte

//go:embed testdata/Peking.json
var jsonPeking []byte

func TestNewGeoMap(t *testing.T) {
	_, err := NewGeoMap("testdata/China.json", "name")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewGeoMaoFromBytes(t *testing.T) {
	_, err := NewGeoMapFromBytes(jsonChina, "name")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewGeoMapFormat(t *testing.T) {
	ff := func(ks []string) string {
		return ks[0] + "-" + ks[1]
	}

	_, err := NewGeoMapFormat("testdata/Peking.json", []string{"name", "level"}, ff)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewGeoMapFormatFromBytes(t *testing.T) {
	ff := func(ks []string) string {
		return ks[0] + "-" + ks[1]
	}

	_, err := NewGeoMapFormatFromBytes(jsonPeking, []string{"name", "level"}, ff)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGeoMap_FindLoc(t *testing.T) {
	m, err := NewGeoMap("testdata/China.json", "name")
	if err != nil {
		t.Fatal(err)
	}

	if m.FindLoc(39.905241, 116.397682) != "北京市" {
		t.Error("loc not match")
	}

	if m.FindLoc(53.467057, 108.156513) != StrNotFound {
		t.Error("loc not match")
	}
}

func TestGeoMap_FindLocFormat(t *testing.T) {
	ff := func(ks []string) string {
		return ks[0] + "-" + ks[1]
	}

	m, err := NewGeoMapFormat("testdata/Peking.json", []string{"name", "level"}, ff)
	if err != nil {
		t.Fatal(err)
	}

	if m.FindLoc(39.905241, 116.397682) != "东城区-district" {
		t.Error("loc not match")
	}

	if m.FindLoc(53.467057, 108.156513) != StrNotFound {
		t.Error("loc not match")
	}
}

func TestGeoMap_ContainLoc(t *testing.T) {
	m, err := NewGeoMap("testdata/China.json", "name")
	if err != nil {
		t.Fatal(err)
	}

	if !m.ContainLoc(39.905241, 116.397682) {
		t.Error("loc not match")
	}

	if m.ContainLoc(53.467057, 108.156513) {
		t.Error("loc not match")
	}
}
