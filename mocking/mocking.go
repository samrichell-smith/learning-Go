package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout)
}


const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func Countdown(out io.Writer) {

	for i := countdownStart; i>0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(out, finalWord)
}