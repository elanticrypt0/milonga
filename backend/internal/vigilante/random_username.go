package vigilante

import (
	"math/rand"
	"strings"
	"time"
)

// GenerateUsername creates a random username based on provided parameters
func GenerateUsername(minLength, maxLength int) string {
	// Initialize a new random number generator with current time as seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define character sets
	letters := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialChars := "_-"

	// Combine all possible characters
	allCharacters := letters + numbers + specialChars

	// Determine random length between minLength and maxLength
	length := r.Intn(maxLength-minLength+1) + minLength

	// Build the username
	var username strings.Builder

	// Ensure first character is a letter
	username.WriteByte(letters[r.Intn(len(letters))])

	// Generate the rest of the username
	for i := 1; i < length; i++ {
		username.WriteByte(allCharacters[r.Intn(len(allCharacters))])
	}

	return username.String()
}
