package main

import (
	"crypto/sha256"
	"bytes"
)


func validate_proof(last_proof []byte, current_proof []byte) bool {
	h := sha256.New()
	h.Write(current_proof)
	h.Write(last_proof)

	if h.String()[0:4] == "0000" {
		return true
	}

	return false
}

func proof_of_work(last_proof []byte) []byte {
	// proof of work algorithm:
	proof := 0

	for {
		h := sha256.New()
		h.Write([]byte(proof))
		h.Write(last_proof)

		if h.String()[0:4] == "0000" {
			return []byte(proof)
		}

		proof = proof + 1
	}
}
