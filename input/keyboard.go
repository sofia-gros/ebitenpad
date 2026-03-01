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
