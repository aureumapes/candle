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

- `ask <VARIABLE_NAME>` Asks the console to insert a number and sets the variable to the result

- `say <VARIABLE> <message>` Asks the user for a number with a message and sets the variable to the result

- `if <VARIABLE_NAME> !=/==/>=/<=/>/< <NUMBER> <COMMAND>` Executes the command if the condition is matched

- `exit` Exits the program

- Comments are only allowed in a line, you can use comments in new lines like this
  
  ```
  # this is a comment
  # this is allowed
  // and this is also allowed
  say x
  # this is another comment
  say x # this is invalid
  ```

## Examples

You can find examples in the [examples](examples/)  folder in the repository.

## Contributors

- [AureumApes](https://github.com/AureumApes)

- [Tilo Behnke](https://github.com/TUBS1001)

- [SomehowPatrick](https://github.com/SomehowPatrick)
