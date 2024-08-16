This is my first blog post using a Gleam-based blogging system. I'm excited to share my thoughts on web development and programming in general.

## Why Gleam?

Gleam is a fantastic language that combines the best of functional programming with strong typing. Here are a few reasons why I love Gleam:

1. **Type Safety**: Gleam's type system helps catch errors at compile-time.
2. **Erlang Compatibility**: We can leverage the power of the BEAM ecosystem.
3. **Friendly Syntax**: Gleam's syntax is clean and easy to read.

## Code Example

Here's a simple Gleam function:

```gleam
pub fn greet(name: String) -> String {
  "Hello, " <> name <> "!"
}
