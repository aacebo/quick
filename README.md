# Quick

a compromise between the `Rust` and `Go` programming languages. It is a compiled and interpreted, typed language that balances ease of use and performance.

## Why?

I started this project because as I used newer languages like `Rust`/`Go` I noticed a trend of awesome new features. `Rust` has a very powerful type system while `Go` keeps things simple by, for example, only have one type of loop statement (`for`). Both excel in very different ways but both have their caviats.

In `Rust` you have to fight with the borrow checker, for low level performance intensive tasks this is the right tool for the job in many cases, but not necessarily for gerneral purpose API development.

In `Go` you have a substantially weaker type system and less features than `Rust`, but in contrast it has a simpler pattern for concurrency and API development.

My ideal outcome for this language is to offer the same simplicity/ease-of-use as `Go` while adding some of the useful features that `Rust` users enjoy.

## Roadmap

| Title                  | Keywords                 | Status    | Example                               |
|------------------------|--------------------------|-----------|---------------------------------------|
| Variables              | `let` `const`            | [ ]       | [example](./examples/variables.gpp)   |
| Primitives             | `string` `number` `bool` | [ ]       | [example](./examples/primitives.gpp)  |
| Nilable                | `?`                      | [ ]       | [example](./examples/nilable.gpp)     |
| Strings                |                          | [ ]       | [example](./examples/strings.gpp)     |
| If                     | `if` `else if` `else`    | [ ]       | [example](./examples/if.gpp)          |
| Match                  | `match`                  | [ ]       | [example](./examples/match.gpp)       |
| Loops                  | `for`                    | [ ]       | [example](./examples/for.gpp)         |
| Functions              | `fn`                     | [ ]       | [example](./examples/fn.gpp)          |
| Structs                | `struct`                 | [ ]       | [example](./examples/struct.gpp)      |
| Garbage Collection     |                          | [ ]       |                                       |
| Threads                | `go`                     | [ ]       | [example](./examples/go.gpp)          |
| Async Functions        | `async` `await`          | [ ]       | [example](./examples/async.gpp)       |
| Modules                | `mod`                    | [ ]       | [example](./examples/mod.gpp)         |
| Imports                | `use`                    | [ ]       | [example](./examples/use.gpp)         |
| Syntax Error Handling  |                          | [ ]       |                                       |
| Runtime Error Handling |                          | [ ]       |                                       |
| Visibility             | `pub`                    | [ ]       | [example](./examples/visibility.gpp)  |
| Inheritance            | `extends`                | [ ]       | [example](./examples/inheritance.gpp) |
| Generics               | `<T>`                    | [ ]       | [example](./examples/generics.gpp)    |
| Exceptions             | `throw` `try` `catch`    | [ ]       | [example](./examples/exceptions.gpp)  |
| Slices                 | `[]`                     | [ ]       | [example](./examples/slices.gpp)      |
| Maps                   | `map[K]V`                | [ ]       | [example](./examples/maps.gpp)        |

## Future (Advanced)

| Title                  | Keywords                 | Status    | Example                               |
|------------------------|--------------------------|-----------|---------------------------------------|
| Struct Tags            |                          | [ ]       | [example](./examples/tags.gpp)        |
| Reflection             |                          | [ ]       | [example](./examples/reflection.gpp)  |
| Testing                |                          | [ ]       | [example](./examples/testing.gpp)     |
| Decorators             | `@`                      | [ ]       | [example](./examples/decorators.gpp)  |

## Author

[Alexander Acebo](mailto:aacebowork@gmail.com)
