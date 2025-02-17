# Integer Overflow in Go Summation Function

This repository demonstrates a potential integer overflow bug in a simple Go function that calculates the sum of elements in an integer slice.  The function `myFunc` iterates through the slice and adds each element to a running sum. If the sum exceeds the maximum value of an int, an integer overflow will occur, leading to incorrect results.

The `bug.go` file contains the buggy code, while `bugSolution.go` provides a solution using `int64` to prevent overflow.  This is a common issue in many programming languages when dealing with large sums or multiplications.  The solution illustrates the importance of selecting appropriate data types to avoid such errors.

## Solution

The solution uses `int64`, which has a significantly larger range than `int`, to prevent overflow.  Consider using even larger integer types (such as `math/big.Int`) for exceptionally large sums.