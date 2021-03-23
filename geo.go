/*
 * Author Linvon
 */

package geohelper

import (
	"errors"
	"fmt"
	geo "github.com/kellydunn/golang-geo"
	geojson "github.com/paulmach/go.geojson"
	"io/ioutil"
)

const StrNotFound = "NotFound"

type GeoMap struct {
	GMap map[string][]*geo.Polygon
}

func getPolyMap(data []byte, keys []string, ff func(ks []string) string) (map[string][]*geo.Polygon, error) {
	fc, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		return nil, err
	}

	polysMap := make(map[string][]*geo.Polygon, 0)
	for _, v := range fc.Features {
		pKeys := make([]string, len(keys))
		for i, key := range keys {
			if _, ok := v.Properties[key]; !ok {
				return nil, errors.New(fmt.Sprintf("file has no key:%v in some features", key))
			}
			pKeys[i] = v.Properties[key].(string)
		}
		key := ff(pKeys)

		geometry := v.Geometry
		mps := make([][][][]float64, 0)
		if geometry.Type == "MultiPolygon" {
			mps = geometry.MultiPolygon

		} else if geometry.Type == "Polygon" {
			mps = append(mps, geometry.Polygon)
		}

		for _, polygon := range mps {
			tmpPointList := make([]*geo.Point, 0)
			for _, points := range polygon {
				for _, point := range points {
					tmpPoint := geo.NewPoint(point[1], point[0])
					tmpPointList = append(tmpPointList, tmpPoint)
				}
			}
			polysMap[key] = append(polysMap[key], geo.NewPolygon(tmpPointList))
		}
	}
	return polysMap, nil
}

func NewGeoMap(file, key string) (*GeoMap, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	g, e := getPolyMap(data, []string{key}, func(ks []string) string { return ks[0] })
	if e != nil {
		return nil, e
	}
	return &GeoMap{g}, nil
}

func NewGeoMapFromBytes(data []byte, key string) (*GeoMap, error) {
	g, e := getPolyMap(data, []string{key}, func(ks []string) string { return ks[0] })
	if e != nil {
		return nil, e
	}
	return &GeoMap{g}, nil
}

func NewGeoMapFormat(file string, keys []string, ff func(ks []string) string) (*GeoMap, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	g, e := getPolyMap(data, keys, ff)
	if e != nil {
		return nil, e
	}
	return &GeoMap{g}, nil
}

func NewGeoMapFormatFromBytes(data []byte, keys []string, ff func(ks []string) string) (*GeoMap, error) {
	g, e := getPolyMap(data, keys, ff)
	if e != nil {
		return nil, e
	}
	return &GeoMap{g}, nil
}

func (g *GeoMap) FindPoint(tmpPoint *geo.Point) string {
	for name, polys := range g.GMap {
		for _, poly := range polys {
			if poly.Contains(tmpPoint) {
				return name
			}
		}
	}
	return StrNotFound
}

func (g *GeoMap) ContainPoint(tmpPoint *geo.Point) bool {
	for _, polys := range g.GMap {
		for _, poly := range polys {
			if poly.Contains(tmpPoint) {
				return true
			}
		}
	}
	return false
}

func (g *GeoMap) FindLoc(lat, lng float64) string {
	tmpPoint := geo.NewPoint(lat, lng)
	return g.FindPoint(tmpPoint)
}

func (g *GeoMap) ContainLoc(lat, lng float64) bool {
	tmpPoint := geo.NewPoint(lat, lng)
	return g.ContainPoint(tmpPoint)
}
