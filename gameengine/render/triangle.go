package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Triangle2D struct {
	VAO, VBO      uint32
	ShaderProgram uint32
}

// Инициализаци хийх: shader үүсгэх, VAO/VBO үүсгэх
func NewTriangle2D(vertices []float32) *Triangle2D {
	// Шейдерийн эх сурвалж
	vertexSrc := `
		#version 460 core
		layout(location = 0) in vec2 aPos;
		void main() {
			gl_Position = vec4(aPos, 0.0, 1.0);
		}
	` + "\x00"

	fragmentSrc := `
		#version 460 core
		out vec4 FragColor;
		void main() {
			FragColor = vec4(1.0, 0.5, 0.2, 1.0);
		}
	` + "\x00"

	// Шейдер үүсгэх
	vs := compileShader(vertexSrc, gl.VERTEX_SHADER)
	fs := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	program := gl.CreateProgram()
	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)
	gl.LinkProgram(program)
	gl.DeleteShader(vs)
	gl.DeleteShader(fs)

	// VAO, VBO үүсгэх
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 2*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	return &Triangle2D{
		VAO:           vao,
		VBO:           vbo,
		ShaderProgram: program,
	}
}

// Зурах функц
func (t *Triangle2D) Draw() {
	gl.UseProgram(t.ShaderProgram)
	gl.BindVertexArray(t.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}
