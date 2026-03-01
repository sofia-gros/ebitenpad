package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// gamepadButtonBinding はアクションにバインドされたゲームパッドのボタンを表します。
type gamepadButtonBinding struct {
	action Action
	button ebiten.StandardGamepadButton
}

// gamepadAxisBinding はアクションにバインドされたゲームパッドの軸を表します。
type gamepadAxisBinding struct {
	action Action
	axisX  int
	axisY  int
}

// gamepadManager はゲームパッド入力を管理します。
type gamepadManager struct {
	buttons []gamepadButtonBinding
	axes    []gamepadAxisBinding
}

// newGamepadManager は新しい gamepadManager を作成します。
func newGamepadManager() *gamepadManager {
	return &gamepadManager{
		buttons: []gamepadButtonBinding{},
		axes:    []gamepadAxisBinding{},
	}
}
