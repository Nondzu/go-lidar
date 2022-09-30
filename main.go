package main

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/nondzu/test1/gogl"

	"github.com/veandco/go-sdl2/sdl"
)

// Go 42 	https://www.youtube.com/watch?v=XjOrF_OmXsY&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x&index=45&ab_channel=JackMott - work on
// Go 43	https://www.youtube.com/watch?v=ogxPnneEPSM&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x&index=43&ab_channel=JackMott - todo
const winWidth = 720
const winHeight = 480

func mglTest() {
	x := mgl32.NewVecN(2)
	fmt.Printf("x: %v\n", x)
}

func main() {

	mglTest()
	// projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(1)/1, 0.1, 10.0)
	// _ = projection
	err := sdl.Init(uint32(sdl.INIT_EVERYTHING))

	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)

	window, err := sdl.CreateWindow("Hello triangle", 50, 50, winWidth, winHeight, uint32(sdl.WINDOW_OPENGL))

	if err != nil {
		panic(err)
	}

	window.GLCreateContext()

	defer window.Destroy()

	gl.Init()
	fmt.Println("OpenGL Version: ", gogl.GetVersion())

	shaderProgram, err := gogl.NewShader("shaders/hello.vert", "shaders/quadtexture.frag")

	if err != nil {
		panic(err)
	}

	texture := gogl.LoadTextureAlpha("assets/tex.png")

	vertices := []float32{
		0.5, 0.5, 0.0, 0.5, 0.5,
		0.5, -0.5, 0.0, 0.5, 0.0,
		-0.5, -0.5, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.0, 0.0, 0.5,
	}

	indices := []uint32{
		0, 1, 3, // triangle 1
		1, 2, 3, // triangle 2
	}

	gogl.GenBindBuffer(gl.ARRAY_BUFFER)
	VAO := gogl.GenBindVertexArray()
	gogl.BufferDataFloat(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)
	gogl.GenBindBuffer(gl.ELEMENT_ARRAY_BUFFER)
	gogl.BufferDataInt(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gogl.UnbindVertexArray()

	var x float32 = 1.0

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// gogl.UseProgram(shaderProgram)
		shaderProgram.Use()

		shaderProgram.SetFloat("x", x)
		shaderProgram.SetFloat("y", 0.0)
		gogl.BindTexture(texture)

		gogl.BindVertexArray(VAO)

		// gl.BindFramebuffer(gl.ARRAY_BUFFER, uint32(VBO))
		// gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))

		window.GLSwap()

		shaderProgram.CheckShadersForChanges()

		x = x + 0.05
	}
}
