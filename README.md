# Candle
Candle is a simple esoteric research programing language. Built using [Go](https://go.dev) and featuring a well-crafted instruction set, it provides fast execution time and produces easy-to-read code.

> ⚠️ Please keep in mind that the project is in its alpha stage. Bugs, crashes, and missing features are expected. The instruction set is also incomplete and features undefined behavior.

## ✏ Language features
- line-by-line execution and interpretation, where instruction has its line
- strings and 64-bit float as datatypes
- Comments, introduced with `//, only working in a line

## 📎 Instructions
Candle features a variety of instructions for you to use in your programs

### Set
The set instruction allows you to set a value of a variable. If the variable name wasn't previously defined a new variable will be created, otherwise, the value of the variable will be overridden by the new one. The set instruction allows for inserting a variable as a value.
If you want to set an string as the value, add the prefix `+` to the value. Variable names must follow the RegEx `[a-z]+[0-9]*`
<br />
Example: `set pi to 3.14`, `set x to y`, `set hi to +Hello World`

### Say
With the say instruction, you can print a variable or string to the console. Prefix strings as input with a `+` like `say +Hello World`, otherwise the interpreter will interpret `Hello World` as a variable name.
<br />
Example: `say x`, `say +Hello World`

### Sayln
Sayln is the same as the say instruction but will append a new line character at the end of the message.
<br />
Example: `sayln x`, `sayln +Hello World`

### Ask
Wanted to ask for input. The `ask` instruction is your go-to option. It follows the syntax `ask <variable> <message>` where the variable is the variable name to store the input in (if not available a new one will be created) and the message is the message asked in the console.
<br />
Example: `ask pi What is pi equal to?`

### Add
The add instruction is used to clalculate the sum of a Variable and anouther Variable or Number.
The result will override the value of the first variable.
<br />
Example: `add a -5`, `add a b`

### Sub, Mul, Div
These instructions do the same as add but instead calculate the difference, product, or quotient of the variable and the other variable or number.
<br />
Example: `sub a 8`, `mul a 5`, `div a 2`

### If
WIth if you can execute a command if a variable is ==, !=, <=, >=, < or > to another variable or other value
<br />
Example: `if a == 9 sayln +a is equal to 9`

## ⌛ Examples
You can find examples in the [examples](https://github.com/Amiraxoba/Candle/tree/v2/examples) folder.

## 📝 Implementations
* [Go](https://github.com/Amiraxoba/Candle/tree/go)(The original Implementation)

## 🚢 Connect with us
We coordinate the development through our [Discord](https://discord.gg/sFaZtaSX9j).