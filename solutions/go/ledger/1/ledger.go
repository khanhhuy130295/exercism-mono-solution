package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	dateWidth        = 10
	descriptionWidth = 25
	changeWidth      = 13
)

var (
	errInvalidCurrency = errors.New("unsupported currency")
	errInvalidLocale   = errors.New("unsupported locale")
	errInvalidDate     = errors.New("invalid date")
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type ledgerFormat struct {
	header         string
	dateLayout     string
	currencyGap    string
	thousands      string
	decimal        string
	positiveSuffix string
	negative       func(symbol string, currencyGap string, cents string) string
}

var formats = map[string]ledgerFormat{
	"en-US": {
		header:         "Date       | Description               | Change       ",
		dateLayout:     "01/02/2006",
		thousands:      ",",
		decimal:        ".",
		negative:       func(symbol string, _ string, cents string) string { return "(" + symbol + cents + ")" },
		currencyGap:    "",
		positiveSuffix: " ",
	},
	"nl-NL": {
		header:     "Datum      | Omschrijving              | Verandering  ",
		dateLayout: "02-01-2006",
		thousands:  ".",
		decimal:    ",",
		negative: func(symbol string, currencyGap string, cents string) string {
			return symbol + currencyGap + "-" + cents + " "
		},
		currencyGap:    " ",
		positiveSuffix: " ",
	},
}

var currencySymbols = map[string]string{
	"EUR": "€",
	"USD": "$",
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	format, ok := formats[locale]
	if !ok {
		return "", errInvalidLocale
	}

	symbol, ok := currencySymbols[currency]
	if !ok {
		return "", errInvalidCurrency
	}

	sortedEntries := append([]Entry(nil), entries...)
	sort.Slice(sortedEntries, func(i, j int) bool {
		left, right := sortedEntries[i], sortedEntries[j]
		if left.Date != right.Date {
			return left.Date < right.Date
		}
		if left.Description != right.Description {
			return left.Description < right.Description
		}
		return left.Change < right.Change
	})

	var output strings.Builder
	output.WriteString(format.header)
	output.WriteByte('\n')
	for _, entry := range sortedEntries {
		line, err := formatEntry(entry, symbol, format)
		if err != nil {
			return "", err
		}
		output.WriteString(line)
	}

	return output.String(), nil
}

func formatEntry(entry Entry, symbol string, format ledgerFormat) (string, error) {
	date, err := time.Parse("2006-01-02", entry.Date)
	if err != nil {
		return "", fmt.Errorf("%w: %q", errInvalidDate, entry.Date)
	}

	formattedDate := date.Format(format.dateLayout)
	description := formatDescription(entry.Description)
	change := formatChange(entry.Change, symbol, format)

	return fmt.Sprintf("%-*s | %-*s | %*s\n", dateWidth, formattedDate, descriptionWidth, description, changeWidth, change), nil
}

func formatDescription(description string) string {
	characters := []rune(description)
	if len(characters) <= descriptionWidth {
		return description
	}

	return string(characters[:descriptionWidth-3]) + "..."
}

func formatChange(change int, symbol string, format ledgerFormat) string {
	negative := change < 0
	magnitude := uint64(change)
	if negative {
		// -(MinInt) overflows, so negate after moving one step toward zero.
		magnitude = uint64(-(change + 1)) + 1
	}

	formattedCents := formatCents(magnitude, format.thousands, format.decimal)
	if negative {
		return format.negative(symbol, format.currencyGap, formattedCents)
	}

	return symbol + format.currencyGap + formattedCents + format.positiveSuffix
}

func formatCents(cents uint64, thousands string, decimal string) string {
	whole, fraction := cents/100, cents%100
	wholeText := fmt.Sprintf("%d", whole)
	wholeText = groupDigits(wholeText, thousands)

	return fmt.Sprintf("%s%s%02d", wholeText, decimal, fraction)
}

func groupDigits(digits string, separator string) string {
	if len(digits) <= 3 {
		return digits
	}

	firstGroupLength := len(digits) % 3
	if firstGroupLength == 0 {
		firstGroupLength = 3
	}

	var grouped strings.Builder
	grouped.WriteString(digits[:firstGroupLength])
	for position := firstGroupLength; position < len(digits); position += 3 {
		grouped.WriteString(separator)
		grouped.WriteString(digits[position : position+3])
	}

	return grouped.String()
}
