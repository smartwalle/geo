package geo4go

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	// https://lbs.amap.com/api/javascript-api/example/calcutation/calculate-distance-between-two-markers
	fmt.Println(Distance(Point{39.922501, 116.387271}, Point{39.923423, 116.368904}))
}
