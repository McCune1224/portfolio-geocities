package handlers

import (
	"geocity/static/pages/pong"

	"github.com/labstack/echo/v4"
)

type Ball struct {
	X int
	Y int
}

// Update the respective X/Y Position for the ball
func (h *Handler) Tick(c echo.Context) error {
	ballOffsetX, ok := c.Get("cookie_x").(int)
	if !ok {
		ballOffsetX = 0
	}
	ballOffsetY, ok := c.Get("cookie_y").(int)
	if !ok {
		ballOffsetY = 0
	}
	ball := &Ball{
		X: ballOffsetX,
		Y: ballOffsetY,
	}
	c.Logger().Print(&ball)
	return Render(c, 200, pong.Ball())
}
