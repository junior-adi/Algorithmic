
/*
Author: Junior ADI
Description: Brief description of the code file
Date: April 8th 2024, 12:17 AM GMT
Location: Abidjan, Cote d'Ivoire.

Original repository: https://github.com/zhenrong-wang/filter-uniq-ints.git

This code is licensed under the MIT License.

Copyright (c) 2024, Junior ADI
*/

package main

import (
    "fmt"
    "time"
)

func removeDuplicates(input []int) []int {
    var output []int
    for _, elem := range input {
        found := false
        for _, val := range output {
            if val == elem {
                found = true
                break
            }
        }
        if !found {
            output = append(output, elem)
        }
    }
    return output
}

func removeDuplicatesImproved(input []int) []int {
    var output []int
    var max, min int

    for _, elem := range input {
        // Update max and min if necessary
        if elem > max {
            max = elem
        }
        if elem < min || len(output) == 0 {
            min = elem
        }

        found := false
        for _, val := range output {
            if val == elem {
                found = true
                break
            }
        }
        if !found {
            output = append(output, elem)
        }
    }
    return output
}

func removeDuplicatesHashTable(input []int) []int {
    seen := make(map[int]bool)
    var output []int

    for _, elem := range input {
        if !seen[elem] {
            seen[elem] = true
            output = append(output, elem)
        }
    }
    return output
}

type hashTableBaseNode struct {
	branchSizeP uint32
	branchSizeN uint32
	ptrBranchP  []int
	ptrBranchN  []int
}

func newHashTableBaseNode(branchSizeP, branchSizeN uint32) *hashTableBaseNode {
	return &hashTableBaseNode{
		branchSizeP: branchSizeP,
		branchSizeN: branchSizeN,
		ptrBranchP:  make([]int, branchSizeP),
		ptrBranchN:  make([]int, branchSizeN),
	}
}

func removeDuplicatesDynamicHashTable(input []int) []int {
	const (
		initialSize = 32
		modValue    = 65536
	)

	hashTable := make([]*hashTableBaseNode, initialSize)

	var output []int

	for _, elem := range input {
		hashIndex := elem % initialSize
		if hashIndex < 0 {
			hashIndex = -hashIndex
		}

		if hashTable[hashIndex] == nil {
			hashTable[hashIndex] = newHashTableBaseNode(modValue, modValue)
		}

		var ptrBranch []int
		if elem >= 0 {
			ptrBranch = hashTable[hashIndex].ptrBranchP
		} else {
			ptrBranch = hashTable[hashIndex].ptrBranchN
		}

		modIndex := elem % modValue
		if modIndex < 0 {
			modIndex = -modIndex
		}

		if ptrBranch[modIndex] == 0 {
			output = append(output, elem)
			ptrBranch[modIndex] = 1
		}
	}

	return output
}

type bitHashTableNode struct {
	branchSizeP uint32
	branchSizeN uint32
	ptrBranchP  []uint8
	ptrBranchN  []uint8
}

func newBitHashTableNode(branchSizeP, branchSizeN uint32) *bitHashTableNode {
	return &bitHashTableNode{
		branchSizeP: branchSizeP,
		branchSizeN: branchSizeN,
		ptrBranchP:  make([]uint8, (branchSizeP+7)/8),
		ptrBranchN:  make([]uint8, (branchSizeN+7)/8),
	}
}

func flipBit(byteArray []uint8, bitPosition int) {
	byteIndex := bitPosition / 8
	bitOffset := uint8(1 << uint(bitPosition%8))
	byteArray[byteIndex] |= bitOffset
}

func checkBit(byteArray []uint8, bitPosition int) bool {
	byteIndex := bitPosition / 8
	bitOffset := uint8(1 << uint(bitPosition%8))
	return (byteArray[byteIndex] & bitOffset) != 0
}

func removeDuplicatesBitHashTable(input []int) []int {
	const (
		modValue = 65536
	)

	hashTable := make([]*bitHashTableNode, modValue)

	var output []int

	for _, elem := range input {
		hashIndex := elem % modValue
		if hashIndex < 0 {
			hashIndex = -hashIndex
		}

		if hashTable[hashIndex] == nil {
			hashTable[hashIndex] = newBitHashTableNode(modValue, modValue)
		}

		var ptrBranch []uint8
		if elem >= 0 {
			ptrBranch = hashTable[hashIndex].ptrBranchP
		} else {
			ptrBranch = hashTable[hashIndex].ptrBranchN
		}

		modIndex := elem % modValue
		if modIndex < 0 {
			modIndex = -modIndex
		}

		if !checkBit(ptrBranch, modIndex) {
			output = append(output, elem)
			flipBit(ptrBranch, modIndex)
		}
	}

	return output
}

func main() {
	input := []int{16, 17, 2, 17, 4, 2, 97, 4, 17}

	/* NAIVE ALGORITHM */

	fmt.Printf("NAIVE ALGORITHM START\n")
	// Capture the time before executing removeDuplicates
	startTime := time.Now()

	// Execute removeDuplicates
	output := removeDuplicates(input)

	// Capture the time after executing removeDuplicates
	endTime := time.Now()

	// Calculate the execution duration
	duration := endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", input)
	fmt.Printf("Filtered array: %v\n", output)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("NAIVE ALGORITHM END\n\n")

	/* IMPROVED ALGORITHM */

	fmt.Printf("IMPROVED ALGORITHM START\n")
	// Capture the time before executing removeDuplicatesImproved
	startTime = time.Now()

	// Execute removeDuplicatesImproved
	output = removeDuplicatesImproved(input)

	// Capture the time after executing removeDuplicatesImproved
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", input)
	fmt.Printf("Filtered array: %v\n", output)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("IMPROVED ALGORITHM END\n\n")


	/* HASH TABLE ALGORITHM */
	fmt.Printf("HASH TABLE ALGORITHM START\n")
	// Capture the time before executing removeDuplicatesHashTable
	startTime = time.Now()

	// Execute removeDuplicatesHashTable
	output = removeDuplicatesHashTable(input)

	// Capture the time after executing removeDuplicatesHashTable
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", input)
	fmt.Printf("Filtered array: %v\n", output)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HASH TABLE ALGORITHM END\n\n")


	/* HASH TABLE DYNAMIC ALGORITHM */
	fmt.Printf("HASH TABLE DYNAMIC ALGORITHM START\n")
	// Capture the time before executing removeDuplicatesDynamicHashTable
	startTime = time.Now()

	// Execute removeDuplicatesDynamicHashTable
	output = removeDuplicatesDynamicHashTable(input)

	// Capture the time after executing removeDuplicatesDynamicHashTable
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", input)
	fmt.Printf("Filtered array: %v\n", output)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HASH TABLE DYNAMIC ALGORITHM END\n\n")


	/* BIT MAP ALGORITHM */
	fmt.Printf("BIT MAP ALGORITHM START\n")
	// Capture the time before executing removeDuplicatesBitMap
	startTime = time.Now()

	// Execute removeDuplicatesBitMap
	output = removeDuplicatesBitHashTable(input)

	// Capture the time after executing removeDuplicatesBitMap
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", input)
	fmt.Printf("Filtered array: %v\n", output)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("BIT MAP ALGORITHM END\n\n")
	
}
