# ebitenpad 設計書 v1.0

============================================
1. 概要
============================================

ebitenpad は Ebitengine 向けの入力統合ライブラリである。

目的：

- デバイス非依存入力
- Actionベース抽象化
- Virtualパッド内蔵
- マルチタッチ対応
- 軽量ポーリング型
- シンプルで拡張可能

============================================
2. 設計思想
============================================

2.1 Action中心設計

`type Action int`

すべての入力は Action に正規化される。
ゲーム側はデバイスを意識しない。

2.2 正規化モデル

```
type ActionState struct {
  pressed      bool
  justPressed  bool
  justReleased bool

  x float64
  y float64

  strength float64
}
```

意味：

pressed        : 押下中
justPressed    : 今フレーム押下
justReleased   : 今フレーム解放
x,y            : -1.0 ～ 1.0 正規化ベクトル
strength       : 0.0 ～ 1.0 入力強度

============================================
3. ディレクトリ構造設計
============================================

`github.com/sofia-gros/ebitenpad`

ebitenpad/
  go.mod
  README.md
  DESIGN.md
  main.go            `サンプル実行用`
  input/
    input.go         `Input, ActionState, デバイス統合`
    keyboard.go      `キーボード処理`
    gamepad.go       `ゲームパッド処理`
  virtual/
    virtual.go       `Virtual UI 基本クラス`
    stick.go         `FixedStick, FreeStick, DPad`
    button.go        `Button`
  examples/
    basic.go         `使用例コード`

============================================
4. コアAPI仕様
============================================

4.1 Input生成

`input := ebitenpad.NewInput()`

4.2 フレーム更新

`input.Update()`
`input.Draw(screen)`

4.3 Action取得API

`Pressed(Action) bool`
`JustPressed(Action) bool`
`JustReleased(Action) bool`

`Vector(Action) (x, y float64)`
`Direction4(Action) int`
`Direction8(Action) int`
`Angle(Action) float64`
`Strength(Action) float64`

============================================
5. Direction仕様
============================================

5.1 Direction4

返却値：

-1 : 無入力
 0 : Up
 1 : Right
 2 : Down
 3 : Left

判定方式：

- abs(x) と abs(y) を比較
- 優勢軸を採用
- threshold未満は無入力

5.2 Direction8

返却値：

-1 : 無入力
 0～7 : 8方向（45度単位）

内部：

- atan2(y, x) により角度取得
- 45度単位で量子化

============================================
6. Virtual UI設計
============================================

6.1 取得方法

`ui := input.Virtual()`

6.2 Stick種類

1) FixedStick
   - 位置固定
   - 半径固定

2) FreeStick
   - タッチ位置が中心になる

3) DPad
   - 4方向専用

6.3 Stick共通仕様

- 二重円構造：中心円 + 指位置円
- 半径指定可能
- 方向量子化可能（4 / 8 / 16 / 360）
- `SetPosition(x, y float64)`
- `SetRadius(r float64)`
- `SetDirections(n int)`

============================================
7. Button仕様
============================================

生成：

`btn := ui.AddButton()`

設定：

`SetPosition(x, y)`
`SetRadius(r)`
`SetImage(path string)`

特性：

- 押下判定は円形
- マルチタッチ対応
- 押下IDロック方式

============================================
8. Actionバインド仕様
============================================

8.1 Virtual UI

`input.BindStick(action, stick)`
`input.BindButton(action, button)`

8.2 キーボード

`input.BindKey(action, ebiten.KeyX)`
`input.BindKeyAxis(action, left, right, up, down)`

8.3 ゲームパッド

`input.BindGamepadButton(action, buttonIndex)`
`input.BindGamepadAxis(action, axisX, axisY)`

============================================
9. マルチタッチ仕様
============================================

- タッチIDごとに管理
- 各Stickは最初に触れたIDをロック
- 他のStickやButtonと干渉しない
- 同時操作可能
- FreeStickは最初の接触点を中心とする

============================================
10. 使用例（完成形）
============================================

```
func (g *Game) Update() error {
  g.input.Update()

  if g.input.Pressed(Move) {
    x, y := g.input.Vector(Move)
    dir4 := g.input.Direction4(Move)
    dir8 := g.input.Direction8(Move)

    player.Move(x, y)
    player.SetDirection(dir8)
  }

  if g.input.JustPressed(Jump) {
    player.Jump()
  }

  return nil
}
```

その他の例：

- バーチャルスティック生成

`moveStick := ui.AddFixedStick().SetPosition(120, 300).SetRadius(80).SetDirections(8)`
`aimStick  := ui.AddFreeStick().SetRadius(70)`
`jumpBtn   := ui.AddButton().SetPosition(500,280).SetRadius(40).SetImage("jump.png")`

- キーボード/ゲームパッドバインド

`input.BindKey(Jump, ebiten.KeySpace)`
`input.BindGamepadButton(Jump, 0)`

============================================
11. 内部構造（概念）
============================================

`Input
 ├─ actions map[Action]*ActionState
 ├─ keyboard
 ├─ gamepad
 └─ virtualUI
        ├─ sticks
        └─ buttons`

ActionState:

`pressed, justPressed, justReleased, x, y, strength`

============================================
12. 設計の強み
============================================

- デバイス非依存
- スマホ・PC両対応
- 2D/ローグライク/シューティング対応
- 軽量
- 拡張容易
- 内部正規化構造が明確

============================================
13. 将来拡張可能項目
============================================

- 入力デッドゾーン設定
- カスタム感度調整
- JSON設定ロード
- 入力リプレイ記録
- ネットワーク同期入力