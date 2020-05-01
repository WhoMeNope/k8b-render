package common

// Renderer defines behavior of a HTML renderer
type Renderer interface {
	// Render renders HTML of the input
	Render([]byte) ([]byte, error)
}
