package geo4go

import "math"

const (
	xPi = 3.14159265358979324 * 3000.0 / 180.0
	a   = 6378245.0
	ee  = 0.00669342162296594323
)

func transformLat(x, y float64) float64 {
	var ret = -100.0 + 2.0*x + 3.0*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(y*math.Pi) + 40.0*math.Sin(y/3.0*math.Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(y/12.0*math.Pi) + 320*math.Sin(y*math.Pi/30.0)) * 2.0 / 3.0
	return ret
}

func transformLon(x, y float64) float64 {
	var ret = 300.0 + x + 2.0*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20.0*math.Sin(6.0*x*math.Pi) + 20.0*math.Sin(2.0*x*math.Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(x*math.Pi) + 40.0*math.Sin(x/3.0*math.Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(x/12.0*math.Pi) + 300.0*math.Sin(x/30.0*math.Pi)) * 2.0 / 3.0
	return ret
}

func transform(p Point) Point {
	if isOutOfChina(p) {
		return p
	}
	var dLat = transformLat(p.Longitude-105.0, p.Latitude-35.0)
	var dLon = transformLon(p.Longitude-105.0, p.Latitude-35.0)
	var radLat = p.Latitude / 180.0 * math.Pi
	var magic = math.Sin(radLat)
	magic = 1 - ee*magic*magic
	var sqrtMagic = math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * math.Pi)
	dLon = (dLon * 180.0) / (a / sqrtMagic * math.Cos(radLat) * math.Pi)
	var nLat = p.Latitude + dLat
	var nLon = p.Longitude + dLon
	return Point{Latitude: nLat, Longitude: nLon}
}

func isOutOfChina(p Point) bool {
	if p.Longitude < 72.004 || p.Longitude > 137.8347 || p.Latitude < 0.8293 || p.Latitude > 55.8271 {
		return true
	}
	return false
}

// GPS84ToGCJ02 84 to 火星坐标系 (GCJ-02) World Geodetic System ==> Mars Geodetic System
func GPS84ToGCJ02(p Point) Point {
	if isOutOfChina(p) {
		return p
	}
	var dLat = transformLat(p.Longitude-105.0, p.Latitude-35.0)
	var dLon = transformLon(p.Longitude-105.0, p.Latitude-35.0)
	var radLat = p.Latitude / 180.0 * math.Pi
	var magic = math.Sin(radLat)
	magic = 1 - ee*magic*magic
	var sqrtMagic = math.Sqrt(magic)
	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * math.Pi)
	dLon = (dLon * 180.0) / (a / sqrtMagic * math.Cos(radLat) * math.Pi)
	var nLat = p.Latitude + dLat
	var nLon = p.Longitude + dLon
	return Point{Latitude: nLat, Longitude: nLon}
}

// GCJ02ToGPS84 火星坐标系 (GCJ-02) to 84
func GCJ02ToGPS84(p Point) Point {
	var gps = transform(p)
	var nLon = p.Longitude*2 - gps.Longitude
	var nLat = p.Latitude*2 - gps.Latitude
	return Point{Latitude: nLat, Longitude: nLon}
}

// GCJ02ToBD09 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 将 GCJ-02 坐标转换成 BD-09 坐标
func GCJ02ToBD09(p Point) Point {
	var x = p.Longitude
	var y = p.Latitude
	var z = math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*xPi)
	var theta = math.Atan2(y, x) + 0.000003*math.Cos(x*xPi)
	var nLon = z*math.Cos(theta) + 0.0065
	var nLat = z*math.Sin(theta) + 0.006
	return Point{Latitude: nLat, Longitude: nLon}
}

// BD09ToGCJ02 火星坐标系 (GCJ-02) 与百度坐标系 (BD-09) 的转换算法 * * 将 BD-09 坐标转换成GCJ-02 坐标
func BD09ToGCJ02(p Point) Point {
	var x = p.Longitude - 0.0065
	var y = p.Latitude - 0.006
	var z = math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPi)
	var theta = math.Atan2(y, x) - 0.000003*math.Cos(x*xPi)
	var nLon = z * math.Cos(theta)
	var nLat = z * math.Sin(theta)
	return Point{Latitude: nLat, Longitude: nLon}
}

// GPS84ToBD09 将gps84转为bd09
func GPS84ToBD09(p Point) Point {
	var gcj02 = GPS84ToGCJ02(p)
	var bd09 = GCJ02ToBD09(gcj02)
	return bd09
}

// BD09ToGPS84 将bd09转为gps84
func BD09ToGPS84(p Point) Point {
	var gcj02 = BD09ToGCJ02(p)
	var gps84 = GCJ02ToGPS84(gcj02)
	return gps84
}
