package ctx

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type RequestContext struct {
	Role     string `header:"USER-ROLE"`
	Username string `header:"USER-NAME"`
}

func FromEchoContext(c echo.Context) *RequestContext {
	req := new(RequestContext)
	if err := (&echo.DefaultBinder{}).BindHeaders(c, req); err != nil {
		panic(fmt.Sprintf("failed to build request context due to %s", err))
	}

	return req
}
