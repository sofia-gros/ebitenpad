# Product Guidelines: ebitenpad

## Core Philosophy
The library should feel like a natural extension of Ebitengine. It should prioritize simplicity and performance while providing enough abstraction to handle complex input scenarios.

## API Design Guidelines
- **Consistency:** Use clear, consistent naming (e.g., `Pressed`, `JustPressed`, `JustReleased`).
- **Simplicity:** Keep the public API surface area as small as possible.
- **Polling-Based:** Follow Ebitengine's architecture by using a polling model (`Update()`, `Draw()`, and state queries).
- **Fluency:** Use method chaining for configuration (e.g., `SetPosition(x, y).SetRadius(r)`).

## Visual & Interaction Guidelines (Virtual UI)
- **Visual Feedback:** All virtual controls should provide clear visual feedback when touched or active.
- **Interactive Feedback:** Virtual sticks should respond immediately to touch without perceived lag.
- **Modern Aesthetics:** Standard virtual sticks and buttons should have a clean, minimalist design that fits a wide range of game styles.

## Documentation Guidelines
- **Clarity:** Use concise and technically accurate language.
- **Code Examples:** Every major feature should include a brief, runnable code snippet.
- **Architecture Explanations:** Provide diagrams or clear descriptions of the internal normalization model when necessary.

## Development Style
- **Performance:** Avoid allocations in the update loop.
- **Safety:** Ensure multi-touch logic is robust and doesn't lead to "stuck" inputs.
- **Compatibility:** Test against various Ebitengine versions and platforms (Web, Mobile, Desktop).
