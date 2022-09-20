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

func CreateProgram(vertexShader ShaderId, fragmentShader ShaderId, shaderProgram ProgramId) {
	// program := gl.CreateProgram()
	gl.AttachShader(uint32(shaderProgram), uint32(vertexShader))
	gl.AttachShader(uint32(shaderProgram), uint32(fragmentShader))
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
	gl.DeleteShader(uint32(vertexShader))
	gl.DeleteShader(uint32(fragmentShader))
}
