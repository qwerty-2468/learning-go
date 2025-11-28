# Sorting data

Write a function that receives a slice as its argument and returns a slice that contains the original values in ascending order (if the slice contains strings, sort from 'A' - 'Z').
The sorting should satisfy the following requirements:
- The function should work with slices of type *`int64 and float64 }`* but is not required to work with other types
- You must not use type checking and the function must not accept a slice with type interface{} (hint: read the name of the parent directory :D)

An example for the usage:
```python
ints := {12, -3, 41, 27}
floats := {3.14, -1.2,2.781}
sortSlice(ints) // returns {-3, 12, 27, 41}
sortSlice(floats) // returns {-1.2, 2.781, 3.14}
```


Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.