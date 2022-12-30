This is an experimental implementation of python 3.11 interpreter, written in golang.

How to use:
```
echo 'print("Hello, world")' > hello_world.py
python python_to_instructions.py hello_world.py output.json && go run main.go
```

# Motivation

1. I want to make smaller (in terms of LOC) and easily maintainable version of CPython. Golang gc, default collections, reflection and so on - should dramatically reduce python codebase.
1. I want to make python version without GIL, with golang concurency. Golang-related features (like goroutines and channels) will be added in additional module.
1. It feels good working on it.

# What's done already?

* A partial implementation of builtin types:

    * list
    * dict
    * tuple
    * str
    * int
    * float
    * function
    * bool
    * NoneType

* 17/165 opcodes

# Roadmap

* 
* Clean-up non-exception related TODO's.
* Add exceptions
* cleanup exception-related TODO's

\* This is not a complete roadmap, only a bunch of goals on a week or two to give readers a better understanding about this project.
