package main

import "fmt"

// Note: All function names same as given on gcr
// command : "go mod init Quiz_Blockchain" to initialize the mod module
// This is a block of blockchain
type Block struct {
	Data     string // data
	Previous *Block // pointer to previous block
}

// function to create newblok when data and pointer to previous block is given
func NewBlock(data string, previous *Block) *Block {
	return &Block{Data: data, Previous: previous}
}

// make changes in block data
func ModifyBlock(block *Block, newData string) {
	block.Data = newData
}

// displaying all blocks data
func DisplayAllBlocks(block *Block) {
	for block != nil { // will display until previous is null like when genisis block previous will come
		fmt.Println("Block Data:", block.Data)
		block = block.Previous
	}
}

func main() {
	//creating the first block that is genisis block  so  its previous pointer will be null
	genesis_block := NewBlock("Genesis_Block", nil)

	// making new blocks block 1 block 2
	block_1 := NewBlock("Block 1", genesis_block) // previous block  is genesis
	block_2 := NewBlock("Block 2", block_1)       // // previous block  iis block1

	//checking modifying function
	ModifyBlock(block_2, "Block 2 Data Changed")

	// Display all blocks
	fmt.Println("All Blocks:")
	DisplayAllBlocks(block_2)
}
