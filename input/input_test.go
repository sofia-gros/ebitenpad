package input

import (
	"testing"
)

func TestNewInput(t *testing.T) {
	input := NewInput()
	if input == nil {
		t.Fatal("NewInput() returned nil")
	}
	if input.actions == nil {
		t.Error("NewInput() did not initialize actions map")
	}
}

func TestActionStateInitialValues(t *testing.T) {
	state := ActionState{}
	if state.pressed != false {
		t.Error("ActionState.pressed should be false by default")
	}
	if state.justPressed != false {
		t.Error("ActionState.justPressed should be false by default")
	}
	if state.justReleased != false {
		t.Error("ActionState.justReleased should be false by default")
	}
	if state.x != 0 {
		t.Errorf("ActionState.x should be 0, got %f", state.x)
	}
	if state.y != 0 {
		t.Errorf("ActionState.y should be 0, got %f", state.y)
	}
	if state.strength != 0 {
		t.Errorf("ActionState.strength should be 0, got %f", state.strength)
	}
}
