package main

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/nondzu/test1/gogl"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 1280
const winHeight = 720

func main() {

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

	vertices := []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}

	cubePositions := []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{3.0, 0.0, 0.0},
		{6.0, 0.0, 0.0},
		{3.0, 3.0, 0.0},
		{3.0, 6.0, 0.0},

		{2.0, 4.5, -15.0},
		{2.0, 5.0, -10.0},
	}
	_ = cubePositions

	gogl.GenBindBuffer(gl.ARRAY_BUFFER)
	VAO := gogl.GenBindVertexArray()
	gogl.BufferDataFloat(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gogl.UnbindVertexArray()

	// var x float32 = 0
	// var z float32 = -3.0

	//new camera
	position := mgl32.Vec3{2.0, 2.5, 10.0}
	worldUp := mgl32.Vec3{0.0, 1.0, 0.0}

	camera := gogl.NewCamera(position, worldUp, -90.0, 0.0, 0.03, 0.05)

	keyboardState := sdl.GetKeyboardState()
	var elapsedTime float32

	for {
		frameStart := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// var dir gogl.Direction
		dir := gogl.Nowhere

		if keyboardState[sdl.SCANCODE_A] != 0 {
			dir = gogl.Left
		}
		if keyboardState[sdl.SCANCODE_D] != 0 {
			dir = gogl.Right
		}

		if keyboardState[sdl.SCANCODE_W] != 0 {
			dir = gogl.Up
		}
		if keyboardState[sdl.SCANCODE_S] != 0 {
			dir = gogl.Down
		}

		if keyboardState[sdl.SCANCODE_Q] != 0 {
			dir = gogl.Forward
		}

		if keyboardState[sdl.SCANCODE_E] != 0 {
			dir = gogl.Backward
		}

		mouseX, mouseY, _ := sdl.GetRelativeMouseState()
		camera.UpdateCamera(dir, elapsedTime, float32(mouseX), float32(mouseY))

		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shaderProgram.Use()
		projectionMatrix := mgl32.Perspective(mgl32.DegToRad(45.0), float32(winWidth)/float32(winHeight), 0.1, 100.0)
		// viewMatrix := mgl32.Ident4()
		viewMatrix := camera.GetViewMatrix()
		// viewMatrix = mgl32.Translate3D(x, -3.0, z)

		shaderProgram.SetMat4("projection", projectionMatrix)
		shaderProgram.SetMat4("view", viewMatrix)
		gogl.BindTexture(texture)
		gogl.BindVertexArray(VAO)

		for i, pos := range cubePositions {
			modelMatrix := mgl32.Ident4()
			modelMatrix = mgl32.Translate3D(pos.X(), pos.Y(), pos.Z()).Mul4(modelMatrix)
			_ = i
			// angle := 90.0 * float32(i)
			// modelMatrix = mgl32.HomogRotate3D(mgl32.DegToRad(angle), mgl32.Vec3{1.0, 0.3, 0.5}).Mul4(modelMatrix)
			shaderProgram.SetMat4("model", modelMatrix)
			gl.DrawArrays(gl.TRIANGLES, 0, 36)
		}

		window.GLSwap()
		shaderProgram.CheckShadersForChanges()
		elapsedTime = float32(time.Since(frameStart).Seconds() * 1000)
	}
}
