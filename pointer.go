package main

import "fmt"

type bill struct {
	names string
	items map[string]float64
	tip   float64
}

func Newbill(name string) bill {
	b := bill{
		names: name,
		items: map[string]float64{
			"pie":  5.99,
			"cake": 3.99,
		},
		tip: 0,
	}
	return b
}

func (b *bill) Format() string {
	fs := "Bill breakdwon \n"
	var total float64 = 0

	for indeks, value := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", indeks+":", value)
		total += value
	}

	fs += fmt.Sprintf("%v ...$%0.2f", "total:", total)
	return fs
}
