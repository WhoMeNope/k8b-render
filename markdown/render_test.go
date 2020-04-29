package render

import (
	"testing"

	"github.com/WhoMeNope/snapshot"
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

		actual, err := Render([]byte(input))
		if err != nil {
			t.Fatal(err)
		}

		err = snapshot.WithLabel(name).Matches(actual.String())
		// returns an error with a pretty diff if the snapshot doesn't match
		if err != nil {
			// print the label and the diff
			t.Fatalf("Snapshot didn't match.\n - '%s'\n%s",
				name,
				err.Error())
		}
	}
}
