package common

type Renderer interface {
	Render([]byte) ([]byte, error)
}
