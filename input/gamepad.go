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
func newGamepadManager() *keyboardManager {
	return &gamepadManager{
		buttons: []gamepadButtonBinding{},
		axes:    []gamepadAxisBinding{},
	}
}

// BindGamepadButton はゲームパッドのボタンをアクションにバインドします。
func (i *Input) BindGamepadButton(action Action, button ebiten.StandardGamepadButton) {
	i.gamepad.buttons = append(i.gamepad.buttons, gamepadButtonBinding{
		action: action,
		button: button,
	})
}

// BindGamepadAxis はゲームパッドのアナログスティック軸をアクションにバインドします。
func (i *Input) BindGamepadAxis(action Action, axisX, axisY int) {
	i.gamepad.axes = append(i.gamepad.axes, gamepadAxisBinding{
		action: action,
		axisX:  axisX,
		axisY:  axisY,
	})
}
