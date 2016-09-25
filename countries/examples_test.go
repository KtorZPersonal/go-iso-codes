// Copyright © 2016 Matthias Benkort
// Use of this source code is governed by the MPL 2.0 license that can be found in the LICENSE file.

package countries

import (
	"fmt"

	"github.com/KtorZ/go-iso-codes/currencies"
	"github.com/KtorZ/go-iso-codes/languages"
)

// sort ascendently sort a slice of countries on the Alpha3
func sort(xs []Country) []Country {
	var _xs []Country
sort:
	for _, a := range xs {
		for i, b := range _xs {
			if a.Alpha3 < b.Alpha3 {
				_xs = append(_xs[:i], append([]Country{a}, _xs[i:]...)...)
				continue sort
			}
		}

		_xs = append(_xs, a)
	}

	return _xs
}

// ExampleLookup illustrates several lookups
func ExampleLookup() {
	results := Lookup(Country{Alpha3: "NLD"})
	fmt.Println(results)

	results = Lookup(Country{
		Languages: []languages.Language{
			languages.Language{Alpha3: "nld"},
		},
		Currencies: []currencies.Currency{
			currencies.Currency{Alpha3: "EUR"},
		},
	})
	fmt.Println(sort(results))

	// Output:
	// [{NL NLD 528 [{EUR 978 2 Euro}] [+31] [{nl nld Dutch}] Netherlands}]
	// [{BE BEL 056 [{EUR 978 2 Euro}] [+32] [{nl nld Dutch} {fr fra French} {de deu German}] Belgium} {NL NLD 528 [{EUR 978 2 Euro}] [+31] [{nl nld Dutch}] Netherlands}]
}
