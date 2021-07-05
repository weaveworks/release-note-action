package action

import "fmt"

type ActionFunc = func(ctx *Context) error

type Action interface {
	Do() error
}

func New(f ActionFunc) (Action, error) {
	if f == nil {
		return nil, FuncRequired
	}

	ctx, err := NewContextFromEnv()
	if err != nil {
		return nil, fmt.Errorf("creating context: %w", err)
	}

	return &actionImpl{
		ctx:       ctx,
		actionFun: f,
	}, nil
}

type actionImpl struct {
	ctx       *Context
	actionFun ActionFunc
}

func (a *actionImpl) Do() error {
	if err := a.actionFun(a.ctx); err != nil {
		return fmt.Errorf("executing action function: %w", err)
	}

	return nil
}
