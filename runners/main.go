package main

type runner interface {
	name() string
	run(distance int) (seconds int)
}

type geoff struct{}
type blithe struct{}
type thom struct{}

func (geoff) name() string  { return "Geoff" }
func (blithe) name() string { return "Blithe" }
func (thom) name() string   { return "Thom" }

func (geoff) run(d int) int  { return d * 2 }
func (blithe) run(d int) int { return d }
func (thom) run(d int) int   { return d / 2 }

func race(distance int, runners ...runner) (winner string, fastest int) {
	for i, runner := range runners {
		seconds := runner.run(distance)
		if i == 0 || seconds < fastest {
			fastest, winner = seconds, runner.name()
		}
	}

	return winner, fastest
}

func main() {
	winner, fastest := race(100, geoff{}, blithe{}, thom{})
	println("The winner is", winner, "with a time of", fastest)

}
