This is an experimental implementation of python interpreter in golang.

How to use:
```
echo 'print("Hello", "world")' > hello_world.py
python python_to_instructions.py hello_world.py output.json && go run main.go
```