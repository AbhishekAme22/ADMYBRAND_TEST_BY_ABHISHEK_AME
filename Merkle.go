package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read in the transactions
	scanner := bufio.NewScanner(file)
	var transactions []string
	for scanner.Scan() {
		transactions = append(transactions, scanner.Text())
	}

	// Compute the Merkle root
	root := computeMerkleRoot(transactions)

	// Output the result
	fmt.Printf("Merkle root: %s\n", root)
}

func computeMerkleRoot(transactions []string) string {
	// If there are no transactions, return an empty string
	if len(transactions) == 0 {
		return ""
	}

	// If there is only one transaction, return its hash
	if len(transactions) == 1 {
		return transactions[0]
	}

	// Compute the hashes of each transaction
	var hashes [][]byte
	for _, transaction := range transactions {
		hash, err := hex.DecodeString(transaction)
		if err != nil {
			panic(err)
		}
		hashes = append(hashes, hash)
	}

	// Compute the Merkle root
	for len(hashes) > 1 {
		if len(hashes)%2 == 1 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}
		var nextLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			combined := append(hashes[i], hashes[i+1]...)
			hash := sha256.Sum256(combined)
			nextLevel = append(nextLevel, hash[:])
		}
		hashes = nextLevel
	}

	// Encode the Merkle root as a hex string and return it
	return hex.EncodeToString(hashes[0])
}
