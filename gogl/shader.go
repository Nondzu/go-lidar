package gogl

import (
	"fmt"
	"os"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
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

	result := &Shader{id, vertexPath, framentPath, getModifiedTime(vertexPath), getModifiedTime(framentPath)}
	loadedShaders[id] = result

	return result, nil
}

func (shader *Shader) Use() {
	// shader := loadedShaders[id]
	UseProgram(shader.id)
}

func (shader *Shader) SetFloat(name string, value float32) {
	name_cstr := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(uint32(shader.id), name_cstr)
	gl.Uniform1f(location, value)
}

func (shader *Shader) CheckShadersForChanges() {
	// for _, shader := range loadedShaders {

	vertexModTime := getModifiedTime(shader.vertexPath)
	fragmentModTime := getModifiedTime(shader.framgentPath)

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
	// }
}

func getModifiedTime(filePath string) time.Time {
	file, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}

	return file.ModTime()
}
