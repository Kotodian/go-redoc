# go-redoc

[![GoDoc](https://godoc.org/github.com/Kotodian/go-redoc?status.svg)](https://godoc.org/github.com/Kotodian/go-redoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kotodian/go-redoc?_=1)](https://goreportcard.com/report/github.com/Kotodian/go-redoc?_=1)

`go-redoc` is an embedded OpenAPI documentation ui for Go using [ReDoc](https://github.com/ReDocly/redoc) and [1.16's embed](https://golang.org/pkg/embed/), with middleware implementations for: `net/http`, `gin` and `echo`. The template is based on the ReDoc's [bundle template](https://github.com/ReDocly/redoc/blob/master/cli/template.hbs) with the script already placed in the html instead of depending on a cdn.

## Usage

```go
import "github.com/Kotodian/go-redoc"

...

doc := redoc.Redoc{
    Title:       "Example API",
    Description: "Example API Description",
    SpecFile:    "./openapi.json",
    SpecPath:    "/openapi.json",
    DocsPath:    "/docs",
}
```

- `net/http`

```go
import (
	"net/http"
	"github.com/Kotodian/go-redoc"
)

...

http.ListenAndServe(address, doc.Handler())
```

- `gin`

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/Kotodian/go-redoc"
	ginredoc "github.com/Kotodian/go-redoc/gin"
)

...

r := gin.New()
r.Use(ginredoc.New(doc))
```

- `echo`

```go
import (
	"github.com/labstack/echo/v4"
	"github.com/Kotodian/go-redoc"
	echoredoc "github.com/Kotodian/go-redoc/echo"
)

...

r := echo.New()
r.Use(echoredoc.New(doc))
```

See [examples](/_examples)
