package main

import "fmt"

func main () {
	blockchain := new_blockchain()

	blockchain.new_transaction(100, "leo", "aleena")
	blockchain.new_transaction(50, "aleena", "leo")
	blockchain.print_chain()

	fmt.Printf("done\n")
}
