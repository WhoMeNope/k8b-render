package render

import (
	"bytes"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func Render(source []byte) (bytes.Buffer, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
			highlighting.NewHighlighting(
				highlighting.WithStyle("murphy"),
				highlighting.WithFormatOptions(
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
	if err := md.Convert(source, &buf); err != nil {
		return buf, err
	}

	return buf, nil
}
