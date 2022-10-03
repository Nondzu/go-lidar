package gogl

import (
	"fmt"
	"os"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var loadedShaders = make(map[ProgramId]*Shader)

type Shader struct {
	id               ProgramId
	vertexPath       string
	framgentPath     string
	vertexModified   time.Time
	fragmentModified time.Time
}

func NewShader(vertexPath string, framentPath string) (*Shader, error) {
	id, err := CreateProgram(vertexPath, framentPath)

	if err != nil {
		return nil, err
	}

	vertTime, err := getModifiedTime(vertexPath)
	if err != nil {
		return nil, err
	}

	fragTime, err := getModifiedTime(framentPath)
	if err != nil {
		return nil, err
	}

	result := &Shader{id, vertexPath, framentPath, vertTime, fragTime}
	loadedShaders[id] = result

	return result, nil
}

func (shader *Shader) Use() {
	UseProgram(shader.id)
}

func (shader *Shader) SetFloat(name string, value float32) {
	name_cstr := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(uint32(shader.id), name_cstr)
	gl.Uniform1f(location, value)
}

func (shader *Shader) SetMat4(name string, mat mgl32.Mat4) {
	name_cstr := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(uint32(shader.id), name_cstr)
	m4 := [16]float32(mat)
	gl.UniformMatrix4fv(location, 1, false, &m4[0])
}

func (shader *Shader) SetVec3(name string, v mgl32.Vec3) {
	name_cstr := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(uint32(shader.id), name_cstr)
	v3 := [3]float32(v)
	gl.Uniform3fv(location, 1, &v3[0])
}

func (shader *Shader) CheckShadersForChanges() error {

	vertexModTime, err := getModifiedTime(shader.vertexPath)
	if err != nil {
		return err
	}

	fragmentModTime, err := getModifiedTime(shader.framgentPath)
	if err != nil {
		return err
	}

	if !vertexModTime.Equal(shader.vertexModified) ||
		!fragmentModTime.Equal(shader.fragmentModified) {
		id, err := CreateProgram(shader.vertexPath, shader.framgentPath)

		if err != nil {
			fmt.Println(err)
		} else {
			gl.DeleteProgram(uint32(shader.id))
			shader.id = id
		}
	}
	return nil
}

func getModifiedTime(filePath string) (time.Time, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}

	return file.ModTime(), nil
}
