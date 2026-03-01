package input

import (
	"testing"
)

func TestNewInput(t *testing.T) {
	input := NewInput()
	if input == nil {
		t.Fatal("NewInput() が nil を返しました")
	}
	if input.actions == nil {
		t.Error("NewInput() が actions マップを初期化しませんでした")
	}
}

func TestActionStateInitialValues(t *testing.T) {
	state := ActionState{}
	if state.pressed != false {
		t.Error("ActionState.pressed のデフォルト値は false である必要があります")
	}
	if state.justPressed != false {
		t.Error("ActionState.justPressed のデフォルト値は false である必要があります")
	}
	if state.justReleased != false {
		t.Error("ActionState.justReleased のデフォルト値は false である必要があります")
	}
	if state.x != 0 {
		t.Errorf("ActionState.x のデフォルト値は 0 である必要があります。現在の値: %f", state.x)
	}
	if state.y != 0 {
		t.Errorf("ActionState.y のデフォルト値は 0 である必要があります。現在の値: %f", state.y)
	}
	if state.strength != 0 {
		t.Errorf("ActionState.strength のデフォルト値は 0 である必要があります。現在の値: %f", state.strength)
	}
}

func TestInputQueries(t *testing.T) {
	const jump Action = 1
	input := NewInput()

	// 初期状態ではすべてのクエリが false を返す必要があります
	if input.Pressed(jump) {
		t.Error("初期状態の Pressed() は false である必要があります")
	}
	if input.JustPressed(jump) {
		t.Error("初期状態の JustPressed() は false である必要があります")
	}
	if input.JustReleased(jump) {
		t.Error("初期状態の JustReleased() は false である必要があります")
	}

	// ジャンプアクションの状態をモックします
	input.actions[jump] = &ActionState{
		pressed:      true,
		justPressed:  true,
		justReleased: false,
	}

	// モックされた状態でのクエリ結果を確認します
	if !input.Pressed(jump) {
		t.Error("state.pressed が true の場合、Pressed() は true である必要があります")
	}
	if !input.JustPressed(jump) {
		t.Error("state.justPressed が true の場合、JustPressed() は true である必要があります")
	}
	if input.JustReleased(jump) {
		t.Error("state.justReleased が false の場合、JustReleased() は false である必要があります")
	}

	// 別の状態をモックします
	input.actions[jump].pressed = false
	input.actions[jump].justPressed = false
	input.actions[jump].justReleased = true

	if input.Pressed(jump) {
		t.Error("state.pressed が false の場合、Pressed() は false である必要があります")
	}
	if input.JustPressed(jump) {
		t.Error("state.justPressed が false の場合、JustPressed() は false である必要があります")
	}
	if !input.JustReleased(jump) {
		t.Error("state.justReleased が true の場合、JustReleased() は true である必要があります")
	}
}
