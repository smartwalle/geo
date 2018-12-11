package geo4go

import "math"

const (
	kEarthRadius = 6378.137 //地球半径
)

// Distance 计算两点之间的距离，返回距离单位为千米
func Distance(p1, p2 Point) float64 {
	var lat1 = degreesToRadians(p1.Latitude)
	var lon1 = degreesToRadians(p1.Longitude)

	var lat2 = degreesToRadians(p2.Latitude)
	var lon2 = degreesToRadians(p2.Longitude)

	var dLat = lat1 - lat2
	var dLon = lon1 - lon2

	var d = 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(dLat/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)))
	d = d * kEarthRadius
	d = math.Round(d*10000) / 10000
	return d
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180.0
}

func radiansToDegrees(r float64) float64 {
	return r * 180.0 / math.Pi
}
