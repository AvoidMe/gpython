This is an experimental implementation of python 3.10 interpreter, written in golang.

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
    * int *(right now backend is just int64)
    * float
    * function
    * bool
    * NoneType

* 17/165 opcodes

# Roadmap

* Right now I use my python-written pseudo-compiller to run things, I want to implement *.pyc parser and run python files from it. In the future I will implement full-fledged python parser, but for now *.pyc will work.
* CPython hash can be negative numbers, but I used uint64, move it to regular int64.
* PyInt implementation right now is based on int64, which is incompattible with python, add a infinity number library and use it as a PyInt backend.
* Clean-up non-exception related TODO's.
* Add exceptions and cleanup exception-related TODO's

\* This is not a complete roadmap, only a bunch of goals on a week or two to give readers a better understanding about this project.