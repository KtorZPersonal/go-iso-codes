// Copyright © 2016 Matthias Benkort
// Use of this source code is governed by the MPL 2.0 license that can be found in the LICENSE file.

package currencies

import "strings"

// Currency materializes a currency object, with its name and ISO identifiers
type Currency struct {
	// Alpha3 refers to a 3-letter alpha ISO 4217 Code
	Alpha3 string
	// Numeric refers to a 3-digit numeric ISO 4217 Code
	Numeric string
	// Decimals refers to the number of decimal accepted in this currency
	Decimals uint
	// Name refers to the common english name
	Name string
}

// Currencies exports all currencies available in the package
var Currencies = reduceCurrencies(currencies)

func reduceCurrencies(dict map[string]Currency) (xs []Currency) {
	for _, currency := range dict {
		xs = append(xs, currency)
	}

	return xs
}

// Lookup retrieves a list of currencies satisfying the criteria(s) given in arguments
func Lookup(query Currency) []Currency {
	if query.Alpha3 != "" {
		res, ok := currencies[query.Alpha3]
		if ok {
			return []Currency{res}
		}

		return make([]Currency, 0)
	}

	if query.Numeric != "" {
		for _, res := range currencies {
			if query.Numeric == res.Numeric {
				return []Currency{res}
			}
		}
	}

	results := make([]Currency, 0)
	if query.Name != "" {
		for _, res := range currencies {
			if strings.Contains(res.Name, query.Name) {
				results = append(results, res)
			}
		}
	}
	return results
}
