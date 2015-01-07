package bitcoinreceiver

import (
	"testing"

	"github.com/stripe/stripe-go/currency"
	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestBitcoinReceiverNew(t *testing.T) {
	bitcoinReceiverParams := &stripe.BitcoinReceiverParams{
		Amount:		1000,
		Currency:	currency.USD,
		Email:		"a@b.com",
		Desc:		"some details",
	}

	target, err := New(bitcoinReceiverParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != bitcoinReceiverParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, bitcoinReceiverParams.Amount)
	}

	if target.Currency != bitcoinReceiverParams.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, bitcoinReceiverParams.Currency)
	}

	if target.Desc != bitcoinReceiverParams.Desc {
		t.Errorf("Desc %q does not match expected description %v\n", target.Desc, bitcoinReceiverParams.Desc)
	}

	if target.Email != bitcoinReceiverParams.Email {
		t.Errorf("Email %q does not match expected email %v\n", target.Email, bitcoinReceiverParams.Email)
	}
}

func TestBitcoinReceiverGet(t *testing.T) {
	bitcoinReceiverParams := &stripe.BitcoinReceiverParams{
		Amount:		1000,
		Currency:	currency.USD,
		Email:		"a@b.com",
		Desc:		"some details",
	}

	res, _ := New(bitcoinReceiverParams)

	target, err := Get(res.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if target.ID != res.ID {
		t.Errorf("BitcoinReceiver id %q does not match expected id %q\n", target.ID, res.ID)
	}
}

func TestBitcoinReceiverList(t *testing.T) {
	params := &stripe.BitcoinReceiverListParams{}
	params.Filters.AddFilter("include[]", "", "total_count")
	params.Filters.AddFilter("limit", "", "5")
	params.Single = true

	i := List(params)
	for i.Next() {
		if i.BitcoinReceiver() == nil {
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