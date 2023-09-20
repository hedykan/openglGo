package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width              = 800
	height             = 800
	rows               = 10
	columns            = 10
	fps                = 1
	vertexShaderSource = `
        #version 410
        in vec3 vp;
        void main() {
            gl_Position = vec4(vp, 1.0);
        }
    ` + "\x00"
	fragmentShaderSource = `
        #version 410
        out vec4 frag_colour;
        void main() {
            frag_colour = vec4(1, 0.5, 0.2, 1.0);
        }
    ` + "\x00"
)

var (
	vertices = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}

	indices = []int{
		0, 1, 3,
		1, 2, 3,
	}
)

type Sprite struct {
	x int
	y int
}

func (sprite *Sprite) move(moveType glfw.Key) {
	switch moveType {
	case glfw.KeyRight:
		sprite.x += 1
	case glfw.KeyLeft:
		sprite.x -= 1
	case glfw.KeyUp:
		sprite.y += 1
	case glfw.KeyDown:
		sprite.y -= 1
	}
	sprite.x %= 10
	sprite.y %= 10
}

func main() {
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()
	vao := makeVao(vertices, indices)

	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
	sprite := Sprite{0, 0}

	for !window.ShouldClose() {
		processInput(window, &sprite)
		draw(vao, window, program)
	}
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	gl.BindVertexArray(vao)
	// gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.Ptr(indices))
	glfw.PollEvents()
	window.SwapBuffers()
}
