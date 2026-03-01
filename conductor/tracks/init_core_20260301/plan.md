# Implementation Plan - Initialize project structure and core Action system

This plan covers the initial scaffolding and core type implementation for `ebitenpad`.

## Phase 1: Project Scaffolding
- [ ] Task: Create directory structure (`input`, `virtual`, `examples`)
- [ ] Task: Initialize `input/input.go` with package declaration and basic imports
- [ ] Task: Conductor - User Manual Verification 'Project Scaffolding' (Protocol in workflow.md)

## Phase 2: Core Types Implementation
- [ ] Task: Define `Action` and `ActionState` types in `input/input.go`
    - [ ] Define `type Action int`
    - [ ] Define `type ActionState struct` with fields: `pressed`, `justPressed`, `justReleased`, `x`, `y`, `strength`
- [ ] Task: Implement `Input` struct and `NewInput` constructor in `input/input.go`
    - [ ] `Input` should contain a map or slice of `ActionState`
- [ ] Task: Write unit tests for `ActionState` and `Input` initialization in `input/input_test.go`
- [ ] Task: Conductor - User Manual Verification 'Core Types Implementation' (Protocol in workflow.md)

## Phase 3: Action Query API
- [ ] Task: Implement `Input.Update()` method skeleton in `input/input.go`
- [ ] Task: Implement query methods in `input/input.go`:
    - [ ] `Pressed(Action) bool`
    - [ ] `JustPressed(Action) bool`
    - [ ] `JustReleased(Action) bool`
- [ ] Task: Write unit tests for `Input` query methods in `input/input_test.go`
- [ ] Task: Conductor - User Manual Verification 'Action Query API' (Protocol in workflow.md)
