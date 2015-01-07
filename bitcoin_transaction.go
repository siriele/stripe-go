package stripe

import "encoding/json"

// BitcoinTransactionListParams is the set of parameters that can be used when listing BitcoinTransactions.
// For more details see https://stripe.com/docs/api/#list_bitcoin_transactions.
type BitcoinTransactionListParams struct {
	ListParams
	Customer	string
	Receiver	string
}

// BitcoinTransactionParams is the set of parameters that can be used when retrieving a BitcoinTransaction.
// For more details see https://stripe.com/docs/api/#retrieve_bitcoin_transaction.
type BitcoinTransactionParams struct {
	Params
}

// BitcoinTransactionList is a list object for BitcoinTransactions.
// For more details see https://stripe.com/docs/api/#list_bitcoin_transactions
type BitcoinTransactionList struct {
	ListMeta
	Values []*BitcoinTransaction `json:"data"`
}

// BitcoinTransaction is the resource representing a Stripe bitcoin transaction.
// For more details see https://stripe.com/docs/api/#bitcoin_transactions
type BitcoinTransaction struct {
	ID				string		`json:"id"`
	Created			int64		`json:"created"`
	Amount			uint64		`json:"amount"`
	Currency		Currency	`json:"currency"`
	BitcoinAmount	uint64		`json:"bitcoin_amount"`
	Receiver		string		`json:"receiver"`
	Customer		string		`json:"customer"`
}

// UnmarshalJSON handles deserialization of a BitcoinTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (bt *BitcoinTransaction) UnmarshalJSON(data []byte) error {
	type bitcoinTransaction BitcoinTransaction
	var t bitcoinTransaction
	err := json.Unmarshal(data, &t)
	if err == nil {
		*bt = BitcoinTransaction(t)
	} else {
		// the id is surrounded by "\" characters, so strip them
		bt.ID = string(data[1 : len(data)-1])
	}

	return nil
}