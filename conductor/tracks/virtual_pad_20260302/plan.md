# 実装計画 - バーチャルスティックとボタンの実装

## フェーズ 1: バーチャル UI の基盤実装
- [~] タスク: `virtual/virtual.go` を作成し、基本構造を定義
- [ ] タスク: `virtual/button.go` の実装（タッチ ID ロックロジック含む）
- [ ] タスク: `virtual/stick.go` の実装（ベクトル正規化ロジック含む）
- [ ] タスク: ユニットテスト作成 (`virtual/button_test.go`, `virtual/stick_test.go`)
- [ ] タスク: Conductor - ユーザー手動検証 'Virtual UI 基本'

## フェーズ 2: 入力システムへの統合
- [ ] タスク: `input/input.go` に `VirtualPad` の統合
- [ ] タスク: `BindButton`, `BindStick` API の実装
- [ ] タスク: `Input.Update()` でのバーチャル入力ポーリングの実装
- [ ] タスク: 統合テスト作成
- [ ] タスク: Conductor - ユーザー手動検証 'Virtual 統合'

## フェーズ 3: WASM 例への反映
- [ ] タスク: `examples/wasm/main.go` にバーチャル UI を追加
- [ ] タスク: デバッグ描画機能の追加
- [ ] タスク: Conductor - ユーザー手動検証 'WASM バーチャル操作'
