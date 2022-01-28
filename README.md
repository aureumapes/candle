# Candle

**Candle** is an interpreted scripting language designed to work mostly with math and numbers. *It is not feature, nor Turing complete and is still in alpha version. You should not use it in production!*



## Instructions

Candle has only a few instructions, which you can use in your program:

- `goto <LINE_NUMBER>` Jumps to a line and performs the action in that line

- `set <VARIABLE_NAME> to <VALUE>` Creates a new variable if a name and sets a default number to it

- `add <VALUE> <VARIABLE_NAME>` Adds a number to a variable

- `sub <VALUE> <VARIABLE_NAME>` Subtracts a number from a variable

- `mult <VALUE> <VARIABLE_NAME>` Multiplies a variable by a number

- `div <VALUE> <VARIABLE_NAME>` Divides a variable by a number

- `sqrt <VARIABLE_NAME> <VALUE/VARIABLE>` Gets the square root of a variable and sets the variable to the result

- `input <VARIABLE_NAME>` Sets the variable to the input from **stdin**

- `say <VARIABLE>` Prints the value of a variable to the screen

- `if <VARIABLE_NAME> !=/==/>=/<= <NUMBER> <COMMAND>` Executes the command if the condition is matched



## Examples
### 1 to 10 Counter
```
set a to 0
add 1 a
say a
if a != 11 goto 2
```

## Contributors

- [AureumApes](https://github.com/AureumApes)

- [Tilo Behnke](https://github.com/TUBS1001)

- [SomehowPatrick](https://github.com/SomehowPatrick)
