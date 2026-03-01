# Initial Concept
ebitenpad is an input integration library for Ebitengine (Go), designed to simplify and unify input management across various devices (keyboard, gamepad, touchscreen).

# Product Guide: ebitenpad

## Vision
To provide Ebitengine developers with a robust, abstraction-based input system that allows them to focus on game logic rather than device-specific polling and mapping.

## Target Users
- Ebitengine (Go) game developers.
- Developers targeting multiple platforms (Mobile, PC, Web) who need consistent input behavior.

## Core Goals
- **Device Abstraction:** Move from "Is the Space key pressed?" to "Is the 'Jump' action triggered?".
- **Unified API:** Provide a consistent interface for keyboard, gamepad, and virtual touch controls.
- **Seamless Multi-platform Support:** Built-in virtual pads for mobile devices that integrate directly into the action system.
- **Developer Experience:** Lightweight, easy to integrate, and highly extensible.

## Key Features
- **Action-Based Input:** Map multiple physical inputs to a single logical action.
- **Normalization Model:** Consistent state reporting (`pressed`, `justPressed`, `justReleased`, `x/y` vectors, `strength`).
- **Built-in Virtual UI:** Ready-to-use virtual sticks (Fixed, Free) and buttons with multi-touch support.
- **Sophisticated Direction Mapping:** Built-in 4-way and 8-way direction normalization.
- **ID-Based Touch Locking:** Multi-touch management to prevent control interference.

## Future Potential
- Input deadzone configuration.
- Replay system and network synchronization.
- JSON-based mapping configurations.
