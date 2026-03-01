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

func TestInputQueries(t *testing.T) {
	const jump Action = 1
	input := NewInput()

	// Initially all queries should return false
	if input.Pressed(jump) {
		t.Error("Pressed() should be false initially")
	}
	if input.JustPressed(jump) {
		t.Error("JustPressed() should be false initially")
	}
	if input.JustReleased(jump) {
		t.Error("JustReleased() should be false initially")
	}

	// Mock state for the jump action
	input.actions[jump] = &ActionState{
		pressed:      true,
		justPressed:  true,
		justReleased: false,
	}

	// Verify query results with mocked state
	if !input.Pressed(jump) {
		t.Error("Pressed() should be true when state.pressed is true")
	}
	if !input.JustPressed(jump) {
		t.Error("JustPressed() should be true when state.justPressed is true")
	}
	if input.JustReleased(jump) {
		t.Error("JustReleased() should be false when state.justReleased is false")
	}

	// Mock another state
	input.actions[jump].pressed = false
	input.actions[jump].justPressed = false
	input.actions[jump].justReleased = true

	if input.Pressed(jump) {
		t.Error("Pressed() should be false when state.pressed is false")
	}
	if input.JustPressed(jump) {
		t.Error("JustPressed() should be false when state.justPressed is false")
	}
	if !input.JustReleased(jump) {
		t.Error("JustReleased() should be true when state.justReleased is true")
	}
}
