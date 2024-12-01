package georedis

import (
	"context"
	ponggame "geocity/pong"
)

// Checks for existing session, and will overwirte with a new id if found, otherwise just creates a new id
func (gr *GeoDB) CreatePongSession(ctx context.Context, id string, board ponggame.BoardState) (string, error) {
	gr.client.Del(ctx, id)
	set := gr.client.Set(ctx, id, board.FlattenBoard(), 0)
	return set.Result()
}

func (gr *GeoDB) DeletePongSession(ctx context.Context, existingString string) error {
	cmd := gr.client.Del(ctx, existingString)
	return cmd.Err()
}

func (gr *GeoDB) GetBoard(ctx context.Context, gameId string) (ponggame.BoardState, error) {
	cmd := gr.client.Get(ctx, gameId)
	res, err := cmd.Result()
	if err != nil {
		return ponggame.BoardState{}, err
	}
	return ponggame.NewBoardFromFlattenedBoard(res), nil
}
