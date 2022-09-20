package gogl

import (
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderId uint32
type ProgramId uint32
type VAOID uint32
type VBOID uint32

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func CreateShader(shaderSource string, shaderType uint32) ShaderId {

	shaderSource = shaderSource + "\x00"
	shaderId := gl.CreateShader(shaderType)
	csource, free := gl.Strs(shaderSource)
	gl.ShaderSource(shaderId, 1, csource, nil)
	free()

	var status int32
	gl.CompileShader(shaderId)
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLenght int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLenght)
		log := strings.Repeat("\x00", int(logLenght+1))
		gl.GetShaderInfoLog(shaderId, logLenght, nil, gl.Str(log))

		panic("Failed to compile shader : \n" + log)
	}

	return ShaderId(shaderId)
}

func CreateProgram(vertStr string, fragStr string) ProgramId {

	vert := CreateShader(vertStr, gl.VERTEX_SHADER)
	frag := CreateShader(fragStr, gl.FRAGMENT_SHADER)

	shaderProgram := gl.CreateProgram()
	gl.AttachShader(uint32(shaderProgram), uint32(vert))
	gl.AttachShader(uint32(shaderProgram), uint32(frag))
	gl.LinkProgram(uint32(shaderProgram))

	var success int32
	gl.GetProgramiv(uint32(shaderProgram), gl.LINK_STATUS, &success)

	if success == gl.FALSE {
		var logLenght int32
		gl.GetProgramiv(uint32(shaderProgram), gl.INFO_LOG_LENGTH, &logLenght)
		log := strings.Repeat("\x00", int(logLenght+1))
		gl.GetProgramInfoLog(uint32(shaderProgram), logLenght, nil, gl.Str(log))

		panic("Failed to link program : \n" + log)
	}
	gl.DeleteShader(uint32(vert))
	gl.DeleteShader(uint32(frag))

	return ProgramId(shaderProgram)
}

func GenBindBuffer(target uint32) VBOID {
	var VBO uint32

	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(target, VBO)
	return VBOID(VBO)
}

func GenBindVertexArray() VAOID {
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	return VAOID(VAO)
}

func BindVertexArray(vaoID VAOID) {
	gl.BindVertexArray(uint32(vaoID))
}

func BufferDataFloat(target uint32, data []float32, usage uint32) {
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
}

func UnbindVertexArray() {
	gl.BindVertexArray(0)
}

func UseProgram(ProgramId ProgramId) {
	gl.UseProgram(uint32(ProgramId))
}
