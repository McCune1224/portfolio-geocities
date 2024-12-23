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

func (h *Handler) NewCacheGameHandler(c echo.Context) error {
	h.DB.DeleteCachePongSession(c)
	h.DB.CreateCachePongSession(c, ponggame.NewBoard(ponggame.GlobalBoardHeight, ponggame.GlobalBoardWidth))

	return c.Redirect(http.StatusFound, "/pong/update")
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

func (h *Handler) UpdateCacheGameHandler(c echo.Context) error {
	cookies := c.Cookies()

	gameId := ""
	dirCookie := ""
	for _, cookie := range cookies {
		if cookie.Name == "game_id" {
			gameId = cookie.Value
		}
		if gameId == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/pong/new")
		}
		if cookie.Name == "player_direction" {
			dirCookie = cookie.Value
		}
	}
	board, err := h.DB.GetCacheBoard(c)
	if err != nil {
		c.Logger().Error("UH OH STINKY, CACHE BOARD NOT FOUND")
		c.Redirect(http.StatusTemporaryRedirect, "/pong/new")
	}

	playerDirection := ponggame.Direction{}
	switch dirCookie {
	case "up":
		playerDirection = ponggame.DirectionUp
	case "down":
		playerDirection = ponggame.DirectionDown
	default:
		playerDirection = ponggame.DirectionNothing
	}
	wallHit, err := board.UpdateBoard(playerDirection)
	c.SetCookie(&http.Cookie{
		Name:  "player_direction",
		Value: " ",
		Path:  "/",
	})
	// c.Logger().Debug()(win)
	if wallHit != "" {
		return Render(c, 200, pong.PongWinner(wallHit))
	}
	if err != nil {
		c.Logger().Error(err)
		return Render(c, 200, pong.StartGameButton())
	}
	h.DB.UpdateCachePongSession(c, board)

	return Render(c, 200, pong.LiveBoard(board.Grid))

}
