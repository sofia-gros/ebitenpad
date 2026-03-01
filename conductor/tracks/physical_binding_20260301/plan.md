# 実装計画 - キーボードとゲームパッドの入力バインド機能の実装

このプランは、物理入力デバイスを論理アクションに紐付ける機能をカバーします。

## フェーズ 1: キーボードバインドの実装
- [ ] タスク: `input/keyboard.go` を作成し、キーボード入力のバインド構造を定義
- [ ] タスク: `BindKey(action, key)` の実装
- [ ] タスク: `BindKeyAxis(action, left, right, up, down)` の実装
- [ ] タスク: `Input.Update()` におけるキーボード入力のポーリング処理の実装
- [ ] タスク: キーボードバインドのユニットテスト作成 (`input/keyboard_test.go`)
- [ ] タスク: Conductor - ユーザー手動検証 'キーボードバインド' (workflow.md 準拠)

## フェーズ 2: ゲームパッドバインドの実装
- [ ] タスク: `input/gamepad.go` を作成し、ゲームパッド入力のバインド構造を定義
- [ ] タスク: `BindGamepadButton(action, button)` の実装
- [ ] タスク: `BindGamepadAxis(action, axisX, axisY)` の実装
- [ ] タスク: `Input.Update()` におけるゲームパッド入力のポーリング処理の実装
- [ ] タスク: ゲームパッドバインドのユニットテスト作成 (`input/gamepad_test.go`)
- [ ] タスク: Conductor - ユーザー手動検証 'ゲームパッドバインド' (workflow.md 準拠)

## フェーズ 3: 複合入力の検証
- [ ] タスク: 複数の入力ソース（キーボード + ゲームパッド）の統合テスト
- [ ] タスク: `JustPressed` / `JustReleased` の状態維持ロジックの改善
- [ ] タスク: Conductor - ユーザー手動検証 '複合入力' (workflow.md 準拠)
