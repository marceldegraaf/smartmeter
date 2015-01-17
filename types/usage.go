package types

import "time"

type Usage struct {
	// Power delivered to the meter by the energy provider
	TotalDeliveredLowTariff  float64 `bson:"tdl" json:"total_delivered_low_tariff"`
	TotalDeliveredHighTariff float64 `bson:"tdh" json:"total_delivered_high_tariff"`

	// Power provided back to the energy provider
	// e.g. through solar cells
	TotalProvidedLowTariff  float64 `bson:"tpl" json:"total_provided_low_tariff"`
	TotalProvidedHighTariff float64 `bson:"tph" json:"total_provided_high_tariff"`

	// High or low tariff (1 = low, 2 = high)
	CurrentTariff int64 `bson:"ct" json:"current_tariff"`

	// Real time power delivered and provided
	CurrentlyDelivered float64 `bson:"cd" json:"currently_delivered"`
	CurrentlyProvided  float64 `bson:"cp" json:"currently_provided"`

	Timestamp time.Time `bson:"ts" json:"timestamp"`
}
