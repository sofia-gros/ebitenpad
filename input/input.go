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
