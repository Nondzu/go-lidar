package main

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/nondzu/test1/gogl"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 720
const winHeight = 480

func main() {

	err := sdl.Init(uint32(sdl.INIT_EVERYTHING))

	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)

	window, err := sdl.CreateWindow("Hello triangle", 200, 200, winWidth, winHeight, uint32(sdl.WINDOW_OPENGL))

	if err != nil {
		panic(err)
	}

	window.GLCreateContext()

	defer window.Destroy()

	gl.Init()
	fmt.Println("OpenGL Version: ", gogl.GetVersion())

	vertexShaderSource :=
		`#version 460 core
		layout (location = 0) in vec3 aPos;

		void main() 
		{
			gl_Position = vec4(aPos.x,aPos.y,aPos.z,1.0f);
		}`

	vertexShader := gogl.CreateShader(vertexShaderSource, gl.VERTEX_SHADER)

	fragmentShaderSource :=
		`#version 330 core
		out vec4 FragColor;

		void main() 
		{
			FragColor = vec4(1.0f,1.0f,0.0f,0.5f);			
		}`
	fragmentShader := gogl.CreateShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	shaderProgram := gl.CreateProgram()
	gogl.CreateProgram(vertexShader, fragmentShader, gogl.ProgramId(shaderProgram))

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	var VBO uint32

	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(vertices[0])), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindVertexArray(0)

	for {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.ClearColor(0.0, 0.0, 0.0, 0.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.GLSwap()

	}
}
