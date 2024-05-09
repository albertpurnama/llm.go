package llmgo

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	stringToBeEncoded := "hello world, this is albert"

	// convert string to utf8
	utf8EncodedBytes := []byte(stringToBeEncoded)

	maxPair := maxBytePairOccurence(utf8EncodedBytes)
	fmt.Println(maxBytePairOccurence(utf8EncodedBytes))

	// convert bytes back to utf-8 string
	utf8DecodedString := string(maxPair)
	fmt.Println("Decoded UTF-8 string from byte pair:", utf8DecodedString)

}
