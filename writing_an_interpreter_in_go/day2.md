# Day2
## 1.1 Lexical Analysis
<img src="https://github.com/tsmrkk/go-training/blob/master/writing_an_interpreter_in_go/day2/1.png">

Lexer: Does lexical analysis. It is also called tokernizer or scanner

Lexical Analysis: The first transfomation, from source code to tokens. This process is also called "lexing".

Tokens: Small, easily categorized data structure that are then fed to parser

Parser: Does the second transfomation and turns the tokens in "Abstract Syntax Tree"

Example: Input into Lexer

```
"let x = 5 + 5"

```

Example: Output from Lexer

```
[
  LET,
  IDENTIFIER("x"),
  EQUAL_SIGN,
  INTEGER(5),
  PLUS_SIGN,
  INTEGER(5),
  SEMICOLON
]
```

What constitutes  a token varies between different lexer implementations. Some lexers, parse "5" to an integer in the parsing stage, while other lexers does not do that when constructing tokens

## 1.2 Defining Our Tokens
