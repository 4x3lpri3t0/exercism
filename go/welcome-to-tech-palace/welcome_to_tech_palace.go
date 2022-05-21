package techpalace

import (
	"strings"
	"unicode"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	ret := strings.Repeat("*", numStarsPerLine)
	ret += "\n" + welcomeMsg + "\n"
	ret += strings.Repeat("*", numStarsPerLine)
	return ret
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimFunc(oldMsg, func(r rune) bool {
		return r == '*' || unicode.IsSpace(r)
	})
}
