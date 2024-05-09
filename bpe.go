package llmgo

func maxBytePairOccurence(input []byte) []byte {
	bytePairCount := make(map[uint16]int)

	bytePairKey := func(i, j uint8) uint16 {
		return uint16(i)<<8 | uint16(j)
	}

	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			// we can directly convert this to uint8 because
			// uint8 ranges from 0 - 255
			key := bytePairKey(uint8(i), uint8(j))
			bytePairCount[key] = 0
		}
	}

	inputLength := len(input)
	for i, b := range input {
		if i+1 == inputLength {
			// in this case we're at last byte, we don't need to add this to
			// bpe map
			continue
		}

		// we want to track the byte pair count here
		currentByte := b
		nextByte := input[i+1] // we know we're not in the last byte.

		// increment counter
		currentPairKey := bytePairKey(uint8(currentByte), uint8(nextByte))
		bytePairCount[currentPairKey] += 1
	}

	// find the highest byte pair count
	var maxCount int
	var maxBytePairs []byte

	for pair, count := range bytePairCount {
		if count > maxCount {
			maxCount = count
			maxBytePairs = []byte{byte(pair >> 8), byte(pair & 0xff)}
		}
	}

	return maxBytePairs
}
