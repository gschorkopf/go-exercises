package main

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go compute(i, c)
	}
	for i := 0; i < 10; i++ {
		println(i, <-c)
	}
}

func compute(id int, c chan int) {
	c <- id
}
