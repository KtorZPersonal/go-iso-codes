// Copyright © 2016 Matthias Benkort
// Use of this source code is governed by the MPL 2.0 license that can be found in the LICENSE file.

package languages

import "strings"

// Language materializes a language object, with its name and ISO identifiers
type Language struct {
	// Alpha2 refers to a 2-letter ISO 639-1 Code, may be blank for some languages
	Alpha2 string
	// Alpha2 refers to a 3-letter ISO 639-2 Code
	Alpha3 string
	// Name refers to the language human-readable english name
	Name string
}

// Languages exports all languages available in the package
var Languages = reduceLanguages(languages)

func reduceLanguages(dict map[string]Language) (xs []Language) {
	for _, language := range dict {
		xs = append(xs, language)
	}

	return xs
}

// Lookup retrieves a list of languages satisfying the criteria(s) given in arguments
func Lookup(query Language) []Language {
	if query.Alpha3 != "" {
		res, ok := languages[query.Alpha3]

		if ok {
			return []Language{res}
		}

		return make([]Language, 0)
	}

	if query.Alpha2 != "" {
		for _, res := range languages {
			if query.Alpha2 == res.Alpha2 {
				return []Language{res}
			}
		}
	}

	results := make([]Language, 0)
	if query.Name != "" {
		for _, res := range languages {
			if strings.Contains(res.Name, query.Name) {
				results = append(results, res)
			}
		}
	}
	return results
}
