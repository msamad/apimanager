package actions

import (
	"github.com/labstack/echo/v4"
	"sync"
)

type Action struct {
	ActionBean map[int64]ActionBean
	NextId int64
	Lock sync.Mutex
}

func (a Action) PostAction(ctx echo.Context) error {
	panic("implement me")
}

func NewAction() *Action {
	return &Action{
		ActionBean: make(map[int64]ActionBean),
		NextId: 1000,
	}
}


