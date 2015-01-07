// Package bitcointransaction provides the /bitcoin/transactions APIs
package bitcointransaction

import (
    "net/url"

    stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /charges APIs.
type Client struct {
    B   stripe.Backend
    Key string
}

// Get returns the details of a bitcoin transaction.
// For more details see https://stripe.com/docs/api#retrieve_bitcoin_transaction.
func Get(id string, params *stripe.BitcoinTransactionParams) (*stripe.BitcoinTransaction, error) {
    return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.BitcoinTransactionParams) (*stripe.BitcoinTransaction, error) {
    var commonParams *stripe.Params

    if params != nil {
        commonParams = &params.Params
    }

    charge := &stripe.BitcoinTransaction{}
    err := c.B.Call("GET", "/bitcoin/transactions/"+id, c.Key, nil, commonParams, charge)

    return charge, err
}

// List returns a list of bitcoin transactions.
// For more details see https://stripe.com/docs/api#list_bitcoin_transactions.
func List(params *stripe.BitcoinTransactionListParams) *Iter {
    return getC().List(params)
}

func (c Client) List(params *stripe.BitcoinTransactionListParams) *Iter {
    type receiverList struct {
        stripe.ListMeta
        Values []*stripe.BitcoinTransaction `json:"data"`
    }

    var body *url.Values
    var lp *stripe.ListParams

    if params != nil {
        body = &url.Values{}

        if len(params.Customer) > 0 {
            body.Add("customer", params.Customer)
        }

        if len(params.Receiver) > 0 {
            body.Add("receiver", params.Receiver)
        }

        params.AppendTo(body)
        lp = &params.ListParams
    }

    return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
        list := &receiverList{}
        err := c.B.Call("GET", "/bitcoin/transactions", c.Key, &b, nil, list)

        ret := make([]interface{}, len(list.Values))
        for i, v := range list.Values {
            ret[i] = v
        }

        return ret, list.ListMeta, err
    })}
}

// Iter is an iterator for lists of BitcoinTransactions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
    *stripe.Iter
}

// BitcoinTransaction returns the most recent BitcoinTransaction
// visited by a call to Next.
func (i *Iter) BitcoinTransaction() *stripe.BitcoinTransaction {
    return i.Current().(*stripe.BitcoinTransaction)
}

func getC() Client {
    return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
