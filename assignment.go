package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)
var x=0
type Block struct {
	to         string
	from	   string
	amount     int
	hash         string
	previousHash string
	timestamp    time.Time
	nounce    int
	
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	
}


func (b Block) calculateHash() string {
	blockData := b.previousHash + b.to +b.from+string(b.amount)+ b.timestamp.String() 
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}





func CreateBlockchain() Blockchain {
	genesisBlock := Block{
			hash:      "0",
			timestamp: time.Now(),
			nounce: 0,
			previousHash: "0",
	}
	genesisBlock.hash=genesisBlock.calculateHash()
	return Blockchain{
			genesisBlock,
			[]Block{genesisBlock},
			
	}
}










func (b *Blockchain) addBlock(from, to string, amount int,nounce int) {
	
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
			from:         from,
			to:         to,
			amount:         amount,
			nounce:         nounce,
			previousHash: lastBlock.calculateHash(),
			timestamp:    time.Now(),
	}
	
	b.chain = append(b.chain, newBlock)
	x=x+1
	b.chain[x].hash=b.chain[x].calculateHash()
}
















func (b Blockchain) isValid() bool {
	for i := 1; i <= x; i++ {
		
			previousBlock := b.chain[i-1]
			currentBlock := b.chain[i]
			if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
					return false
			}
	}
	return true
}


func (b Blockchain) print(){
for i := 0; i <= x; i++ {
	fmt.Println("Block creation time is: ",b.chain[i].timestamp)
    fmt.Println("Sender is: ",b.chain[i].from)
	fmt.Println("Reciever is: ",b.chain[i].to)
	fmt.Println("Amount is: ",b.chain[i].amount)
	fmt.Println("Nouce is: ",b.chain[i].nounce)
	fmt.Println("Hash is: ",b.chain[i].hash)
	fmt.Println("Previous block hash is: ",b.chain[i].previousHash)
	fmt.Println()
  }

}

func (b Blockchain) modify(blocknum int,from, to string, amount int,nounce int){
    b.chain[blocknum].from=from	
	b.chain[blocknum].to=to
	b.chain[blocknum].amount=amount
	b.chain[blocknum].nounce=amount
	b.chain[blocknum].hash=b.chain[blocknum].calculateHash()
	}



func main() {
	blockchain := CreateBlockchain()
	blockchain.addBlock("Alice", "Bob", 5,32)
	blockchain.addBlock("John", "Bob", 2,32)
	blockchain.addBlock("man", "dob", 4,55)	
	fmt.Println(blockchain.chain[2].hash)
//blockchain.modify(2,"jawad","babds",23,23)
//fmt.Println(blockchain.chain[2].hash)
//fmt.Println(blockchain.chain[2].data)
	//fmt.Println(blockchain.chain[1].hash)
	//fmt.Println(blockchain.chain[2].hash)
	blockchain.print()
	fmt.Println(blockchain.isValid())
	fmt.Println("after modification")
	blockchain.modify(2,"jawad","babds",23,23)
	fmt.Println(blockchain.isValid())
	//fmt.Println(blockchain.chain[3].hash)
	//blockchain.chain[3].nounce=100
	//blockchain.chain[3].hash=blockchain.chain[3].calculateHash()
	//fmt.Println(blockchain.chain[3].hash)
	//fmt.Println(blockchain.isValid())
	//fmt.Println(blockchain.chain[2])
	//fmt.Println(blockchain.chain[0].hash)
	//fmt.Println(blockchain.chain[1].previousHash)
	//fmt.Println(blockchain.chain[1].hash)
	//fmt.Println(blockchain.chain[2].previousHash)
	//fmt.Println(blockchain.chain[2].hash)
	//fmt.Println(blockchain.chain[3].previousHash)
	
	//for i := 0; i < x; i++ {
	//	fmt.Println(blockchain.chain[i])
	//  }

	//fmt.Println(blockchain.isValid())

	//fmt.Println("hello")

}
