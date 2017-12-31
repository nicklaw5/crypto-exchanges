package exchanges

// A list of cryptocurrency exchanges.
const (
	Binance  = "Binance"
	Bittrex  = "Bittrex"
	Coinbase = "Coinbase"
)

// GetAllExchanges returns a list of all cryptocurrency exchanges.
func GetAllExchanges() []string {
	return []string{
		Binance,
		Bittrex,
		Coinbase,
	}
}
