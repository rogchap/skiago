# Skia Go

Skia Go provides a binding to the 2D graphics library [Skia](https://skia.org/).

Skia Go's main purpose is to support [PDF Go](https://rogchap.com/pdf); a feature rich PDF generation library backed by the [Skia PDF backend](https://skia.org/docs/user/sample/pdf/).
As such, only a small subset of bindings are required for the PDF Go project. Feature bindings will be added on a needs basis; either by the PDF Go project or by the community wishing to make use of these binding.

## Features

* "go get-able" module; no need to build skia yourself
* Idiomatic Go API, for example: output to an `io.Writer`
* Latest Skia release (chrome/m124)

## Example

```go
var buf bytes.Buffer
doc := skia.NewDocument(buf)
canvas := doc.BeginPage(400, 600, nil)
// Draw using the canvas
doc.EndPage()
doc.Close()
os.WriteFile("output.pdf", buf.Bytes(), 0o666)
``` 
