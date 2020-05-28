**ARCHIVED : this project was merged into [k8b](https://github.com/WhoMeNope/k8b).**

# k8b-render

[![Go Report Card](https://goreportcard.com/badge/github.com/WhoMeNope/k8b-render)](https://goreportcard.com/report/github.com/WhoMeNope/k8b-render)

Renderer library for [k8b](https://github.com/WhoMeNope/k8b).

## Preview

Includes a simple CLI tool to render files.

`go run main.go [path to a file to preview]`

When run will:

1. Print the raw contents of the file
2. Print the renderer HTML for this file
3. Start a web server and serve the rendered file at
   [http://localhost:5000](http://localhost:5000)

## Library

Use this as a HTML renderer library.

```go
import (
    markdown "github.com/WhoMeNope/k8b-render/markdown"
)
```

```go
rendered, err := markdown.NewRenderer().Render(data)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(rendered))
```

## Internals

Uses [chroma](https://github.com/alecthomas/chroma) for code snippet highlighting.

