package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}


const finalWord = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}


type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i>0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}