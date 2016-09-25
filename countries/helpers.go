// Copyright © 2016 Matthias Benkort
// Use of this source code is governed by the MPL 2.0 license that can be found in the LICENSE file.

package countries

import (
	"github.com/KtorZ/go-iso-codes/currencies"
	"github.com/KtorZ/go-iso-codes/languages"
)

// reduceCountries constructs a list of Country from a given map of country
func reduceCountries(dict map[string]country) (xs []Country) {
	for _, country := range dict {
		xs = append(xs, toCountry(country))
	}

	return xs
}

// toCountry transforms a country into a Country
func toCountry(country country) Country {
	var countryCurrencies []currencies.Currency
	for _, code := range country.Currencies {
		countryCurrencies = mergeCurrencies(countryCurrencies, currencies.Lookup(currencies.Currency{Alpha3: code}))
	}

	var countryLanguages []languages.Language
	for _, code := range country.Languages {
		countryLanguages = mergeLanguages(countryLanguages, languages.Lookup(languages.Language{Alpha3: code}))
	}

	return Country{
		Alpha2:       country.Alpha2,
		Alpha3:       country.Alpha3,
		Numeric:      country.Numeric,
		Currencies:   countryCurrencies,
		DialingCodes: country.DialingCodes,
		Languages:    countryLanguages,
		Name:         country.Name,
	}
}

// mergeCurrencies computes the union of two lists of currencies, without duplicates
func mergeCurrencies(dest []currencies.Currency, elems []currencies.Currency) []currencies.Currency {
	var merge []currencies.Currency

uniqueOnly:
	for _, a := range elems {
		for _, b := range dest {
			if a.Alpha3 == b.Alpha3 {
				continue uniqueOnly
			}
		}

		merge = append(dest, a)
	}

	return merge
}

// mergeLanguages computes the union of two lists of languages, without duplicates
func mergeLanguages(dest []languages.Language, elems []languages.Language) []languages.Language {
	var merge []languages.Language

uniqueOnly:
	for _, a := range elems {
		for _, b := range dest {
			if a.Alpha3 == b.Alpha3 {
				continue uniqueOnly
			}
		}

		merge = append(dest, a)
	}

	return merge
}
