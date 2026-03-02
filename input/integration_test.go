package input

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

type mockScanner struct {
	pressedKeys    map[ebiten.Key]bool
	pressedButtons map[ebiten.StandardGamepadButton]bool
	gamepadAxes    map[int]float64
}

func (s *mockScanner) IsKeyPressed(key ebiten.Key) bool {
	return s.pressedKeys[key]
}

func (s *mockScanner) AppendGamepadIDs(ids []ebiten.GamepadID) []ebiten.GamepadID {
	return append(ids, 0)
}

func (s *mockScanner) IsStandardGamepadButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return s.pressedButtons[button]
}

func (s *mockScanner) StandardGamepadAxisValue(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64 {
	return s.gamepadAxes[int(axis)]
}

func TestMultipleInputSourcesIntegration(t *testing.T) {
	const jump Action = 1
	input := NewInput()

	mock := &mockScanner{
		pressedKeys:    make(map[ebiten.Key]bool),
		pressedButtons: make(map[ebiten.StandardGamepadButton]bool),
		gamepadAxes:    make(map[int]float64),
	}
	input.keyboardScanner = mock
	input.gamepadScanner = mock

	// バインド設定
	input.BindKey(jump, ebiten.KeySpace)
	input.BindGamepadButton(jump, ebiten.StandardGamepadButtonCenterLeft)

	// フレーム 1: キーボードのみ
	mock.pressedKeys[ebiten.KeySpace] = true
	input.Update()

	if !input.Pressed(jump) {
		t.Error("Frame 1: Action should be Pressed (Keyboard)")
	}
	if !input.JustPressed(jump) {
		t.Error("Frame 1: Action should be JustPressed")
	}

	// フレーム 2: キーボード継続 + ゲームパッド追加
	mock.pressedButtons[ebiten.StandardGamepadButtonCenterLeft] = true
	input.Update()

	if !input.Pressed(jump) {
		t.Error("Frame 2: Action should be Pressed (Both)")
	}
	if input.JustPressed(jump) {
		t.Error("Frame 2: Action should NOT be JustPressed (Already Pressed)")
	}

	// フレーム 3: キーボード離す + ゲームパッド継続
	mock.pressedKeys[ebiten.KeySpace] = false
	input.Update()

	if !input.Pressed(jump) {
		t.Error("Frame 3: Action should be Pressed (Gamepad)")
	}
	if input.JustReleased(jump) {
		t.Error("Frame 3: Action should NOT be JustReleased (Still Pressed by gamepad)")
	}

	// フレーム 4: 全て離す
	mock.pressedButtons[ebiten.StandardGamepadButtonCenterLeft] = false
	input.Update()

	if input.Pressed(jump) {
		t.Error("Frame 4: Action should be released")
	}
	if !input.JustReleased(jump) {
		t.Error("Frame 4: Action should be JustReleased")
	}
}
