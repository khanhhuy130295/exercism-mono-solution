package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var builder strings.Builder
	var star string = strings.Repeat("*", numStarsPerLine)
	builder.WriteString(star)
	builder.WriteString("\n")
	builder.WriteString(welcomeMsg)
	builder.WriteString("\n")
	builder.WriteString(star)
	return builder.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var new_str string = strings.ReplaceAll(oldMsg, "*", "")
	new_str = strings.TrimSpace(new_str)
	return new_str
}
