package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("🦆 RUBBER DUCKY RACE! 🦆")
	fmt.Println("Press Enter to start...")
	fmt.Scanln()

	ducky1Position := 0
	ducky2Position := 0
	finishLine := 20

	for ducky1Position < finishLine && ducky2Position < finishLine {
		// Random movement
		ducky1Position += rand.Intn(3) // moves 0-2 spaces
		ducky2Position += rand.Intn(3)

		// Draw the race
		fmt.Println("\n" + makeLane(ducky1Position, "🦆"))
		fmt.Println(makeLane(ducky2Position, "🦆"))
		fmt.Println("🏁" + strings.Repeat("═", finishLine-1) + "🏁")

		time.Sleep(500 * time.Millisecond)
	}

	if ducky1Position >= finishLine {
		fmt.Println("\n🎉 Ducky 1 wins!")
	} else {
		fmt.Println("\n🎉 Ducky 2 wins!")
	}
}

func makeLane(position int, duck string) string {
	lane := ""
	for i := 0; i < position; i++ {
		lane += "~"
	}
	lane += duck
	return lane
}
