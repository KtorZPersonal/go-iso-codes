// Copyright © 2016 Matthias Benkort
// Use of this source code is governed by the MPL 2.0 license that can be found in the LICENSE file.

package languages

import "fmt"

// sort ascendently sort a slice of languages on the Alpha3
func sort(xs []Language) []Language {
	var _xs []Language
sort:
	for _, a := range xs {
		for i, b := range _xs {
			if a.Alpha3 < b.Alpha3 {
				_xs = append(_xs[:i], append([]Language{a}, _xs[i:]...)...)
				continue sort
			}
		}

		_xs = append(_xs, a)
	}

	return _xs
}

// ExampleLookup illustrates a lookup made based on the Alpha3 field
func ExampleLookup() {
	results := Lookup(Language{Alpha3: "fra"})
	fmt.Println(results)

	results = Lookup(Language{Name: "French"})
	fmt.Println(sort(results))

	// Output:
	// [{fr fra French}]
	// [{ cpf Creoles and pidgins, French-based} {fr fra French} { frm French, Middle (ca.1400-1600)} { fro French, Old (842-ca.1400)}]
}
