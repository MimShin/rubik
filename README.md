# rubik
A simple go program to solve the Rubik's cube (finds the optimum solution)
Currently it only works for solutions up to 8 moves (8 moves can take up to 35 mins)

# usage
The following command scrambles the cube with n random moves and then tries to solve it
```
$ go run cmd/rubik_main.go <n>
```
eg. 
```
$ go run cmd/rubik_main.go 8
```
