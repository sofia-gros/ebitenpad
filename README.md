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
	"github.com/sofia-gros/ebitenpad/input"
)

// アクションの定義
const (
	ActionMove input.Action = iota
	ActionJump
)

type Game struct {
	in *input.Input
}

func NewGame() *Game {
	in := input.NewInput()

	// キーボードのバインド
	in.BindKeyAxis(ActionMove, ebiten.KeyA, ebiten.KeyD, ebiten.KeyW, ebiten.KeyS)
	in.BindKey(ActionJump, ebiten.KeySpace)

    // ゲームパッドのバインド
	in.BindGamepadButton(ActionJump, ebiten.StandardGamepadButtonRightBottom)
	in.BindGamepadAxis(ActionMove, 0, 1)

	// バーチャルパッドの設定
	v := in.Virtual()
	stick := v.AddFixedStick(100, 300, 60)
	btn := v.AddButton(540, 300, 40)

	// バーチャルパッドのバインド
	in.BindStick(ActionMove, stick)
	in.BindButton(ActionJump, btn)

	return &Game{in: in}
}

func (g *Game) Update() error {
	// 毎フレーム更新
	g.in.Update()

	// アクションの状態を確認
	if g.in.Pressed(ActionMove) {
		// 移動ベクトルの取得 (-1.0 ~ 1.0)
		state, _ := g.in.GetActionState(ActionMove)
		dx, dy := state.X, state.Y
		// プレイヤーの移動処理など
	}

	if g.in.JustPressed(ActionJump) {
		// ジャンプ処理
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// バーチャルパッドの描画
	g.in.Virtual().Draw(screen)
}

func main() {
	ebiten.RunGame(NewGame())
}
```

## ライブラリの構成

- `input`: メインの入力マネージャー。アクションのバインドと状態管理を行います。
- `virtual`: バーチャルスティックやボタンなどのUIコンポーネント。
- `examples/wasm`: WebAssembly で動作するデモコード。

## ライセンス

MIT License
