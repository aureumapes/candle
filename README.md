Candle
======

Candle is a simple esoteric research programing language. Built using [Go](https://go.dev) and featuring a well-crafted instruction set, it provides fast execution time and produces easy-to-read code.

> ⚠️ Please keep in mind that the project is in its alpha stage. Bugs, crashes, and missing features are expected. The instruction set is also incomplete and features undefined behavior.

✏ Language features
-------------------

*   line-by-line execution and interpretation, where instruction has its line
*   strings and 64-bit float as datatypes
*   Comments, introduced with `//`, only working in a line
*   [Functions](#functions)

📩 Installation
------------
### Arch based Distos
Use a aur-helper of your choice, I use yay:
`yay -S candle`<br>
OR of course build it manually:
`git clone https://aur.archlinux.org/candle.git && cd candle && makepkg -si`

### Other Distors
1. Download the repo: `git clone https://github.com/aureumapes/candle.git`
2. Build the source: `cd candle && go build`
3. Copy the candle-binary where you can execute it


📎 Instructions
---------------

Candle features a variety of instructions for you to use in your programs

### Set

The set instruction allows you to set a value of a variable. If the variable name wasn't previously defined a new variable will be created, otherwise, the value of the variable will be overridden by the new one. The set instruction allows for inserting a variable as a value. If you want to set an string as the value, add the prefix `+` to the value. Variable names must follow the RegEx `[a-z][a-z|0-9]*`  
Example: `set pi to 3.14`, `set x to y`, `set hi to +Hello World`

### Say

With the say instruction, you can print a variable or string to the console. Prefix strings as input with a `+` like `say +Hello World`, otherwise the interpreter will interpret `Hello World` as a variable name.  
Example: `say x`, `say +Hello World`

### Sayln

Sayln is the same as the say instruction but will append a new line character at the end of the message.  
Example: `sayln x`, `sayln +Hello World`

### Ask

Wanted to ask for input. The `ask` instruction is your go-to option. It follows the syntax `ask <variable> <message>` where the variable is the variable name to store the input in (if not available a new one will be created) and the message is the message asked in the console.  
Example: `ask pi What is pi equal to?`

### Read
Read saves the content of a file into a variable<br/>
Example: `read main.cndl cat`

### Add

The add instruction is used to clalculate the sum of a Variable and anouther Variable or Number. The result will override the value of the first variable.  
Example: `add a -5`, `add a b`

### Sub, Mul, Div

These instructions do the same as add but instead calculate the difference, product, or quotient of the variable and the other variable or number.  
Example: `sub a 8`, `mul a 5`, `div a 2`

### If

With if you can execute a command if a variable is ==, !=, <=, >=, < or > to another variable or other value  
Example: `if a == 9 sayln +a is equal to 9`

### Functions

The `function` keyword initiates functions. You end the function body with the `end` keyword.  
The body will be executed, when the function is created and when the fuction name is called in a line, as a instruction.  
Example:  
`function main`<br />
`sayln +Hello World`<br />
`end` <br />
`main`

### Import
`import` reads another candle file and writes it into a function in your current program<br/>
Make sure to give the File name without the extension(.cndl)<br/>
Examples:
`import hello-world`

⌛ Examples
----------

You can find full example files in the [examples](https://github.com/Amiraxoba2/candle/tree/master/examples) folder.

📝 Implementations In different Languages
-----------------------------------------

*   [Go](https://github.com/Amiraxoba/Candle/) (The original Implementation)

🚢 Connect with us
------------------

We coordinate the development through our [Discord](https://discord.gg/sFaZtaSX9j).
