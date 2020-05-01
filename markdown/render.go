package render

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func Render(source []byte) ([]byte, error) {
	var css bytes.Buffer

	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	p := bluemonday.UGCPolicy()

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("murphy"),
				highlighting.WithCSSWriter(&css),
				highlighting.WithFormatOptions(
					// use classes rather than inline styles
					html.WithClasses(true),
					// copy-friendly line numbers
					html.WithLineNumbers(true),
					html.LineNumbersInTable(true),
					// tab width
					html.TabWidth(4),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)

	var buf bytes.Buffer
	if err := md.Convert(p.SanitizeBytes(source), &buf); err != nil {
		return nil, err
	}

	var output bytes.Buffer
	_, err := fmt.Fprintf(&output, "<style>\n%s</style>\n%s", css.Bytes(), buf.Bytes())
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
