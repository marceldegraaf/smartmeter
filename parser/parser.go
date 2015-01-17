package parser

import (
	"regexp"
	"strconv"
	"time"

	"github.com/marceldegraaf/smartmeter/log"
	"github.com/marceldegraaf/smartmeter/types"
)

// Note that we ignore the summer/winter time indicator
// that is present in the IEC 62056-21 timestamp spec.
const timeFormat = "060102150405"

var (
	payload                   string
	deliveredLowTariffRegexp  = regexp.MustCompile("1.8.1\\(([0-9.]+)")
	deliveredHighTariffRegexp = regexp.MustCompile("1.8.2\\(([0-9.]+)")
	providedLowTariffRegexp   = regexp.MustCompile("2.8.1\\(([0-9.]+)")
	providedHighTariffRegexp  = regexp.MustCompile("2.8.2\\(([0-9.]+)")
	currentTariffRegexp       = regexp.MustCompile("96.14.0\\(([0-9]+)")
	currentlyDeliveredRegexp  = regexp.MustCompile("1.7.0\\(([0-9.]+)")
	currentlyProvidedRegexp   = regexp.MustCompile("2.7.0\\(([0-9.]+)")
	timestampRegexp           = regexp.MustCompile("1.0.0\\(([0-9]+)")

	Incoming = make(chan types.Usage, 16)
)

func Parse(telegram types.Telegram) {
	payload = telegram.Payload.(string)

	// TODO: calling the converion functions directly, and letting
	// them return empty values when they occur an error, can result
	// in bogus usage data being stored. We should probably handle
	// parsing/conversion errors more decently, and fail parsing an
	// item instead of passing it on with empty values as if it were
	// parsed.
	usage := types.Usage{
		TotalDeliveredLowTariff:  matchAsFloat(deliveredLowTariffRegexp),
		TotalDeliveredHighTariff: matchAsFloat(deliveredHighTariffRegexp),
		TotalProvidedLowTariff:   matchAsFloat(providedLowTariffRegexp),
		TotalProvidedHighTariff:  matchAsFloat(providedHighTariffRegexp),
		CurrentTariff:            matchAsInt(currentTariffRegexp),
		CurrentlyDelivered:       matchAsFloat(currentlyDeliveredRegexp),
		CurrentlyProvided:        matchAsFloat(currentlyProvidedRegexp),
		Timestamp:                parsedTimestamp(timestampRegexp),
	}

	Incoming <- usage

}

func findFirstMatch(r *regexp.Regexp) string {
	match := r.FindStringSubmatch(payload)

	if len(match) == 0 {
		return ""
	}

	return match[1]
}

func matchAsInt(r *regexp.Regexp) int64 {
	match := findFirstMatch(r)

	i, err := strconv.ParseInt(match, 0, 0)
	if err != nil {
		log.Errorf("Could not parse %q to int: %s", match, err)
		return 0
	}

	return i
}

func matchAsFloat(r *regexp.Regexp) float64 {
	match := findFirstMatch(r)

	flt, err := strconv.ParseFloat(match, 64)
	if err != nil {
		log.Errorf("Could not parse %q to float64: %s", match, err)
		return 0
	}

	return flt
}

func parsedTimestamp(r *regexp.Regexp) time.Time {
	match := findFirstMatch(r)

	t, err := time.Parse(timeFormat, match)
	if err != nil {
		log.Errorf("Could not parse time %q: %s", match, err)
		return time.Now()
	}

	return t
}
