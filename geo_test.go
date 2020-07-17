/*
 * Author Linvon
 */

package geohelper

import "testing"

func TestNewGeoMap(t *testing.T) {
	_, err := NewGeoMap("testdata/China.json", "name")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewGeoMapList(t *testing.T) {
	_, errs := NewGeoMapList([]string{"testdata/China.json", "testdata/Peking.json"}, []string{"name", "name"})
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
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
