package render

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

const multiline = `
# Title

Content
`

var testCases = map[string]string{
	"basic":     "# Title",
	"multiline": multiline,
}

func TestRender(t *testing.T) {
	// iterate over all tests
	for name, input := range testCases {
		t.Run(name, func(t *testing.T) {
			actual, err := Render([]byte(input))
			if err != nil {
				t.Fatal(err)
			}

			cupaloy.SnapshotT(t, actual.String())
		})
	}
}
