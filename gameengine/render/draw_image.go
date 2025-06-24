package render

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// LoadTexture loads an image file and creates an OpenGL texture.
func LoadTexture(path string) (uint32, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		unsafe.Pointer(&rgba.Pix[0]),
	)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return texture, nil
}

var (
	quadShaderProgram uint32
	shaderInitialized bool
)

const (
	vertexShaderSource = `#version 330 core
layout(location = 0) in vec3 vert;
layout(location = 1) in vec2 vertTexCoord;
out vec2 fragTexCoord;
uniform mat4 projection;
void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * vec4(vert, 1.0);
}`
	fragmentShaderSource = `#version 330 core
in vec2 fragTexCoord;
out vec4 outputColor;
uniform sampler2D tex;
void main() {
    outputColor = texture(tex, fragTexCoord);
}`
)

func initQuadShader(windowWidth, windowHeight int) {
	if shaderInitialized {
		return
	}
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	csource, free := gl.Strs(vertexShaderSource + "\x00")
	gl.ShaderSource(vertexShader, 1, csource, nil)
	free()
	gl.CompileShader(vertexShader)

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	csource, free = gl.Strs(fragmentShaderSource + "\x00")
	gl.ShaderSource(fragmentShader, 1, csource, nil)
	free()
	gl.CompileShader(fragmentShader)

	quadShaderProgram = gl.CreateProgram()
	gl.AttachShader(quadShaderProgram, vertexShader)
	gl.AttachShader(quadShaderProgram, fragmentShader)
	gl.LinkProgram(quadShaderProgram)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	shaderInitialized = true
}

func ortho(left, right, bottom, top, near, far float32) [16]float32 {
	return [16]float32{
		2 / (right - left), 0, 0, 0,
		0, 2 / (top - bottom), 0, 0,
		0, 0, -2 / (far - near), 0,
		-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(far + near) / (far - near), 1,
	}
}

// DrawTexturedQuad draws a textured quad at (x, y) with given width and height.
// windowWidth and windowHeight are required for orthographic projection.
func DrawTexturedQuadWithWindow(texture uint32, x, y, width, height float32, windowWidth, windowHeight int) {
	initQuadShader(windowWidth, windowHeight)
	gl.UseProgram(quadShaderProgram)

	proj := ortho(0, float32(windowWidth), float32(windowHeight), 0, -1, 1)
	projLoc := gl.GetUniformLocation(quadShaderProgram, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projLoc, 1, false, &proj[0])

	vertices := []float32{
		x, y, 0, 0, 0,
		x + width, y, 0, 1, 0,
		x + width, y + height, 0, 1, 1,
		x, y + height, 0, 0, 1,
	}
	indices := []uint32{0, 1, 2, 2, 3, 0}

	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, unsafe.Pointer(&vertices[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, unsafe.Pointer(&indices[0]), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.Ptr(nil))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*4, gl.Ptr(uintptr(3*4)))

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	texLoc := gl.GetUniformLocation(quadShaderProgram, gl.Str("tex\x00"))
	gl.Uniform1i(texLoc, 0)

	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteBuffers(1, &ebo)
	gl.DeleteVertexArrays(1, &vao)
	gl.UseProgram(0)
}
