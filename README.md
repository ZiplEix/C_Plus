# C+

## Introduction

C+ is a simple programming language that is designed to enhance the `c` syntaxe. It is a statically typed language that is transpiled to C. The language is designed to be more readible than C, while still being powerful enough to write real programs. and wildly inspired by the `Go` programming language.

C+ is a superset of C, meaning that all valid C programs are also valid C+ programs. This means that you can take any C program and compile it with the C+ compiler without any modifications.

It is important to note that C+ is not a new language, but rather a new syntaxe for C. This means that the C+ compiler transpiles C+ code to C code, which is then compiled with a C compiler. This means that C+ has the same performance as C, and can be used in the same contexts as C.

For the C compilation, The `C+` compiler will first try to use the `clang` compiler, and if it is not available, it will use the `gcc` compiler. If you want to use a specific compiler, you can use the `--compiler` flag to specify the compiler you want to use.

As the `C+` compiler use `C` compiler under the hood, you can use the `--flags` flag to pass flags to the `C` compiler. For example, you can use the `--flags="-lm -o2"` flag to link the math library and enable optimization level 2.

## Features

State | Feature | Desctiption |
| --- | --- | --- |
| ❌ | Type inference | C+ can infer the type of a variable from its initialization. This means that you don't have to write the type of a variable when you declare it. For example, you can write `x := 5;` instead of `int x = 5;`. |
| ❌ | Structs | C+ has a more powerful struct syntax than C. You can define a struct with the `struct` keyword, and access its fields with the `.` operator. You can then create struct related functions. You can then create struct related functions with this syntaxe: `int (maStruct *m) myMemeberFunction() { ... }`. |

## Parameters

The C+ compiler has the following parameters:

State | short | long | Description | Example |
| --- | --- | --- | --- | --- |
| ❌ | -h | --help | Display the help message. | `c+ -h` |
| ❌ | -v | --version | Display the version of the compiler. | `c+ -v` |
| ❌ | -c | --compiler | Specify the C compiler to use | `c+ -c gcc` |
| ❌ | -f | --flags | Pass flags to the C compiler. | `c+ -f "-lm -o2"` |
| ❌ | -o | --output | Specify the output directory. C+ will write the transpiled C code to this directory with the same file architecture. | `c+ -o output` |

## Command

The C+ compiler has the following command:

State | Command | Description | Example |
| --- | --- | --- | --- |
| ❌ | run | Compile and run the C+ program. | `c+ run program.cp` |
| ❌ |  | Compile the C+ programme in C and then built it | `c+ program.cp` |

## Example

Here is an example of a simple C+ program using the C+ syntaxe:

```c
#include <stdio.h>

typedef struct {
    int x;
    int y;
} point;

void (point *p) print(char *name) {
    printf("%s: (%d, %d)\n", name, p->x, p->y);
}

int main() {
    x := 5;
    y := 10;

    point p;

    p.x = x;
    p.y = y;

    p.print("The point");

    return 0;
}
```

Under ther hood, this program is transpiled to the following C code:

```c
#include <stdio.h>

typedef struct {
    int x;
    int y;
} Point;

void print_structPointMemberFunction(Point *p, char *name) {
    printf("%s: (%d, %d)\n", name, p->x, p->y);
}

int main() {
    int x = 5;
    int y = 10;

    Point p;

    p.x = x;
    p.y = y;

    print_structPointMemberFunction(&p, "The point");

    return 0;
}
```

## Installation

For the moment, the C+ compiler is not available. However, you can clone the repository and build the compiler yourself. To do this, you will need to have `GoLang` installed on your system. You can then run the following commands to build the compiler:

```bash
git clone
cd c+
go build src/main.go
```
