# Get to work problem
Author: Frank Yung
Go version: 1.20

### Some notes and explanations
- In `main.go`, function `readFileAsString` is for reading content from the `input.txt` file
- "calculator" package contains most of the logic 
- A custom sort is created for helping to calculate the minimum cars needed for each town
- The most straightforward way is implemented to try to solve the problem
- Due to time constraints(90mins), some improvement cannot be done and some unit tests are missing (coverage: 91.1% of statements in calculator package)

### Run the project
```go
go run main.go
```

### Unit test
``` sh
make unit-test
```
Unit test report (coverage.html) will be generated 

### Input parameters
Please provide input in input.txt and place it in the project root folder

