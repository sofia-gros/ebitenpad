# 仕様書: キーボードとゲームパッドの入力バインド機能の実装

## 概要
このトラックでは、`Action` を物理デバイスの入力（キーボード、ゲームパッド）にバインドする機能を実装します。これにより、`Input.Update()` が各フレームの物理的な入力をスキャンし、`ActionState` を自動的に更新できるようになります。

## ユーザーストーリー
- 開発者として、特定のキーやゲームパッドのボタンを論理的な「アクション」に割り当てたい。
- 開発者として、1つのアクションに対して複数の入力ソース（例: スペースキーとゲームパッドの A ボタンの両方で「ジャンプ」）を割り当てたい。

## 要件

### 1. キーボードバインド API (`input/keyboard.go`)
- `BindKey(action Action, key ebiten.Key)`: 単一のキーをアクションにバインドします。
- `BindKeyAxis(action Action, left, right, up, down ebiten.Key)`: 4つのキーをベクトルアクション（スティック入力など）としてバインドします。

### 2. ゲームパッドバインド API (`input/gamepad.go`)
- `BindGamepadButton(action Action, button ebiten.StandardGamepadButton)`: 標準ゲームパッドボタンをアクションにバインドします。
- `BindGamepadAxis(action Action, axisX, axisY int)`: ゲームパッドのアナログスティック軸をアクションにバインドします。

### 3. 入力更新ロジック (`input/input.go`)
- `Input.Update()`:
  - 内部でキーボードとゲームパッドの入力をチェックし、対応する `ActionState` の `pressed`, `justPressed`, `justReleased`, `x`, `y`, `strength` を更新します。
  - 同じアクションに複数のバインドがある場合、いずれかが有効であれば `pressed` を true にします（ベクトルの場合は合成または優先順位に従う）。

## 受入条件
- [ ] キーボードの `BindKey` で登録したアクションが `Pressed` で正しく判定されること。
- [ ] ゲームパッドのボタン入力がアクションとして判定されること。
- [ ] `Update()` フレームごとに `JustPressed` / `JustReleased` が正しく切り替わること。
- [ ] 全てのコードコメントとドキュメントが日本語で記述されていること。
