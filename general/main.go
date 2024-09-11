package general

import (
	"fmt"
	"os"
)

// Define the character set
const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

// Global variable to count the number of combinations generated
var count int

// Function to generate combinations and write the first 10 results to a file
func generateCombinations(length int, current string, file *os.File) {
	if count >= 10 {
		return
	}
	if len(current) == length {
		_, err := file.WriteString(current + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		count++
		return
	}
	for _, char := range charset {
		generateCombinations(length, current+string(char), file)
	}
}

func main() {
	length := 3 // Adjust this to the desired length
	file, err := os.Create("combinations.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Starting generation of combinations...")
	generateCombinations(length, "", file)
	fmt.Println("Combination generation completed. Results are saved in 'combinations.txt'.")
}
