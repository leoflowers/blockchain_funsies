package main

import (
	"fmt"
	"log"
	"time"
	"crypto/sha256"
	"bytes"
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
	proof		[]byte
	previous_hash	[]byte
}


type Blockchain struct {
	transactions	[]Transaction
	chain		[]*Block
}


// creates new blockchain with a genesis block
func new_blockchain() *Blockchain {
	new_bc := new(Blockchain)
	new_bc.transactions = make([]Transaction, 0)
	new_bc.chain = make([]*Block, 0)

	new_bc.chain = append(new_bc.chain, new_bc.new_block([]byte(1), []byte(100)))	// genesis block
	return new_bc
}

// creates a single block and adds it to chain
func (bc Blockchain) new_block(pf []byte, prev_hash []byte) *Block {
	block := &Block{index: len(bc.chain) + 1,
			timestamp: time.Now(),
			transactions: bc.transactions,
			proof: pf,
			previous_hash: prev_hash,
		}

	bc.transactions = nil
	bc.chain = append(bc.chain, block)

	return block
}


// adds new transaction to list of currently floating transactions
func (bc Blockchain) new_transaction(amt int, s string, r string) int {
	bc.transactions = append(bc.transactions, Transaction{amt, s, r})
	return len(bc.chain) + 1
}

// hashes block and returns its hash
func hash_block(b Block) []byte {
	//buffer := new(bytes.Buffer)

	//err := binary.Write(buffer, binary.BigEndian, b)
	//if err != nil {
	//	log.Fatal("encode error: ", err)
	//}

	h := sha256.New()
	h.Write([]byte(b))
	return h.Sum(nil)
}

// prints current blockchain
func (bc Blockchain) print_chain() {
	for _, b := range bc.chain {
		fmt.Println("block index: ", b.index)
		fmt.Println("	timestamp: ", b.timestamp.Format(time.UnixDate))
		fmt.Println("	transactions:\n")

		for _, t := range b.transactions {
			fmt.Println("\t\tsender: ", t.sender)
			fmt.Println("\t\trecipient: ", t.recipient)
			fmt.Println("\t\tamount: ", t.amount)
		}
	}
}
