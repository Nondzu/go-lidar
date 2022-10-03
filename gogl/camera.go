package gogl

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Direction int

const (
	Forward Direction = iota
	Backward
	Left
	Right
	Up
	Down
	Nowhere
)

type Camera struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
	Right    mgl32.Vec3
	WordlUp  mgl32.Vec3

	Yaw           float32
	Pitch         float32
	MovementSpeed float32
	MouseSens     float32
	Zoom          float32
}

func NewCamera(position mgl32.Vec3, wordlUp mgl32.Vec3, yaw, pitch, speed, sens float32) *Camera {
	camera := Camera{}
	camera.Position = position
	camera.WordlUp = wordlUp
	camera.Yaw = yaw
	camera.Pitch = pitch
	camera.MovementSpeed = speed
	camera.MouseSens = sens
	camera.updateVectors()

	return &camera
}

func (camera *Camera) UpdateCamera(dir Direction, deltaT, xoffset, yoffset float32) {

	magnitude := camera.MovementSpeed * deltaT
	switch dir {
	case Forward:
		camera.Position = camera.Position.Add(camera.Front.Mul(magnitude))
	case Backward:
		camera.Position = camera.Position.Sub(camera.Front.Mul(magnitude))
	case Left:
		camera.Position = camera.Position.Sub(camera.Right.Mul(magnitude))
	case Right:
		camera.Position = camera.Position.Add(camera.Right.Mul(magnitude))
	case Up:
		camera.Position = camera.Position.Add(camera.Up.Mul(magnitude))
	case Down:
		camera.Position = camera.Position.Sub(camera.Up.Mul(magnitude))
	case Nowhere:
	}

	xoffset *= camera.MouseSens
	yoffset *= camera.MouseSens

	camera.Yaw += xoffset
	camera.Pitch += yoffset
	camera.updateVectors()
}

func (camera *Camera) GetViewMatrix() mgl32.Mat4 {
	center := camera.Position.Add(camera.Front)
	return mgl32.LookAt(camera.Position.X(), camera.Position.Y(), camera.Position.Z(), center.X(), center.Y(), center.Z(), camera.Up.X(), camera.Up.Y(), camera.Up.Z())
}

func (camera *Camera) updateVectors() {
	front := mgl32.Vec3{
		float32(math.Cos(float64(mgl32.DegToRad(camera.Yaw))) * math.Cos(float64(mgl32.DegToRad(camera.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(camera.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(camera.Yaw))) * math.Cos(float64(mgl32.DegToRad(camera.Pitch)))),
	}

	camera.Front = front.Normalize()
	camera.Right = camera.Front.Cross(camera.WordlUp).Normalize()
	camera.Up = camera.Right.Cross(camera.Front).Normalize()
}
