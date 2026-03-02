package input

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sofia-gros/ebitenpad/virtual"
)

// Action は論理的なゲームアクションを表すカスタム整数型です。
type Action int

// ActionState は論理的なゲームアクションの状態を表します。
type ActionState struct {
	Pressed      bool
	JustPressed  bool
	JustReleased bool

	X float64
	Y float64

	Strength float64

	lastPressed bool
}

// GetActionState は指定されたアクションの状態のコピーを返します。
func (i *Input) GetActionState(action Action) (ActionState, bool) {
	if state, ok := i.actions[action]; ok {
		return *state, true
	}
	return ActionState{}, false
}

// KeyboardScanner はキーボードの状態をスキャンするインターフェースです。
type KeyboardScanner interface {
	IsKeyPressed(key ebiten.Key) bool
}

// GamepadScanner はゲームパッドの状態をスキャンするインターフェースです。
type GamepadScanner interface {
	AppendGamepadIDs(ids []ebiten.GamepadID) []ebiten.GamepadID
	IsStandardGamepadButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool
	StandardGamepadAxisValue(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64
}

// DefaultKeyboardScanner は ebiten の標準 API を使用するデフォルトのスキャナーです。
type DefaultKeyboardScanner struct{}

func (s *DefaultKeyboardScanner) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

// DefaultGamepadScanner は ebiten の標準 API を使用するデフォルトのスキャナーです。
type DefaultGamepadScanner struct{}

func (s *DefaultGamepadScanner) AppendGamepadIDs(ids []ebiten.GamepadID) []ebiten.GamepadID {
	return ebiten.AppendGamepadIDs(ids)
}

func (s *DefaultGamepadScanner) IsStandardGamepadButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return ebiten.IsStandardGamepadButtonPressed(id, button)
}

func (s *DefaultGamepadScanner) StandardGamepadAxisValue(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) float64 {
	return ebiten.StandardGamepadAxisValue(id, axis)
}

type virtualButtonBinding struct {
	action Action
	button *virtual.Button
}

type virtualStickBinding struct {
	action Action
	stick  *virtual.Stick
}

// Input はアクションベースの入力を管理するメインマネージャーです。
type Input struct {
	actions  map[Action]*ActionState
	keyboard *keyboardManager
	gamepad  *gamepadManager
	virtual  *virtual.VirtualPad

	virtualButtons []virtualButtonBinding
	virtualSticks  []virtualStickBinding

	keyboardScanner KeyboardScanner
	gamepadScanner  GamepadScanner
}

// NewInput は新しい Input インスタンスを作成し、初期化します。
func NewInput() *Input {
	return &Input{
		actions:         make(map[Action]*ActionState),
		keyboard:        newKeyboardManager(),
		gamepad:         newGamepadManager(),
		virtual:         virtual.NewVirtualPad(),
		virtualButtons:  []virtualButtonBinding{},
		virtualSticks:   []virtualStickBinding{},
		keyboardScanner: &DefaultKeyboardScanner{},
		gamepadScanner:  &DefaultGamepadScanner{},
	}
}

// BindButton はバーチャルボタンをアクションにバインドします。
func (i *Input) BindButton(action Action, button *virtual.Button) {
	i.virtualButtons = append(i.virtualButtons, virtualButtonBinding{
		action: action,
		button: button,
	})
}

// BindStick はバーチャルスティックをアクションにバインドします。
func (i *Input) BindStick(action Action, stick *virtual.Stick) {
	i.virtualSticks = append(i.virtualSticks, virtualStickBinding{
		action: action,
		stick:  stick,
	})
}

// Virtual はバーチャル UI マネージャーを返します。
func (i *Input) Virtual() *virtual.VirtualPad {
	return i.virtual
}

// Update はすべてのアクションの状態を更新します。
func (i *Input) Update() {
	// 状態のリセット
	for _, state := range i.actions {
		lastPressed := state.Pressed
		state.Pressed = false
		state.JustPressed = false
		state.JustReleased = false
		state.X = 0
		state.Y = 0
		state.Strength = 0

		state.lastPressed = lastPressed
	}

	// 各デバイスのポーリング
	i.keyboard.update(i.actions, i.keyboardScanner)
	i.gamepad.update(i.actions, i.gamepadScanner)
	i.updateVirtual()

	// JustPressed / JustReleased の確定
	for _, state := range i.actions {
		if state.Pressed && !state.lastPressed {
			state.JustPressed = true
		}
		if !state.Pressed && state.lastPressed {
			state.JustReleased = true
		}
	}
}

func (i *Input) updateVirtual() {
	i.virtual.Update()

	for _, b := range i.virtualButtons {
		state := getOrInitState(i.actions, b.action)
		if b.button.Pressed() {
			state.Pressed = true
			state.Strength = 1.0
		}
	}

	for _, s := range i.virtualSticks {
		state := getOrInitState(i.actions, s.action)
		vx, vy := s.stick.Vector()
		if vx != 0 || vy != 0 {
			state.Pressed = true
			// 軸合成
			if math.Abs(vx) > math.Abs(state.X) {
				state.X = vx
			}
			if math.Abs(vy) > math.Abs(state.Y) {
				state.Y = vy
			}
			if s.stick.Strength() > state.Strength {
				state.Strength = s.stick.Strength()
			}
		}
	}
}

// Pressed はアクションが現在押されている場合に true を返します。
func (i *Input) Pressed(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.Pressed
	}
	return false
}

// JustPressed は現在のフレームでアクションが押されたばかりの場合に true を返します。
func (i *Input) JustPressed(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.JustPressed
	}
	return false
}

// JustReleased は現在のフレームでアクションが離されたばかりの場合に true を返します。
func (i *Input) JustReleased(action Action) bool {
	if state, ok := i.actions[action]; ok {
		return state.JustReleased
	}
	return false
}
