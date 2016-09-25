# CODES [![](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/KtorZ/codes) [![](https://img.shields.io/badge/license-MPL%202.0-blue.svg?style=flat-square)](https://www.mozilla.org/en-US/MPL/2.0/) [![](https://travis-ci.org/KtorZ/codes.svg?style=flat-square)](https://travis-ci.org/KtorZ/codes)

## Overview

codes aggregates data to get various ISO code and details about countries, currencies and
languages among the world.

The package offers methods to quickly lookup a country, language or currency based on a code, a
fraction of the name or a combination of other criterias.

The full lists of data are also available and can be easily generated to JSON via the
`encoding/json` package.

## Getting Started

Three packages are available should you need to only import `currencies` or `languages`. Note
nevertheless that the `countries` package uses the two others.

```go
import (
    "fmt"

    "github.com/KtorZ/go-iso-codes/countries"
    "github.com/KtorZ/go-iso-codes/currencies"
    "github.com/KtorZ/go-iso-codes/languages"
)

func main() {
    eurCurrency := currencies.Lookup(currencies.Currency{Alpha3: "EUR"})
    fmt.Println(eurCurrency)
    // [{EUR 978 2 Euro}]

    allCurrencies := currencies.Currencies
    fmt.Println(allCurrencies)
    // A lot of stuff

    europeanEnglishCountries := countries.Lookup(countries.Country{
        Currencies: []currencies.Currency{eurCurrency},
        Languages: []languages.Language{
            Alpha2: "EN",
        },
    })
    fmt.Println(europeanEnglishCountries)
    // Also a lot of stuff
}
```

## Release Notes

- **0.1.0** Sept 25, 2016
    - First release
