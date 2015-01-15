package parser

import (
	"regexp"
	"strconv"

	"github.com/marceldegraaf/smartmeter/log"
	"github.com/marceldegraaf/smartmeter/types"
)

var (
	payload                   string
	deliveredLowTariffRegexp  = regexp.MustCompile("1.8.1\\(([0-9.]+)")
	deliveredHighTariffRegexp = regexp.MustCompile("1.8.2\\(([0-9.]+)")
	providedLowTariffRegexp   = regexp.MustCompile("2.8.1\\(([0-9.]+)")
	providedHighTariffRegexp  = regexp.MustCompile("2.8.2\\(([0-9.]+)")
	currentTariffRegexp       = regexp.MustCompile("96.14.0\\(([0-9]+)")
	currentlyDeliveredRegexp  = regexp.MustCompile("1.7.0\\(([0-9.]+)")
	currentlyProvidedRegexp   = regexp.MustCompile("2.7.0\\(([0-9.]+)")
)

func Parse(telegram types.Telegram) {
	payload = telegram.Payload.(string)

	usage := types.Usage{
		TotalDeliveredLowTariff:  matchAsFloat(deliveredLowTariffRegexp),
		TotalDeliveredHighTariff: matchAsFloat(deliveredHighTariffRegexp),
		TotalProvidedLowTariff:   matchAsFloat(providedLowTariffRegexp),
		TotalProvidedHighTariff:  matchAsFloat(providedHighTariffRegexp),
		CurrentTariff:            matchAsInt(currentTariffRegexp),
		CurrentlyDelivered:       matchAsFloat(currentlyDeliveredRegexp),
		CurrentlyProvided:        matchAsFloat(currentlyProvidedRegexp),
	}

	log.Infof("Usage: %#v", usage)
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
