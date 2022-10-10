package main

import (
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

func calculateCubePos(angle, distance float32) {

}
