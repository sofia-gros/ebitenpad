package input

import (
	"math"

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

// update はゲームパッド入力をポーリングし、各アクションの状態を更新します。
func (m *gamepadManager) update(actions map[Action]*ActionState, scanner GamepadScanner) {
	ids := scanner.AppendGamepadIDs(nil)
	if len(ids) == 0 {
		return
	}
	// 簡易化のため、最初のゲームパッドのみを対象とします。
	id := ids[0]

	for _, b := range m.buttons {
		state := getOrInitState(actions, b.action)
		if scanner.IsStandardGamepadButtonPressed(id, b.button) {
			state.pressed = true
			state.strength = 1.0
		}
	}

	for _, b := range m.axes {
		state := getOrInitState(actions, b.action)
		x := scanner.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxis(b.axisX))
		y := scanner.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxis(b.axisY))

		// デッドゾーン処理などは将来の拡張として、ここでは生値を採用
		if x != 0 || y != 0 {
			state.pressed = true
			// すでに入力がある場合は合成する（簡易的に大きい方を採用）
			if math.Abs(x) > math.Abs(state.x) {
				state.x = x
			}
			if math.Abs(y) > math.Abs(state.y) {
				state.y = y
			}
			// 入力の強さを計算
			state.strength = 1.0 // 簡易実装
		}
	}
}
