// Copyright © 2016 Matthias Benkort
// Use of this source code is governed by the MPL 2.0 license that can be found in the LICENSE file.

package currencies

import "fmt"

// sort ascendently sort a slice of Currency on the Alpha3
func sort(xs []Currency) []Currency {
	var _xs []Currency
sort:
	for _, a := range xs {
		for i, b := range _xs {
			if a.Alpha3 < b.Alpha3 {
				_xs = append(_xs[:i], append([]Currency{a}, _xs[i:]...)...)
				continue sort
			}
		}

		_xs = append(_xs, a)
	}

	return _xs
}

// ExampleLookup illustrates several lookups
func ExampleLookup() {
	results := Lookup(Currency{Alpha3: "EUR"})
	fmt.Println(results)

	results = Lookup(Currency{Name: "kron"})
	fmt.Println(sort(results))

	// Output:
	// [{EUR 978 2 Euro}]
	// [{DKK 208 2 Danish krone} {NOK 578 2 Norwegian krone} {SEK 752 2 Swedish krona/kronor}]
}
