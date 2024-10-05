package handlers

import (
	"geocity/georedis"
	"geocity/static/pages"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *georedis.GeoDB
}

func NewHandler(db *georedis.GeoDB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) AttachRoutes(e *echo.Echo) {
	e.GET("/", HomeHandler)
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Index())
}
