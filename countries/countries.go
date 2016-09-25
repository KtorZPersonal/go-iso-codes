// Copyright © 2016 Matthias Benkort
// Use of this source code is governed by the MPL 2.0 license that can be found in the LICENSE file.

package countries

import (
	"strings"

	"github.com/KtorZ/go-iso-codes/currencies"
	"github.com/KtorZ/go-iso-codes/languages"
)

// Country materializes a country object, with its name and ISO identifiers
type Country struct {
	// Alpha2 refers to a 2-letter alpha ISO 3166-1 Code
	Alpha2 string `json:"alpha2"`
	// Alpha3 refers to a 3-letter alpha ISO 3166-1 Code
	Alpha3 string `json:"alpha3"`
	// Numeric refers to a 3-digit numeric ISO 3166-1 Code
	Numeric string `json:"numeric"`
	// Currencies refers to all currencies used in a country
	Currencies []currencies.Currency `json:"currencies"`
	// DialingCodes refers to all dialing / calling phone codes / prefixes
	DialingCodes []string `json:"dialing_codes"`
	// Languages refers to all languages spoken in a country
	Languages []languages.Language `json:"languages"`
	// Name refers to the common english name of a country
	Name string `json:"name"`
}

// Countries exports all countries available in the package
var Countries = reduceCountries(countries)

// country is anb opaque type to define the set of testdata without repeating currencies and
// languages every time.
type country struct {
	Country
	Currencies []string
	Languages  []string
}

// Lookup retrieves a list of countries satisfying the criteria(s) given in arguments
//
// Criteria such as `Alpha3`, `Alpha2` and `Numeric` return a unique Country, if valid
// Criteria `Name` is matched against inclusion (e.g. "United Kingdom" will be matched by "United"
// Criteria `Currencies` and `Languages` use Lookup() function for the related type
func Lookup(query Country) []Country {
	if query.Alpha3 != "" {
		res, ok := countries[query.Alpha3]
		if ok {
			return []Country{toCountry(res)}
		}

		return make([]Country, 0)
	}

	if query.Alpha2 != "" {
		for _, res := range Countries {
			if query.Alpha2 == res.Alpha2 {
				return []Country{res}
			}
		}
	}

	if query.Numeric != "" {
		for _, res := range Countries {
			if query.Numeric == res.Numeric {
				return []Country{res}
			}
		}
	}

	// Other criteria might return several countries, we build a pipeline of goroutine to
	// filter all countries according to one criteria at a time
	chCountry := make(chan Country, len(Countries))
	go func(chOut chan Country) {
		for _, country := range Countries {
			chOut <- country
		}

		close(chOut)
	}(chCountry)

	chCountry = filterName(query.Name, chCountry)
	chCountry = filterDialingCode(query.DialingCodes, chCountry)
	chCountry = filterLanguages(query.Languages, chCountry)
	chCountry = filterCurrencies(query.Currencies, chCountry)

	results := make([]Country, 0)
	for res := range chCountry {
		results = append(results, res)
	}

	return results
}

// filterCurrencies constructs the filter for currency matching
func filterCurrencies(queries []currencies.Currency, chIn <-chan Country) (chOut chan Country) {
	var queryCurrencies []currencies.Currency
	for _, query := range queries {
		queryCurrencies = mergeCurrencies(queryCurrencies, currencies.Lookup(query))
	}
	queryLength := len(queryCurrencies)

	chOut = make(chan Country)

	go func() {
		for in := range chIn {
			if queryLength == 0 {
				chOut <- in
				continue
			}

			nbFound := 0

			for _, query := range queryCurrencies {
				for _, in := range in.Currencies {
					if in.Numeric == query.Numeric {
						nbFound++
					}
				}
			}

			if nbFound == queryLength {
				chOut <- in
			}
		}

		close(chOut)
	}()

	return chOut
}

// filterDialingCode constructs the filter for dialing matching
func filterDialingCode(query []string, chIn <-chan Country) (chOut chan Country) {
	queryLength := len(query)

	chOut = make(chan Country)

	go func() {
		for in := range chIn {
			if queryLength == 0 {
				chOut <- in
				continue
			}
			nbFound := 0

			for _, queryCode := range query {
				for _, inCode := range in.DialingCodes {
					if inCode == queryCode {
						nbFound++
					}
				}
			}

			if nbFound == queryLength {
				chOut <- in
			}
		}

		close(chOut)
	}()

	return chOut
}

// filterLanguages constructs the filter for language matching
func filterLanguages(queries []languages.Language, chIn <-chan Country) (chOut chan Country) {
	var queryLanguages []languages.Language
	for _, query := range queries {
		queryLanguages = mergeLanguages(queryLanguages, languages.Lookup(query))
	}
	queryLength := len(queryLanguages)

	chOut = make(chan Country)

	go func() {
		for in := range chIn {
			if queryLength == 0 {
				chOut <- in
				continue
			}

			nbFound := 0

			for _, query := range queryLanguages {
				for _, in := range in.Languages {
					if in.Alpha3 == query.Alpha3 {
						nbFound++
					}
				}
			}

			if nbFound == queryLength {
				chOut <- in
			}
		}

		close(chOut)
	}()

	return chOut
}

// filterName constructs the filter for name matching
func filterName(query string, chIn <-chan Country) (chOut chan Country) {
	chOut = make(chan Country)

	go func() {
		for in := range chIn {
			if query == "" {
				chOut <- in
				continue
			}

			if strings.Contains(in.Name, query) {
				chOut <- in
			}
		}

		close(chOut)
	}()

	return chOut
}
