package main

import (
	"fmt"
	"math/rand"
	"time"
)

var try int
var testtally int = 1

func main() {
	origin := rand.NewSource(time.Now().UnixNano())
	random := rand.New(origin)
	confidential := random.Intn(10)

	fmt.Println("Its Time to Judge the Number!")

	for {
		fmt.Println("Make Your Judgment")
		fmt.Scan(&try)
		if testtally > 5 {
			fmt.Println("Tough Luck! Have Another Try")
			break
		} else {
			if try > confidential {
				fmt.Println("No! The Magic Number is Smaller")
			} else if try < confidential {
				fmt.Println("No! The Magic Number is Bigger")
			} else {
				fmt.Println("BullsEye")
				break
			}

		}

		testtally++
	}
}
