/*
 * Author Linvon
 */

package geohelper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	geo "github.com/kellydunn/golang-geo"
	geojson "github.com/paulmach/go.geojson"
)

const StrNotFound = "NotFound"

type GeoMap struct {
	GMap map[string][]*geo.Polygon
	File string
}

func getPolyMap(file string, keys []string, bm bool, ff func(ks []string) string) (map[string][]*geo.Polygon, error) {
	t := time.Now()
	provinces, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fc1, err := geojson.UnmarshalFeatureCollection(provinces)
	if err != nil {
		return nil, err
	}

	polysMap := make(map[string][]*geo.Polygon, 0)
	for _, v := range fc1.Features {
		pKeys := make([]string, len(keys))
		for i, key := range keys {
			if _, ok := v.Properties[key]; !ok {
				return nil, errors.New(fmt.Sprintf("file:%v has no key:%v in some features", file, key))
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
	if bm {
		fmt.Printf("File %-30v loaded, Got %-2v area, Elapse %v\n", file, len(fc1.Features), time.Since(t))
	}
	return polysMap, nil
}

func NewGeoMap(file, key string) (*GeoMap, error) {
	g, e := getPolyMap(file, []string{key}, false, func(ks []string) string {return ks[0]})
	if e != nil {
		return nil, e
	}
	return &GeoMap{g, file}, nil
}

func NewGeoMapFormat(file string, keys []string, ff func(ks []string) string) (*GeoMap, error) {
	g, e := getPolyMap(file, keys, false, ff)
	if e != nil {
		return nil, e
	}
	return &GeoMap{g, file}, nil
}

func NewGeoMapWithBenchmark(file, key string) (*GeoMap, error) {
	g, e := getPolyMap(file, []string{key}, true, func(ks []string) string {return ks[0]})
	if e != nil {
		return nil, e
	}
	return &GeoMap{g, file}, nil
}

func NewGeoMapList(files, keys []string) ([]*GeoMap, []error) {
	if len(files) != len(keys) {
		return nil, []error{errors.New("params are not matched")}
	}
	l := make([]*GeoMap, len(files))
	le := make([]error, len(files))
	wg := new(sync.WaitGroup)
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			g, e := getPolyMap(files[index], []string{keys[index]}, false, func(ks []string) string {return ks[0]})
			if e != nil {
				l[index] = nil
			} else {
				l[index] = &GeoMap{g, files[index]}
			}
			le[index] = e
		}(i, wg)
	}
	wg.Wait()
	return l, le
}

func NewGeoMapListFormat(files []string, keys [][]string, ff func(ks []string) string) ([]*GeoMap, []error) {
	if len(files) != len(keys) {
		return nil, []error{errors.New("params are not matched")}
	}
	l := make([]*GeoMap, len(files))
	le := make([]error, len(files))
	wg := new(sync.WaitGroup)
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			g, e := getPolyMap(files[index], keys[index], false, ff)
			if e != nil {
				l[index] = nil
			} else {
				l[index] = &GeoMap{g, files[index]}
			}
			le[index] = e
		}(i, wg)
	}
	wg.Wait()
	return l, le
}

func NewGeoMapListWithBenchmark(files, keys []string) ([]*GeoMap, []error) {
	if len(files) != len(keys) {
		return nil, []error{errors.New("params are not matched")}
	}
	l := make([]*GeoMap, len(files))
	le := make([]error, len(files))
	wg := new(sync.WaitGroup)
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			g, e := getPolyMap(files[index], []string{keys[index]}, true, func(ks []string) string {return ks[0]})
			if e != nil {
				l[index] = nil
			} else {
				l[index] = &GeoMap{g, files[index]}
			}
			le[index] = e
		}(i, wg)
	}
	wg.Wait()
	return l, le
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

func (g *GeoMap) FindPointBenchmark(p *geo.Point) {
	t := time.Now()
	fmt.Printf("File %-30v Result: %-15v  Elapse: %v\n", g.File, g.FindPoint(p), time.Since(t))
}

func (g *GeoMap) FindLocBenchmark(lat, lng float64) {
	p := geo.NewPoint(lat, lng)
	t := time.Now()
	fmt.Printf("File %-30v Result: %-15v  Elapse: %v\n", g.File, g.FindPoint(p), time.Since(t))
}
