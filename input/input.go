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
}

// Input はアクションベースの入力を管理するメインマネージャーです。
type Input struct {
	actions  map[Action]*ActionState
	keyboard *keyboardManager
}

// NewInput は新しい Input インスタンスを作成し、初期化します。
func NewInput() *Input {
	return &Input{
		actions:  make(map[Action]*ActionState),
		keyboard: newKeyboardManager(),
	}
}

// Update はすべてのアクションの状態を更新します。
func (i *Input) Update() {
	// TODO: バインドされたデバイスを反復処理し、アクションの状態を更新します。
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
