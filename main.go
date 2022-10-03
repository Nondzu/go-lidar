package main

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/nondzu/test1/gogl"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 720
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

	shaderProgram, err := gogl.NewShader("shaders/light.vert", "shaders/quadtexture-light.frag")

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

	normals := make([]float32, 36*3)

	for tri := 0; tri < 12; tri++ {

		index := tri * 15

		p1 := mgl32.Vec3{vertices[index], vertices[index+1], vertices[index+2]}
		index += 5

		p2 := mgl32.Vec3{vertices[index], vertices[index+1], vertices[index+2]}
		index += 5

		p3 := mgl32.Vec3{vertices[index], vertices[index+1], vertices[index+2]}
		normal := gogl.TriangleNormal(p1, p2, p3)

		normals[tri*9] = normal.X()
		normals[tri*9+1] = normal.Y()
		normals[tri*9+2] = normal.Z()

		normals[tri*9+3] = normal.X()
		normals[tri*9+4] = normal.Y()
		normals[tri*9+5] = normal.Z()

		normals[tri*9+6] = normal.X()
		normals[tri*9+7] = normal.Y()
		normals[tri*9+8] = normal.Z()
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

	gogl.GenBindBuffer(gl.ARRAY_BUFFER)
	VAO := gogl.GenBindVertexArray()
	gogl.BufferDataFloat(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	//NAO
	gogl.GenBindBuffer(gl.ARRAY_BUFFER)
	gogl.BufferDataFloat(gl.ARRAY_BUFFER, normals, gl.STATIC_DRAW)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(2)
	//
	//

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

		// mouseX, mouseY, _ := sdl.GetRelativeMouseState()
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

		//set light
		shaderProgram.SetVec3("lightPos", mgl32.Vec3{0.0, 0.0, 1.0})
		shaderProgram.SetVec3("lightColor", mgl32.Vec3{0.4, 0.4, 0.4})
		shaderProgram.SetVec3("ambientColor", mgl32.Vec3{0.5, 0.5, 0.0})
		shaderProgram.SetVec3("viewPos", camera.Position)

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
