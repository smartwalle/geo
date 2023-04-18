package geo_test

import (
	"github.com/smartwalle/geo"
	"testing"
)

func TestDistance(t *testing.T) {
	// https://lbs.amap.com/api/javascript-api/example/calcutation/calculate-distance-between-two-markers
	t.Log(geo.Distance(geo.Point{Latitude: 39.922501, Longitude: 116.387271}, geo.Point{Latitude: 39.923423, Longitude: 116.368904}))
}
