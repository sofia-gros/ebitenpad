package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// keyBinding はアクションにバインドされた単一のキーを表します。
type keyBinding struct {
	action Action
	key    ebiten.Key
}

// keyAxisBinding はアクションにバインドされた4方向のキーを表します。
type keyAxisBinding struct {
	action Action
	left   ebiten.Key
	right  ebiten.Key
	up     ebiten.Key
	down   ebiten.Key
}

// keyboardManager はキーボード入力を管理します。
type keyboardManager struct {
	keys   []keyBinding
	axes   []keyAxisBinding
}

// newKeyboardManager は新しい keyboardManager を作成します。
func newKeyboardManager() *keyboardManager {
	return &keyboardManager{
		keys: []keyBinding{},
		axes: []keyAxisBinding{},
	}
}

// BindKey は単一のキーをアクションにバインドします。
func (i *Input) BindKey(action Action, key ebiten.Key) {
	i.keyboard.keys = append(i.keyboard.keys, keyBinding{
		action: action,
		key:    key,
	})
}

// BindKeyAxis は4つのキーをベクトルアクションとしてバインドします。
func (i *Input) BindKeyAxis(action Action, left, right, up, down ebiten.Key) {
	i.keyboard.axes = append(i.keyboard.axes, keyAxisBinding{
		action: action,
		left:   left,
		right:  right,
		up:     up,
		down:   down,
	})
}

// update はキーボード入力をポーリングし、各アクションの状態を更新します。
func (m *keyboardManager) update(actions map[Action]*ActionState) {
	for _, b := range m.keys {
		state := getOrInitState(actions, b.action)
		if ebiten.IsKeyPressed(b.key) {
			state.pressed = true
			state.strength = 1.0
		}
	}

	for _, b := range m.axes {
		state := getOrInitState(actions, b.action)
		var dx, dy float64
		if ebiten.IsKeyPressed(b.left) {
			dx -= 1.0
		}
		if ebiten.IsKeyPressed(b.right) {
			dx += 1.0
		}
		if ebiten.IsKeyPressed(b.up) {
			dy -= 1.0
		}
		if ebiten.IsKeyPressed(b.down) {
			dy += 1.0
		}

		if dx != 0 || dy != 0 {
			state.pressed = true
			state.x = dx
			state.y = dy
			state.strength = 1.0 // 簡易実装。
		}
	}
}

func getOrInitState(actions map[Action]*ActionState, action Action) *ActionState {
	state, ok := actions[action]
	if !ok {
		state = &ActionState{}
		actions[action] = state
	}
	return state
}
