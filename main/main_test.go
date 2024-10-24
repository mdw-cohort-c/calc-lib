package main

import (
	"errors"
	"testing"
)

func TestHandler_WrongNumberOfArguments(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	if !errors.Is(err, errWrongNumberOfArgs) {
		t.Error("wrong error")
	}
}
func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "3"})
	if !errors.Is(err, errInvalidArg) {
		t.Error("wrong error")
	}
}
