package exchanges

import (
	"reflect"
	"testing"
)

func TestGetAllExchanges(t *testing.T) {
	t.Parallel()

	excpectedExchanges := []string{
		"Binance",
		"Bittrex",
		"Coinbase",
	}

	if !reflect.DeepEqual(excpectedExchanges, GetAllExchanges()) {
		t.Errorf("Expected exchanges slice does not match the slice returned from GetAllExchanges()")
	}
}
