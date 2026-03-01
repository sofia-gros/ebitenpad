# 実装計画 - キーボードとゲームパッドの入力バインド機能の実装

このプランは、物理入力デバイスを論理アクションに紐付ける機能をカバーします。

## フェーズ 1: キーボードバインドの実装 [checkpoint: 0a2f82f]
- [x] タスク: `input/keyboard.go` を作成し、キーボード入力のバインド構造を定義 89c6239
- [x] タスク: `BindKey(action, key)` の実装 fc98917
- [x] タスク: `BindKeyAxis(action, left, right, up, down)` の実装 9e51889
- [x] タスク: `Input.Update()` におけるキーボード入力のポーリング処理の実装 c5444c9
- [x] タスク: キーボードバインドのユニットテスト作成 (`input/keyboard_test.go`) 9ec525e
- [x] タスク: Conductor - ユーザー手動検証 'キーボードバインド' (workflow.md 準拠) 0a2f82f

## フェーズ 2: ゲームパッドバインドの実装 [checkpoint: 179ef66]
- [x] タスク: `input/gamepad.go` を作成し、ゲームパッド入力のバインド構造を定義 71b6a1d
- [x] タスク: `BindGamepadButton(action, button)` の実装 11d3719
- [x] タスク: `BindGamepadAxis(action, axisX, axisY)` の実装 e405a41
- [x] タスク: `Input.Update()` におけるゲームパッド入力のポーリング処理の実装 b250b0c
- [x] タスク: ゲームパッドバインドのユニットテスト作成 (`input/gamepad_test.go`) 6e01c1d
- [x] タスク: Conductor - ユーザー手動検証 'ゲームパッドバインド' (workflow.md 準拠) 179ef66

## フェーズ 3: 複合入力の検証
- [ ] タスク: 複数の入力ソース（キーボード + ゲームパッド）の統合テスト
- [ ] タスク: `JustPressed` / `JustReleased` の状態維持ロジックの改善
- [ ] タスク: Conductor - ユーザー手動検証 '複合入力' (workflow.md 準拠)
