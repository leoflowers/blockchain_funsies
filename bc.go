package main

import (
	"fmt"
	"time"
	"errors"
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


func new_blockchain() *Blockchain {
	new_bc := new(Blockchain)
	new_bc.transactions = make([]Transaction, 0)
	new_bc.chain = make([]Block, 0)

	new_bc.chain = append(new_bc.chain, new_bc.new_block(1, "100"))	// genesis block
	return new_bc
}


func (bc Blockchain) new_block(pf int, prev_hash string) Block {
	block := Block{
		index: len(bc.chain) + 1,
		timestamp: time.Now(),
		transactions: bc.transactions,
		proof: pf,
		previous_hash: prev_hash,
	}

	bc.transactions = nil

	return block
}


func (bc Blockchain) new_transaction(amt int, s string, r string) int {
	bc.transactions = append(bc.transactions, Transaction{amt, s, r})
	return len(bc.chain) + 1
}


