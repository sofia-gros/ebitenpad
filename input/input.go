package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Action は論理的なゲームアクションを表すカスタム整数型です。
type Action int

// ActionState は論理的なゲームアクションの状態を表します。
type ActionState struct {
	pressed      bool
	justPressed  bool
	justReleased bool

	x float64
	y float64

	strength float64

	lastPressed bool
}

// KeyboardScanner はキーボードの状態をスキャンするインターフェースです。
type KeyboardScanner interface {
	IsKeyPressed(key ebiten.Key) bool
}

// GamepadScanner はゲームパッドの状態をスキャンするインターフェースです。
type GamepadScanner interface {
	AppendGamepadIDs(ids []ebiten.GamepadID) []ebiten.GamepadID
	IsStandardGamepadButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool
	StandardGamepadAxisValue(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64
}

// DefaultKeyboardScanner は ebiten の標準 API を使用するデフォルトのスキャナーです。
type DefaultKeyboardScanner struct{}

func (s *DefaultKeyboardScanner) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

// DefaultGamepadScanner は ebiten の標準 API を使用するデフォルトのスキャナーです。
type DefaultGamepadScanner struct{}

func (s *DefaultGamepadScanner) AppendGamepadIDs(ids []ebiten.GamepadID) []ebiten.GamepadID {
	return ebiten.AppendGamepadIDs(ids)
}

func (s *DefaultGamepadScanner) IsStandardGamepadButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return ebiten.IsStandardGamepadButtonPressed(id, button)
}

func (s *DefaultGamepadScanner) StandardGamepadAxisValue(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64 {
	return ebiten.StandardGamepadAxisValue(id, axis)
}

// Input はアクションベースの入力を管理するメインマネージャーです。
type Input struct {
	actions  map[Action]*ActionState
	keyboard *keyboardManager
	gamepad  *gamepadManager

	keyboardScanner KeyboardScanner
	gamepadScanner  GamepadScanner
}

// NewInput は新しい Input インスタンスを作成し、初期化します。
func NewInput() *Input {
	return &Input{
		actions:         make(map[Action]*ActionState),
		keyboard:        newKeyboardManager(),
		gamepad:         newGamepadManager(),
		keyboardScanner: &DefaultKeyboardScanner{},
		gamepadScanner:  &DefaultGamepadScanner{},
	}
}

// Update はすべてのアクションの状態を更新します。
func (i *Input) Update() {
	// 状態のリセット
	for _, state := range i.actions {
		lastPressed := state.pressed
		state.pressed = false
		state.justPressed = false
		state.justReleased = false
		state.x = 0
		state.y = 0
		state.strength = 0

		state.lastPressed = lastPressed
	}

	// 各デバイスのポーリング
	i.keyboard.update(i.actions, i.keyboardScanner)
	i.gamepad.update(i.actions, i.gamepadScanner)

	// JustPressed / JustReleased の確定
	for _, state := range i.actions {
		if state.pressed && !state.lastPressed {
			state.justPressed = true
		}
		if !state.pressed && state.lastPressed {
			state.justReleased = true
		}
	}
}

// Pressed はアクションが現在押されている場合に true を返します。
func (i *Input) Pressed(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.pressed
	}
	return false
}

// JustPressed は現在のフレームでアクションが押されたばかりの場合に true を返します。
func (i *Input) JustPressed(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.justPressed
	}
	return false
}

// JustReleased は現在のフレームでアクションが離されたばかりの場合に true を返します。
func (i *Input) JustReleased(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.justReleased
	}
	return false
}
