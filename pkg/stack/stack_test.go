package stack_test

import (
	"testing"

	"github.com/daronenko/calc/pkg/stack"
)

func TestStackPushPop(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if value, ok := s.Pop(); !ok || value != 3 {
		t.Errorf("got %d, expected 3", value)
	}

	if value, ok := s.Pop(); !ok || value != 2 {
		t.Errorf("got %d, expected 2", value)
	}

	if value, ok := s.Pop(); !ok || value != 1 {
		t.Errorf("got %d, expected 1", value)
	}

	if _, ok := s.Pop(); ok {
		t.Error("got true, expected false")
	}
}

func TestStackTop(t *testing.T) {
	s := stack.New[string]()

	if _, ok := s.Top(); ok {
		t.Error("got true, expected false")
	}

	s.Push("first")
	if value, ok := s.Top(); !ok || *value != "first" {
		t.Errorf("got '%s', expected 'first'", *value)
	}
	s.Push("second")
	if value, ok := s.Top(); !ok || *value != "second" {
		t.Errorf("got '%s', expected 'second'", *value)
	}

	s.Pop()
	if value, ok := s.Top(); !ok || *value != "first" {
		t.Errorf("got '%s', expected 'first'", *value)
	}
}

func TestStackSize(t *testing.T) {
	s := stack.New[float64]()

	if size := s.Len(); size != 0 {
		t.Errorf("got %d, expected size 0", size)
	}

	s.Push(1.1)
	s.Push(2.2)
	if size := s.Len(); size != 2 {
		t.Errorf("got %d, expected size 2", size)
	}

	s.Pop()
	if size := s.Len(); size != 1 {
		t.Errorf("got %d, expected size 1", size)
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := stack.New[bool]()

	if !s.IsEmpty() {
		t.Error("got false, expected true")
	}

	s.Push(true)
	if s.IsEmpty() {
		t.Error("got true, expected false")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Error("got false, expected true")
	}
}
