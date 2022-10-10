package main

import (
	"math"

	"github.com/PerformLine/go-stockutil/colorutil"
	"github.com/go-gl/mathgl/mgl32"
)

func valueToRGB(v, max float32) mgl32.Vec4 {

	hMax := 240
	v2 := float32(hMax) / max * v

	hue := v2
	r, g, b := colorutil.HsvToRgb(float64(hue), float64(1), float64(1))

	ret := mgl32.Vec4{float32(r) / 256, float32(g) / 256, float32(b) / 256, 1.0}
	return ret
}

// https://stackoverflow.com/questions/30619901/calculate-3d-point-coordinates-using-horizontal-and-vertical-angles-and-slope-di
/*
distance - distance
theta: the zenith angle,
phi: the azimuth angle
*/
func calculateCubePos(distance, theta, phi float32) mgl32.Vec3 {

	x := distance * float32(math.Sin(float64(mgl32.DegToRad(theta)))*math.Cos(float64(mgl32.DegToRad(phi))))
	y := distance * float32(math.Sin(float64(mgl32.DegToRad(theta)))*math.Sin(float64(mgl32.DegToRad(phi))))
	z := distance * float32(math.Cos(float64(mgl32.DegToRad(theta))))

	return mgl32.Vec3{x, y, z}
}
