package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Action is a custom integer type to represent logical game actions.
type Action int

// ActionState represents the state of a logical game action.
type ActionState struct {
	pressed      bool
	justPressed  bool
	justReleased bool

	x float64
	y float64

	strength float64
}

// Input is the main manager for action-based input.
type Input struct {
	actions map[Action]*ActionState
}

// NewInput creates and initializes a new Input instance.
func NewInput() *Input {
	return &Input{
		actions: make(map[Action]*ActionState),
	}
}

// Update updates the state of all actions.
func (i *Input) Update() {
	// TODO: Iterate over bound devices and update action states.
}

// Pressed returns true if the action is currently pressed.
func (i *Input) Pressed(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.pressed
	}
	return false
}

// JustPressed returns true if the action was just pressed in the current frame.
func (i *Input) JustPressed(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.justPressed
	}
	return false
}

// JustReleased returns true if the action was just released in the current frame.
func (i *Input) JustReleased(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.justReleased
	}
	return false
}
