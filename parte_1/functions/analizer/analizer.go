package main

const (
	defaultCounters = 1
)

type Response struct {
	Total           float32        `json:"total"`
	NotPurchased    int            `json:"not_purchased"`
	HighestPurchase float32        `json:"highest_purchase"`
	PurchasesTDC    map[string]int `json:"purchases_tdc"`
}

func NewResponse() *Response {
	return &Response{
		PurchasesTDC: make(map[string]int),
	}
}

func (r *Response) analizer(inputSales Sales) {

	for _, record := range inputSales {
		if !record.Compro {
			r.NotPurchased++
			continue
		}

		r.Total += record.Monto

		if record.Monto > r.HighestPurchase {
			r.HighestPurchase = record.Monto
		}

		if _, ok := r.PurchasesTDC[record.TDC]; ok {
			r.PurchasesTDC[record.TDC]++
			continue
		}

		r.PurchasesTDC[record.TDC] = defaultCounters

	}
}
