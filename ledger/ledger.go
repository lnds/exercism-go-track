package ledger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry
	// step 1 remove loop
	entriesCopy = append(entriesCopy, entries...)

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	// step 2 remove m1 & m2 maps
	es := entriesCopy
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if e.Date < first.Date || e.Description < first.Description || e.Change < first.Change {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}

	// step 3: initialization of s based on locale refactored in header function
	var s string
	s, err := header(locale)
	if err != nil {
		return s, err
	}
	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			// step 4: pasrse date
			d, err := formatDate(entry.Date, locale)

			if err != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: err}
			}

			// step 5: format description
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			}
			de = fmt.Sprintf("%-25s", de)

			// step 6: format change
			a, err := formatChange(entry.Change, currency, locale)
			if err != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: err}
			}
			var al int
			for range a {
				al++
			}
			co <- struct {
				i int
				s string
				e error
			}{i: i, s: fmt.Sprintf("%10s", d) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
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
		return "", errors.New("")
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

	var curSymbol string
	switch currency {
	case "EUR":
		curSymbol = "€"
	case "USD":
		curSymbol = "$"
	default:
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
		return "", nil
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
