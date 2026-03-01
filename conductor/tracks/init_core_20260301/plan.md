# Implementation Plan - Initialize project structure and core Action system

This plan covers the initial scaffolding and core type implementation for `ebitenpad`.

## Phase 1: Project Scaffolding [checkpoint: 569f484]
- [x] Task: Create directory structure (`input`, `virtual`, `examples`) bd338cf
- [x] Task: Initialize `input/input.go` with package declaration and basic imports c9662ff
- [x] Task: Conductor - User Manual Verification 'Project Scaffolding' (Protocol in workflow.md) 569f484

## Phase 2: Core Types Implementation [checkpoint: f6afbac]
- [x] Task: Define `Action` and `ActionState` types in `input/input.go` a722933
    - [x] Define `type Action int`
    - [x] Define `type ActionState struct` with fields: `pressed`, `justPressed`, `justReleased`, `x`, `y`, `strength`
- [x] Task: Implement `Input` struct and `NewInput` constructor in `input/input.go` abf6a0b
    - [x] `Input` should contain a map or slice of `ActionState`
- [x] Task: Write unit tests for `ActionState` and `Input` initialization in `input/input_test.go` 07c119f
- [x] Task: Conductor - User Manual Verification 'Core Types Implementation' (Protocol in workflow.md) f6afbac

## Phase 3: Action Query API [checkpoint: ed62dee]
- [x] Task: Implement `Input.Update()` method skeleton in `input/input.go` f6e9f93
- [x] Task: Implement query methods in `input/input.go`: 800ff4f
    - [x] `Pressed(Action) bool`
    - [x] `JustPressed(Action) bool`
    - [x] `JustReleased(Action) bool`
- [x] Task: Write unit tests for `Input` query methods in `input/input_test.go` c6b9e6d
- [x] Task: Conductor - User Manual Verification 'Action Query API' (Protocol in workflow.md) ed62dee
