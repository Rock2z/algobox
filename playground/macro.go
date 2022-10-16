package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	go func() {
		select {
		case <-time.After(time.Second * 10):
			os.Exit(0)
		}
	}()
	low()
	//F2Chan := make(chan bool)
	//f2Macro(F2Chan)
}

func f2Macro(ch chan bool) {
	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		fmt.Println("1 is pressed, e = ", e)
		for true {
			fmt.Println("auto press 2")
			err := robotgo.KeyPress("2")
			if err != nil {
				return
			}
			robotgo.MilliSleep(500)
		}
	})
	s := hook.Start()
	<-hook.Process(s)
}

func getKeepPressFn(flag bool) func(e hook.Event) {
	return func(e hook.Event) {
	}
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}

func event() {
	ok := hook.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	keve := hook.AddEvent("k")
	if keve {
		fmt.Println("you press... ", "k")
	}

	mleft := hook.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}
