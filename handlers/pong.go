package handlers

import (
	ponggame "geocity/pong"
	"geocity/static/pages/pong"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RenderBoardHandler(c echo.Context) error {
	board := ponggame.NewBoard(ponggame.GlobalBoardHeight, ponggame.GlobalBoardWidth)
	return Render(c, 200, pong.LiveBoard(board.Grid))
}

func (h *Handler) NewGameHandler(c echo.Context) error {
	cookies := c.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "game_id" {
			h.DB.DeletePongSession(c.Request().Context(), cookie.Value)
		}
	}
	gameId := uuid.New().String()
	_, err := h.DB.CreatePongSession(
		c.Request().Context(),
		gameId,
		ponggame.NewBoard(ponggame.GlobalBoardHeight, ponggame.GlobalBoardWidth))
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:  "game_id",
		Value: gameId,
	})
	return c.Redirect(http.StatusFound, "/pong/update")
}

func (h *Handler) UpdateGameHandler(c echo.Context) error {
	cookies := c.Cookies()

	gameId := ""
	dirCookie := ""
	for _, cookie := range cookies {
		if cookie.Name == "game_id" {
			gameId = cookie.Value
		}
		if cookie.Name == "direction" {
			dirCookie = cookie.Value
		}
		if gameId == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/pong/new")
		}
	}
	board, err := h.DB.GetBoard(c.Request().Context(), gameId)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	playerDirection := ponggame.Direction{}
	switch dirCookie {
	case "u":
		playerDirection = ponggame.DirectionUp
	case "d":
		playerDirection = ponggame.DirectionDown
	default:
		playerDirection = ponggame.DirectionNothing
	}
	err = board.UpdateBoard(playerDirection)
	if err != nil {
		c.Logger().Error(err)
		return Render(c, 200, pong.StartGameButton())
	}
	_, err = h.DB.UpdatePongSession(c.Request().Context(), gameId, board)
	if err != nil {
		c.Logger().Error(err)
		return Render(c, 200, pong.StartGameButton())
	}

	return Render(c, 200, pong.LiveBoard(board.Grid))

}
