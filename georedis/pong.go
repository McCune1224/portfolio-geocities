package georedis

import (
	"context"
	"errors"
	ponggame "geocity/pong"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	sessionDuration = time.Hour * 2
)

// Checks for existing session, and will overwirte with a new id if found, otherwise just creates a new id
func (gr *GeoDB) CreatePongSession(ctx context.Context, id string, board ponggame.BoardState) (string, error) {
	gr.client.Del(ctx, id)
	set := gr.client.Set(ctx, id, board.FlattenBoard(), sessionDuration)
	return set.Result()
}

func (gr *GeoDB) DeletePongSession(ctx context.Context, existingString string) error {
	cmd := gr.client.Del(ctx, existingString)
	return cmd.Err()
}

func (gr *GeoDB) UpdatePongSession(ctx context.Context, id string, board ponggame.BoardState) (string, error) {
	set := gr.client.Set(ctx, id, board.FlattenBoard(), sessionDuration)
	return set.Result()
}

func (gr *GeoDB) GetBoard(ctx context.Context, gameId string) (ponggame.BoardState, error) {
	cmd := gr.client.Get(ctx, gameId)
	res, err := cmd.Result()
	if err != nil {
		return ponggame.BoardState{}, err
	}
	return ponggame.NewBoardFromFlattenedBoard(res), nil
}

// Checks for existing session, and will overwirte with a new id if found, otherwise just creates a new id
func (gr *GeoDB) CreateCachePongSession(ctx echo.Context, board ponggame.BoardState) string {
	cookies := ctx.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "game" {
			gr.client.Del(ctx.Request().Context(), cookie.Value)
		}
	}
	boardState := board.FlattenBoard()
	newGameCookie := &http.Cookie{
		Name:    "game",
		Value:   boardState,
		Expires: time.Now().Add(sessionDuration),
	}
	ctx.SetCookie(newGameCookie)
	return boardState
}

func (gr *GeoDB) DeleteCachePongSession(ctx echo.Context) {
	cookies := ctx.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "game" {
			gr.client.Del(ctx.Request().Context(), cookie.Value)
		}
	}
}

func (gr *GeoDB) UpdateCachePongSession(ctx echo.Context, board ponggame.BoardState) string {
	cookies := ctx.Cookies()
	found := false
	for _, cookie := range cookies {
		if cookie.Name == "game" {
			found = true
			break
		}
	}
	if !found {
		return gr.CreateCachePongSession(ctx, board)
	}

	boardState := board.FlattenBoard()
	ctx.SetCookie(&http.Cookie{
		Name:    "game",
		Value:   boardState,
		Expires: time.Now().Add(sessionDuration),
	})
	return boardState
}

func (gr *GeoDB) GetCacheBoard(ctx echo.Context) (ponggame.BoardState, error) {
	cookies := ctx.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "game" {
			return ponggame.NewBoardFromFlattenedBoard(cookie.Value), nil
		}
	}
	return ponggame.BoardState{}, errors.New("No game cookie found")
}
