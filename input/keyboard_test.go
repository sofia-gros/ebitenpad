package input

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestBindKey(t *testing.T) {
	const jump Action = 1
	input := NewInput()

	input.BindKey(jump, ebiten.KeySpace)

	if len(input.keyboard.keys) != 1 {
		t.Errorf("BindKey が正しく実行されませんでした。期待値: 1, 実際: %d", len(input.keyboard.keys))
	}

	if input.keyboard.keys[0].action != jump || input.keyboard.keys[0].key != ebiten.KeySpace {
		t.Error("バインドされたアクションまたはキーが正しくありません")
	}
}

func TestBindKeyAxis(t *testing.T) {
	const move Action = 1
	input := NewInput()

	input.BindKeyAxis(move, ebiten.KeyA, ebiten.KeyD, ebiten.KeyW, ebiten.KeyS)

	if len(input.keyboard.axes) != 1 {
		t.Errorf("BindKeyAxis が正しく実行されませんでした。期待値: 1, 実際: %d", len(input.keyboard.axes))
	}

	axis := input.keyboard.axes[0]
	if axis.action != move || axis.left != ebiten.KeyA || axis.right != ebiten.KeyD || axis.up != ebiten.KeyW || axis.down != ebiten.KeyS {
		t.Error("バインドされたアクションまたは軸キーが正しくありません")
	}
}

// 実際のポーリングは Ebitengine の API に依存するため、Update のテストは難しいが
// 内部の状態管理ロジックのテストは別途検討が必要。
