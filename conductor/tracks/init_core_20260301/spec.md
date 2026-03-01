# Specification: Initialize project structure and core Action system

## Overview
This track focuses on establishing the base architecture and core data structures for `ebitenpad`, following the design specifications laid out in the `README.md`.

## User Stories
- As a developer, I want a clear directory structure so I know where to add new features.
- As a developer, I want a unified `Action` system so I can abstract away physical input devices.

## Requirements

### 1. Directory Structure
Create the following directory structure:
- `input/`: Core input handling and device abstraction.
- `virtual/`: Virtual stick and button implementations.
- `examples/`: Example usage of the library.

### 2. Core Types (`input/input.go`)
- **Action:** A custom integer type (`type Action int`) to represent logical game actions.
- **ActionState:** A struct to represent the state of an action:
  ```go
  type ActionState struct {
      pressed      bool
      justPressed  bool
      justReleased bool
      x            float64
      y            float64
      strength     float64
  }
  ```

### 3. Input Manager (`input/input.go`)
- **Input Struct:** Manage the mapping of actions to states.
- **Constructor:** `NewInput() *Input`.
- **API Methods:**
  - `Update()`: Frame update logic.
  - `Pressed(Action) bool`
  - `JustPressed(Action) bool`
  - `JustReleased(Action) bool`

## Acceptance Criteria
- [ ] Directory structure is created according to the design.
- [ ] `Action` and `ActionState` types are defined with the specified fields.
- [ ] `NewInput` returns a valid `Input` instance.
- [ ] Unit tests verify that `ActionState` correctly tracks pressed/justPressed/justReleased states (initially mocked or manual).
