package gameengine

type GameObject interface {
	Start()
	Update()
	Render()
	Destroy()
}
