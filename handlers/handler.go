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
	e.GET("/", h.HomeHandler)
	e.GET("/spooky", h.SpookyHandler)
	e.GET("/gaming", h.GamingHandler)
	e.GET("/ball", h.Tick)
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

func (h *Handler) HomeHandler(c echo.Context) error {
	viewTally, err := h.DB.IncrementSiteViewCount(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	return Render(c, http.StatusOK, pages.Index(viewTally))
}

func (h *Handler) SpookyHandler(c echo.Context) error {

	return Render(c, http.StatusOK, pages.SpookyPage())
}

func (h *Handler) GamingHandler(c echo.Context) error {

	return Render(c, http.StatusOK, pages.GamingPage())
}
