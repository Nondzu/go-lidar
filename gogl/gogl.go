package gogl

import (
	"errors"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type ShaderId uint32
type ProgramId uint32
type BufferID uint32

type TextureID uint32

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func LoadShader(path string, shaderType uint32) (ShaderId, error) {
	shaderFile, err := ioutil.ReadFile(path)

	if err != nil {
		return 0, err
		// panic(err)
	}

	shaderFileStr := string(shaderFile)

	shaderId, err := CreateShader(shaderFileStr, shaderType)
	if err != nil {

		return 0, err
	}

	return shaderId, nil
}

func LoadTextureAlpha(filename string) TextureID {
	infile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	if err != nil {
		panic(err)
	}

	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}
	texture := GenBindTexture()

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.REPEAT)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture
}

func GenBindTexture() TextureID {
	var texId uint32

	gl.GenTextures(1, &texId)
	gl.BindTexture(gl.TEXTURE_2D, texId)
	return TextureID(texId)
}

func BindTexture(id TextureID) {
	gl.BindTexture(gl.TEXTURE_2D, uint32(id))
}

func CreateShader(shaderSource string, shaderType uint32) (ShaderId, error) {

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
		fmt.Println("Failed to compile shader : \n" + log)
		return 0, errors.New("failed to compile shader")
	}

	return ShaderId(shaderId), nil
}

func CreateProgram(vertPath string, fragPath string) (ProgramId, error) {

	vert, err := LoadShader(vertPath, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	frag, err := LoadShader(fragPath, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

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

		return 0, fmt.Errorf("Failed to link program: \n" + log)
	}

	// TODO finish hotloading shader
	// file, err := os.Stat(path)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// modTime := file.ModTime()
	// loadedShaders = append(loadedShaders, ShaderInfo{path: path, modTime: modTime})

	gl.DeleteShader(uint32(vert))
	gl.DeleteShader(uint32(frag))

	return ProgramId(shaderProgram), nil
}

func GenBindBuffer(target uint32) BufferID {
	var VBO uint32

	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(target, VBO)
	return BufferID(VBO)
}

func GenBindVertexArray() BufferID {
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	return BufferID(VAO)
}

func GenEBO() BufferID {
	var EBO uint32
	gl.GenBuffers(1, &EBO)
	return BufferID(EBO)
}

func BindVertexArray(vaoID BufferID) {
	gl.BindVertexArray(uint32(vaoID))
}

func BufferDataFloat(target uint32, data []float32, usage uint32) {
	gl.BufferData(target, len(data)*4, gl.Ptr(data), usage)
}

func BufferDataInt(target uint32, data []uint32, usage uint32) {
	gl.BufferData(target, len(data)*4, gl.Ptr(data), usage)
}

func UnbindVertexArray() {
	gl.BindVertexArray(0)
}

func UseProgram(ProgramId ProgramId) {
	gl.UseProgram(uint32(ProgramId))
}
