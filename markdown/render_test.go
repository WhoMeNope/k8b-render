package render

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

const multiline = `
# Title

Content
`

const code = `
# Title

` + "```go" + `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
` + "```" + `

More content here.

## Heading level 2

More content.

` + "```go" + `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
` + "```"

var testCases = map[string]string{
	"basic":     "# Title",
	"multiline": multiline,
	"code":      code,
}

// TestRender executes all testCases for Render by checking with snapshots.
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
