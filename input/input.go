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
