package main

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/nondzu/test1/gogl"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 1700
const winHeight = 900

func main() {

	go calcAngle()
	err := sdl.Init(uint32(sdl.INIT_EVERYTHING))

	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)

	window, err := sdl.CreateWindow("Hello triangle", 50, 50, winWidth, winHeight, uint32(sdl.WINDOW_OPENGL))

	sdl.SetRelativeMouseMode(true)

	if err != nil {
		panic(err)
	}

	window.GLCreateContext()

	defer window.Destroy()

	gl.Init()
	gl.Enable(gl.DEPTH_TEST)

	fmt.Println("OpenGL Version: ", gogl.GetVersion())

	shaderProgram, err := gogl.NewShader("shaders/hello.vert", "shaders/quadtexture.frag")

	if err != nil {
		panic(err)
	}

	texture := gogl.LoadTextureAlpha("assets/tex2.png")

	cubePositions := []mgl32.Vec3{
		{0.0, 0.0, 0.0},
	}

	gogl.GenBindBuffer(gl.ARRAY_BUFFER)
	VAO := gogl.GenBindVertexArray()

	gogl.BufferDataFloat(gl.ARRAY_BUFFER, cube1, gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)
	gogl.UnbindVertexArray()

	//new camera
	position := mgl32.Vec3{2.0, 2.5, 10.0}
	worldUp := mgl32.Vec3{0.0, 1.0, 0.0}
	camera := gogl.NewCamera(position, worldUp, -90.0, 0, 0.02, 0.15)

	keyboardState := sdl.GetKeyboardState()
	var elapsedTime float32
	var mouseX, mouseY int32

	for {
		frameStart := time.Now()
		mouseX = 0
		mouseY = 0
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch e := event.(type) {
			case *sdl.MouseMotionEvent:
				mouseX = e.XRel
				mouseY = e.YRel

			case *sdl.QuitEvent:
				return

			}
		}

		dir := gogl.Nowhere
		if keyboardState[sdl.SCANCODE_A] != 0 {
			dir = gogl.Left
		}
		if keyboardState[sdl.SCANCODE_D] != 0 {
			dir = gogl.Right
		}

		if keyboardState[sdl.SCANCODE_Q] != 0 {
			dir = gogl.Up
		}
		if keyboardState[sdl.SCANCODE_E] != 0 {
			dir = gogl.Down
		}

		if keyboardState[sdl.SCANCODE_W] != 0 {
			dir = gogl.Forward
		}

		if keyboardState[sdl.SCANCODE_S] != 0 {
			dir = gogl.Backward
		}

		camera.UpdateCamera(dir, elapsedTime, float32(mouseX), float32(mouseY))

		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shaderProgram.Use()
		projectionMatrix := mgl32.Perspective(mgl32.DegToRad(45.0), float32(winWidth)/float32(winHeight), 0.1, 100.0)
		viewMatrix := camera.GetViewMatrix()

		shaderProgram.SetMat4("projection", projectionMatrix)
		shaderProgram.SetMat4("view", viewMatrix)

		gogl.BindTexture(texture)
		gogl.BindVertexArray(VAO) //set triangles array

		for i, pos := range cubePositions {
			modelMatrix := mgl32.Ident4()
			modelMatrix = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix)
			_ = i

			shaderProgram.SetVec4("color", mgl32.Vec4{1.0, 0.3, 0.0, 0.0})
			shaderProgram.SetMat4("model", modelMatrix) // set model
			gl.DrawArrays(gl.TRIANGLES, 0, 36)          //draw model

		}

		//draw demo
		for b := 0; b < 180; b = b + 1 {
			for a := 180; a < 360; a = a + 1 {
				//draw small cube
				{
					pos := calculateCubePos(13, float32(b), float32(a))
					modelMatrix2 := mgl32.Ident4()
					// modelMatrix2 = mgl32.HomogRotate3D(mgl32.DegToRad(float32(angle)), mgl32.Vec3{0.0, 1.0, 0.0}).Mul4(modelMatrix2)
					// modelMatrix2 = mgl32.HomogRotate3D(mgl32.DegToRad(float32(angle)), mgl32.Vec3{1.0, 0.0, 0.0}).Mul4(modelMatrix2)
					modelMatrix2 = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix2)
					shaderProgram.SetVec4("color", mgl32.Vec4{0.0, 0.0, 1.0, 0.0})
					shaderProgram.SetMat4("model", modelMatrix2) // set model
					gl.DrawArrays(gl.TRIANGLES, 36, 73)          //draw model
				}
			}
		}

		drawCube(angle2, angle, shaderProgram)

		{
			modelMatrix2 := mgl32.Ident4()
			modelMatrix2 = mgl32.Translate3D(2, 2, 0.0).Mul4(modelMatrix2)
			shaderProgram.SetVec4("color", mgl32.Vec4{1.0, 1.0, 0.0, 0.0})
			shaderProgram.SetMat4("model", modelMatrix2) // set model
			gl.DrawArrays(gl.TRIANGLES, 36, 73)          //draw model
		}

		{
			modelMatrix2 := mgl32.Ident4()
			modelMatrix2 = mgl32.Translate3D(2, 2, 2.0).Mul4(modelMatrix2)
			shaderProgram.SetVec4("color", mgl32.Vec4{1.0, 1.0, 0.0, 0.0})
			shaderProgram.SetMat4("model", modelMatrix2) // set model
			gl.DrawArrays(gl.TRIANGLES, 36, 73)          //draw model
		}

		//
		window.GLSwap()
		shaderProgram.CheckShadersForChanges()
		elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)
	}
}

var angle float32 = 0.0
var angle2 float32 = 0.0

func drawCube(zenith, azimuth float32, shader *gogl.Shader) {
	pos := calculateCubePos(5, float32(zenith), float32(360-azimuth))
	modelMatrix2 := mgl32.Ident4()
	modelMatrix2 = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix2)
	shader.SetVec4("color", mgl32.Vec4{0.0, 1.0, 0.0, 0.0})
	shader.SetMat4("model", modelMatrix2) // set model
	gl.DrawArrays(gl.TRIANGLES, 36, 73)   //draw model
}

func calcAngle() {
	angle = 0.0
	for {
		time.Sleep(time.Millisecond * 5)
		angle = (angle + float32(1.0))
		if angle > 360 {
			angle = 0
			angle2 = (angle2 + float32(5.0))
			if angle2 > 180 {
				angle2 = 0
			}
		}
	}
}
