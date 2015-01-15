package types

type Telegram struct {
	Payload interface{}
}

type Usage struct {
	// Power delivered to the meter by the energy provider
	TotalDeliveredLowTariff  float64
	TotalDeliveredHighTariff float64

	// Power provided back to the energy provider
	// e.g. through solar cells
	TotalProvidedLowTariff  float64
	TotalProvidedHighTariff float64

	// High or low tariff (1 = low, 2 = high)
	CurrentTariff int64

	// Real time power delivered and provided
	CurrentlyDelivered float64
	CurrentlyProvided  float64
}
