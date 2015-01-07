package bitcointransaction

import (
    "testing"

    "github.com/stripe/stripe-go/currency"
    "github.com/stripe/stripe-go/bitcoinreceiver"
    stripe "github.com/stripe/stripe-go"
    . "github.com/stripe/stripe-go/utils"
)

func init() {
    stripe.Key = GetTestKey()
}

func TestBitcoinTransactionGet(t *testing.T) {
    bitcoinReceiverParams := &stripe.BitcoinReceiverParams{
        Amount:     1000,
        Currency:   currency.USD,
        Email:      "do+fill_now@email.address",
        Desc:       "some details",
    }

    receiver, _ := bitcoinreceiver.New(bitcoinReceiverParams)

    receiverRes, _ := bitcoinreceiver.Get(receiver.ID, nil)

    if len(receiverRes.Transactions.Values) < 1 {
        t.Errorf("Expected receiver to have at least one transaction\n")
    }

    res := receiverRes.Transactions.Values[0]

    target, err := Get(res.ID, nil)

    if err != nil {
        t.Error(err)
    }

    if target.ID != res.ID {
        t.Errorf("BitcoinTransaction id %q does not match expected id %q\n", target.ID, res.ID)
    }
}

func TestBitcoinTransactionList(t *testing.T) {
    params := &stripe.BitcoinTransactionListParams{}
    params.Filters.AddFilter("include[]", "", "total_count")
    params.Filters.AddFilter("limit", "", "5")
    params.Single = true

    i := List(params)
    for i.Next() {
        if i.BitcoinTransaction() == nil {
            t.Error("No nil values expected")
        }

        if i.Meta() == nil {
            t.Error("No metadata returned")
        }
    }
    if err := i.Err(); err != nil {
        t.Error(err)
    }
}