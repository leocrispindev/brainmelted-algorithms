package main

func main() {
	balancelimit := 1000

	value := 1001
	balance := 0

	newB := balance - value

	isless := newB < -balancelimit

	println(isless)
	println(newB)
	println(-balancelimit)

}
