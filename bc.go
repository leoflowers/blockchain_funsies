package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	amount		int
	sender		string
	recipient	string
}


type Block struct {
	index		int
	timestamp	time.Time
	transactions	[]Transaction
	proof		int
	previous_hash	string
}


type Blockchain struct {
	transactions	[]Transaction
	chain		[]Block
}
