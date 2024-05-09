package main

import (
	"fmt"

	"github.com/joshcarp/llmgo"
)

func main() {
	fmt.Println(byte('h'))
	fmt.Println(string([]byte{104}))
	llmgo.NewTokenizer("./gpt2_tokenizer.bin")
}
