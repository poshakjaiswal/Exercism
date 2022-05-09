package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	//panic("Please implement the WelcomeMessage() function")
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	//panic("Please implement the AddBorder() function")
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	replaced := strings.ReplaceAll(oldMsg, "*", " ")
	cleanup := strings.TrimSpace(replaced)
	//panic("Please implement the CleanupMessage() function")
	return cleanup
}
