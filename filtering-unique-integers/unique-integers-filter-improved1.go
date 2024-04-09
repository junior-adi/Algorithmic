
/******************************************************************************

                            Author: Junior ADI
				Description: Brief description of the code file
				    Date: April 8th 2024, 12:17 AM GMT
						Location: Abidjan, Cote d'Ivoire.

    Original repository: https://github.com/zhenrong-wang/filter-uniq-ints.git

					This code is licensed under the MIT License.

                        Copyright (c) 2024, Junior ADI

*******************************************************************************/
package main

import (
    "fmt"
	"math/rand"
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


func generateRandomInputArr(arr []int, numElems, randMax int) error {
    if arr == nil {
        return fmt.Errorf("array is nil")
    }
    if numElems < 1 {
        return fmt.Errorf("number of elements is less than 1")
    }
    if randMax < 1 {
        return fmt.Errorf("randMax is less than 1")
    }

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < numElems; i++ {
        signFlag := rand.Intn(2)
        randNum := rand.Intn(randMax)
        if signFlag%2 == 0 {
            arr[i] = randNum
        } else {
            arr[i] = -randNum
        }
    }
    return nil
}

func generateGrowingArr(arr []int, numElems int) error {
    if arr == nil {
        return fmt.Errorf("array is nil")
    }
    if numElems < 1 {
        return fmt.Errorf("number of elements is less than 1")
    }

    for i := 0; i < numElems; i++ {
        arr[i] = i
    }
    return nil
}

// This function generates a random array of integers with specified number of elements and maximum random value. Numbers can be positive or negative, and there can be duplicates.
//
// Parameters:
//   - arr: the array to generate random integers in
//   - numElems: the number of elements in the array
//   - randMax: the maximum value for random integers
//
// Returns:
//   an error if the array is nil, numElems is less than 1, or randMax is less than 1
//
// Example:
//  arr := make([]int, 10)
//  err := generateRandomInputArr(arr, 10, 100)
//  if err != nil {
//      fmt.Println(err)
//  } else {
//      fmt.Println(arr)
//  }
func generateRandomInputArrImproved(arr []int, numElems, randMax int) error {
    if arr == nil {
        return fmt.Errorf("array is nil")
    }
    if numElems < 1 {
        return fmt.Errorf("number of elements is less than 1")
    }
    if randMax < 1 {
        return fmt.Errorf("randMax is less than 1")
    }

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < numElems; i++ {
        signFlag := rand.Intn(2)
        randNum := rand.Intn(randMax)
        if signFlag%2 == 0 {
            arr[i] = randNum
        } else {
            arr[i] = -randNum
        }
    }
    return nil
}

// This function generates a growing array of integers with specified number of elements. It fills a given array with increasing numbers,
// starting with 0 and increasing by 1 each time.
//
// Parameters:
//   - arr: the array to generate growing integers in
//   - numElems: the number of elements in the array
//
// Returns:
//   an error if the array is nil or numElems is less than 1
//
// Example:
//  arr := make([]int, 10)
//  err := generateGrowingArr(arr, 10)
//  if err != nil {
//      fmt.Println(err)
//  } else {
//      fmt.Println(arr)
//  }
func generateGrowingArrImproved(arr []int, numElems int) error {
    if arr == nil {
        return fmt.Errorf("array is nil")
    }
    if numElems < 1 {
        return fmt.Errorf("number of elements is less than 1")
    }

    for i := 0; i < numElems; i++ {
        arr[i] = i % (numElems / 2)  // Modulo operation allows repetition of elements
    }
    return nil
}


// This function generates a random array of integers with specified number of elements, maximum random value, and number of duplicates.
//
// Parameters:
//   - arr: the array to generate random integers in
//   - numElems: the number of elements in the array
//   - randMax: the maximum value for random integers
//   - numDuplicates: the number of duplicates for each random integer
//
// Returns:
//   an error if the array is nil, numElems is less than 1, randMax is less than 1, or numDuplicates is less than 1
//
// Example:
//  arr := make([]int, 10)
//  err := generateRandomInputArr(arr, 10, 100, 5)
//  if err != nil {
//      fmt.Println(err)
//  } else {
//      fmt.Println(arr)
//  }
func generateRandomInputArrImproved2(arr []int, numElems, randMax, numDuplicates int) error {
    if arr == nil {
        return fmt.Errorf("array is nil")
    }
    if numElems < 1 {
        return fmt.Errorf("number of elements is less than 1")
    }
    if randMax < 1 {
        return fmt.Errorf("randMax is less than 1")
    }
    if numDuplicates < 1 {
        return fmt.Errorf("numDuplicates is less than 1")
    }

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < numElems; i++ {
        signFlag := rand.Intn(2)
        randNum := rand.Intn(randMax / numDuplicates)
        if signFlag%2 == 0 {
            arr[i] = randNum
        } else {
            arr[i] = -randNum
        }
    }
    return nil
}

// This function generates a growing array of integers with specified number of elements and number of duplicates.
//
// Parameters:
//   - arr: the array to generate growing integers in
//   - numElems: the number of elements in the array
//   - numDuplicates: the number of duplicates for each growing integer
//
// Returns:
//   an error if the array is nil, numElems is less than 1, or numDuplicates is less than 1
//
// Example:
//  arr := make([]int, 10)
//  err := generateGrowingArr(arr, 10, 5)
//  if err != nil {
//      fmt.Println(err)
//  } else {
//      fmt.Println(arr)
//  }
func generateGrowingArrImproved2(arr []int, numElems, numDuplicates int) error {
    if arr == nil {
        return fmt.Errorf("array is nil")
    }
    if numElems < 1 {
        return fmt.Errorf("number of elements is less than 1")
    }
    if numDuplicates < 1 {
        return fmt.Errorf("numDuplicates is less than 1")
    }

    for i := 0; i < numElems; i++ {
        arr[i] = i % (numElems / numDuplicates)
    }
    return nil
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
	
	/*----------------------------------- HUGE SETS TESTING -----------------------------------*/

	fmt.Printf("/*------------------- HUGE SETS TESTING -------------------*/\n")

	var huge_input_arr1 = make([]int, 10)
    err := generateRandomInputArr(huge_input_arr1, 10, 100)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    // fmt.Println("Random input array:", huge_input_arr1)

    var huge_input_arr2 = make([]int, 10)
    err = generateGrowingArr(huge_input_arr2, 10)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    // fmt.Println("Growing array:", huge_input_arr2)

	/* HUGE SETS TESTING WITH NAIVE ALGORITHM */

	fmt.Printf("HUGE SETS TESTING WITH NAIVE ALGORITHM\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateRandomInputArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output1 := removeDuplicates(huge_input_arr1)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr1)
	fmt.Printf("Filtered array: %v\n", output1)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH NAIVE ALGORITHM END\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateGrowingArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output2 := removeDuplicates(huge_input_arr2)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr2)
	fmt.Printf("Filtered array: %v\n", output2)
	fmt.Printf("Execution time: %v\n", duration)

	/* HUGE SETS TESTING WITH IMPROVED ALGORITHM */

	fmt.Printf("HUGE SETS TESTING WITH IMPROVED ALGORITHM\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateRandomInputArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output3 := removeDuplicatesImproved(huge_input_arr1)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr1)
	fmt.Printf("Filtered array: %v\n", output3)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH IMPROVED ALGORITHM END\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateGrowingArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output4 := removeDuplicatesImproved(huge_input_arr2)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr2)
	fmt.Printf("Filtered array: %v\n", output4)
	fmt.Printf("Execution time: %v\n", duration)

	fmt.Printf("HUGE SETS TESTING WITH IMPROVED ALGORITHM END\n\n")

	/* HUGE SETS TESTING WITH HASH TABLE ALGORITHM */

	fmt.Printf("HUGE SETS TESTING WITH HASH TABLE ALGORITHM\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateRandomInputArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output5 := removeDuplicatesHashTable(huge_input_arr1)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr1)
	fmt.Printf("Filtered array: %v\n", output5)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH HASH TABLE ALGORITHM END\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateGrowingArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output6 := removeDuplicatesHashTable(huge_input_arr2)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr2)
	fmt.Printf("Filtered array: %v\n", output6)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH HASH TABLE ALGORITHM END\n\n")

	/* HUGE SETS TESTING WITH HASH TABLE DYNAMIC ALGORITHM */

	fmt.Printf("HUGE SETS TESTING WITH HASH TABLE DYNAMIC ALGORITHM START\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateRandomInputArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output7 := removeDuplicatesDynamicHashTable(huge_input_arr1)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr1)
	fmt.Printf("Filtered array: %v\n", output7)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH HASH TABLE DYNAMIC ALGORITHM END\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateGrowingArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output8 := removeDuplicatesDynamicHashTable(huge_input_arr2)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr2)
	fmt.Printf("Filtered array: %v\n", output8)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH HASH TABLE DYNAMIC ALGORITHM END\n\n")

	/* HUGE SETS TESTING WITH BIT MAP ALGORITHM */

	fmt.Printf("HUGE SETS TESTING WITH BIT MAP ALGORITHM START\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateRandomInputArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output9 := removeDuplicatesBitHashTable(huge_input_arr1)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr1)
	fmt.Printf("Filtered array: %v\n", output9)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH BIT MAP ALGORITHM END\n\n")

	fmt.Printf("GENERATED RANDOM HUGE SET WITH generateGrowingArr() ALGORITHM\n")

	// Capture the time before executing removeDuplicates
	startTime = time.Now()

	// Execute removeDuplicates
	output10 := removeDuplicatesBitHashTable(huge_input_arr2)

	// Capture the time after executing removeDuplicates
	endTime = time.Now()

	// Calculate the execution duration
	duration = endTime.Sub(startTime)

	// displaying arrays
	fmt.Printf("Original array: %v\n", huge_input_arr2)
	fmt.Printf("Filtered array: %v\n", output10)
	fmt.Printf("Execution time: %v\n", duration)
	fmt.Printf("HUGE SETS TESTING WITH BIT MAP ALGORITHM END\n\n")

	/* TESTING NEW GENERATOR METHODS */

	fmt.Println("random arrays")

	// Generate a random array of integers with 10 elements and a maximum random value of 100
    arr1 := make([]int, 10)
    err1 := generateRandomInputArrImproved(arr1,10, 100)
    if err1 != nil {
        fmt.Println(err1)
    } else {
        fmt.Println(arr1)
    }

    // Generate a growing array of integers with 10 elements
    arr2 := make([]int, 10)
    err2 := generateGrowingArrImproved(arr2, 10)
    if err2 != nil {
        fmt.Println(err2)
    } else {
        fmt.Println(arr2)
    }

    // Generate a random array of integers with 10 elements, a maximum random value of 100, and 5 duplicates
    arr3 := make([]int, 10)
    err3 := generateRandomInputArrImproved2(arr3, 10, 100, 5)
    if err3 != nil {
        fmt.Println(err3)
    } else {
        fmt.Println(arr3)
    }

    // Generate a growing array of integers with 10 elements and 5 duplicates
    arr4 := make([]int, 10)
    err4 := generateGrowingArrImproved2(arr4, 10, 5)
    if err4 != nil {
        fmt.Println(err4)
    } else {
        fmt.Println(arr4)
    }

}
