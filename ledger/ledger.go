package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var currencies = map[string]string{
	"EUR": "€",
	"USD": "$",
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {

	if _, ok := currencies[currency]; !ok {
		return "", errors.New("invalid curency")
	}

	var entriesCopy []Entry
	// step 1 remove loop
	entriesCopy = append(entriesCopy, entries...)

	// step 2 remove m1 & m2 maps
	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date < entriesCopy[j].Date {
			return true
		} else if entriesCopy[i].Description < entriesCopy[j].Description {
			return true
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	// step 3: initialization of s based on locale refactored in header function
	s, err := header(locale)
	if err != nil {
		return s, err
	}

	// step 8 remove parallelism
	for _, entry := range entriesCopy {
		// step 4: pasrse date
		date, err := formatDate(entry.Date, locale)
		if err != nil {
			return "", err
		}
		// step 6: format change
		change, err := formatChange(entry.Change, currency, locale)
		if err != nil {
			return "", err
		}

		// step 5: format description
		description := entry.Description
		if len(description) > 25 {
			description = description[:22] + "..."
		}

		s += fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, change)
	}

	return s, nil
}

func header(locale string) (string, error) {
	switch locale {
	case "nl-NL":
		return fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering"), nil
	case "en-US":
		return fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change"), nil
	default:
		return "", errors.New("no header")
	}
}

func formatDate(date, locale string) (string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date, err
	}
	switch locale {
	case "nl-NL":
		return t.Format("02-01-2006"), nil
	case "en-US":
		return t.Format("01/02/2006"), nil
	default:
		return date, errors.New("date for localer not supported")
	}
}

func formatChange(change int, currency, locale string) (string, error) {

	curSymbol, ok := currencies[currency]
	if !ok {
		return "", errors.New("invalid currency")
	}

	negative := false
	cents := change
	if cents < 0 {
		cents = cents * -1
		negative = true
	}

	switch locale {
	case "nl-NL":
		centStr := formatCents(cents, ".", ",", 2)
		if negative {
			return fmt.Sprintf("%s %s-", curSymbol, centStr), nil
		}
		return fmt.Sprintf("%s %s ", curSymbol, centStr), nil
	case "en-US":
		centStr := formatCents(cents, ",", ".", 3)
		if negative {
			return fmt.Sprintf("(%s%s)", curSymbol, centStr), nil
		}
		return fmt.Sprintf(" %s%s ", curSymbol, centStr), nil
	default:
		return "", errors.New("invalid locale")
	}
}

func formatCents(cents int, sep, point string, decs int) string {
	centsStr := fmt.Sprintf("%0*d", decs, cents)
	rest := centsStr[:len(centsStr)-2]
	var parts []string
	for len(rest) > 3 {
		parts = append([]string{rest[len(rest)-3:]}, parts...)
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append([]string{rest}, parts...)
	}
	return fmt.Sprintf("%s%s%s", strings.Join(parts, sep), point, centsStr[len(centsStr)-2:])
}
