package commands

import (
	"context"
	"errors"
	"reflect"
)

// ICommandHandler là interface cho các command handler
type ICommandHandler interface {
	Handle(ctx context.Context, command interface{}) error
}

// CommandBus quản lý việc gửi command tới các handler phù hợp
type CommandBus struct {
	handlers map[reflect.Type]ICommandHandler
}

// NewCommandBus tạo một CommandBus mới
func NewCommandBus() *CommandBus {
	return &CommandBus{handlers: make(map[reflect.Type]ICommandHandler)}
}

// RegisterHandler đăng ký một handler mới với command tương ứng
func (cb *CommandBus) RegisterHandler(commandType reflect.Type, handler ICommandHandler) {
	cb.handlers[commandType] = handler
}

// Dispatch gửi command tới handler phù hợp
func (cb *CommandBus) Dispatch(ctx context.Context, command interface{}) error {
	handler, ok := cb.handlers[reflect.TypeOf(command)]
	if !ok {
		return errors.New("no handler registered for command")
	}
	return handler.Handle(ctx, command)
}
