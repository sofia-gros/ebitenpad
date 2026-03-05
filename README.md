# ebitenpad

ebitenpad は Ebitengine (ebiten) 向けの、アクションベースの入力管理ライブラリです。
キーボード、ゲームパッド、そしてバーチャルパッド（タッチ操作）を一つの論理的な「アクション」に統合して扱うことができます。

## 特徴

- **アクションベースの抽象化**: ゲームロジックをデバイス（キーボード、ゲームパッド、タッチ）から切り離します。
- **バーチャルパッド内蔵**: スティックやボタンなどのUIを簡単に作成し、マルチタッチで操作可能です。
- **シンプルなAPI**: `Pressed`, `JustPressed`, `JustReleased` などの直感的なメソッドを提供します。
- **デバイス統合**: 1つのアクションに対して複数のデバイス入力をバインドできます。

## インストール

```bash
go get github.com/sofia-gros/ebitenpad
```

## 基本的な使い方

```go
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sofia-gros/ebitenpad/input"
)

const (
	ActionJump input.Action = 1
	ActionMove input.Action = 2
)

type Game struct {
	in *input.Input
}

func (g *Game) Update() error {
	g.in.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vx, vy := 0.0, 0.0
	strength := 0.0
	if state, ok := g.in.GetActionState(ActionMove); ok {
		vx, vy = state.X, state.Y
		strength = state.Strength
	}

	// バーチャル UI の描画
	g.in.Virtual().Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	in := input.NewInput()

	// Bind Keyboard
	in.BindKey(ActionJump, ebiten.KeySpace)
	in.BindKeyAxis(ActionMove, ebiten.KeyA, ebiten.KeyD, ebiten.KeyW, ebiten.KeyS)

	// Bind Gamepad
	in.BindGamepadButton(ActionJump, ebiten.StandardGamepadButtonRightBottom)
	in.BindGamepadAxis(ActionMove, 0, 1)

	// Virtual Padの設定
	vpad := in.Virtual()
	jumpBtn := vpad.AddButton().SetPosition(550, 400).SetRadius(40)
	moveStick := vpad.AddStick().SetPosition(100, 380).SetRadius(60)

	in.BindButton(ActionJump, jumpBtn)
	in.BindStick(ActionMove, moveStick)

	g := &Game{in: in}

	ebiten.SetWindowTitle("ebitenpad WASM Example")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
```

## ライブラリの構成

- `input`: メインの入力マネージャー。アクションのバインドと状態管理を行います。
- `virtual`: バーチャルスティックやボタンなどのUIコンポーネント。
- `examples/wasm`: WebAssembly で動作するデモコード。

## ライセンス

MIT License
