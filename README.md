[![builds.sr.ht status](https://builds.sr.ht/~tristan957/monkey.svg)](https://builds.sr.ht/~tristan957/monkey?)

# Monkey

An interpreter for the Monkey programming language.

## Customizations

- The lexer operates on buffered I/O instead of reading the entire input as a
  string. This saves on memory usage.
- Tokens have spans associated with them, so in the case of unknown input, the
  interpreter can report line and column numbers. Spans mark the beginning and
  end of a token.
- Added binary, octal, and hexadecimal integer literal support.
  - `0b0101`
  - `0o1234`
  - `0xffa4`
