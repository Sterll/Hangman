package main

import (
	"container/list"
	"fmt"
	"os"
)

type STATE int64

const (
	WAITING STATE = iota
	PLAYING
	END
)

var GAME_STATE STATE

func main() {
	fmt.Print("f")
	list.New()
	os.Exit(0)

	for {
		if GAME_STATE == WAITING || GAME_STATE == END {
			break
		}
		
	}
}
