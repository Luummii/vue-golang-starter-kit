package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// HTTPErrorRedirectHandler -
func (s *Store) HTTPErrorRedirectHandler(err error, c echo.Context) {
	codeHTML := fmt.Sprintf(`<!DOCTYPE html><html><head><meta charset=utf-8><meta name=viewport content="width=device-width,initial-scale=1"><title>unrealengine.help</title><link href=/static/css/app.28372e26e62074a07d6c8921bda63f39.css rel=stylesheet></head><body><div id=app></div><script type=text/javascript src=/static/js/manifest.2ae2e69a05c33dfc65f8.js></script><script type=text/javascript src=/static/js/vendor.3fae27b6d0a0572472a3.js></script><script type=text/javascript src=/static/js/app.fd66f988f9f7ca49cd6e.js></script></body></html>`)
	c.HTML(http.StatusOK, codeHTML)
}
