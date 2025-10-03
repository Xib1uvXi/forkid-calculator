package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/forkid"
)

var (
	genesisPath = flag.String("genesis", "genesis.json", "path to genesis.json")
)

func main() {
	flag.Parse()

	genesisFile, err := os.ReadFile(*genesisPath)
	if err != nil {
		fmt.Printf("Failed to read genesis file: %v\n", err)
		os.Exit(1)
	}

	var genesis core.Genesis
	if err := json.Unmarshal(genesisFile, &genesis); err != nil {
		fmt.Printf("Failed to unmarshal genesis file: %v\n", err)
		os.Exit(1)
	}

	genesisBlock := genesis.ToBlock()

	id := forkid.NewID(genesis.Config, genesisBlock, 0, 0)

	fmt.Printf("ForkID Hash: %#x\n", id.Hash)
	fmt.Printf("ForkID Next: %d\n", id.Next)
}
