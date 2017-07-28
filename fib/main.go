package main

func fibonacci(n int, c chan int) {
	x := 1
	y := 1

	for i := 0; i <= n; i++ {
		c <- x

		// This works
		x, y = y, x+y

		// Why doesn't this work?
		// x = y
		// y = x+y

	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		println(i)
	}
}
