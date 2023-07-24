# nill

[![Github Actions](https://github.com/RekeningkuDev/nill/actions/workflows/ci.yml/badge.svg)](https://github.com/RekeningkuDev/nill/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/RekeningkuDev/nill?status.svg)](https://godoc.org/github.com/RekeningkuDev/nill)
[![Go Report Card](https://goreportcard.com/badge/github.com/RekeningkuDev/nill)](https://goreportcard.com/report/github.com/RekeningkuDev/nill)

## Features
* Support for echo (web framework) https://echo.labstack.com/
* The zero-value is 0, and is safe to use without initialization
* Addition, subtraction, multiplication, division, modulo and etc.
* JSON serialization/deserialization
* Support for Custom JSON Encoder/Decoder

## Install

Run `go get github.com/RekeningkuDev/nill`

## Requirements

nill library requires Go version `>=1.18`

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/RekeningkuDev/nill"
	goJson "github.com/goccy/go-json"
)

func init() {
	// Custom JSON unmarshal or Unmarshal
	nill.JSONMarshal = goJson.Marshal
	nill.JSONUnmarshal = goJson.Unmarshal
}

func main() {
	price := nill.NewFloat64(50.99)
	// Add 
	subtotal := price.Mul(nill.NewFloat64(8))
	fmt.Println("Subtotal:", subtotal)    // Subtotal: 407.92
	
	// Custom type
	type product struct {
		Name nill.String `json:"name"`
		Price nill.Float64 `json:"price"`
    }
	
	type information struct {
		Product nill.Type[product] `json:"product"`
		Qty nill.Int `json:"qty"`
    }
	
	// JSON serialization/deserialization
	p := information{
		Product: nill.Type[product]{
			Valid: true,
			Value: product{
				Name: nill.NewString("Product A"),
				Price: nill.NewFloat64(50.99),
			},
		},
		Qty: nill.NewInt(8),
	}
	
	// Marshal
	b, err := json.Marshal(p)
	if err != nil {
        fmt.Println(err)
	}
	
	fmt.Println(string(b)) // {"product":{"name":"Product A","price":50.99},"qty":8}
}
```

## Documentation

http://godoc.org/github.com/RekeningkuDev/nill

## License
The MIT License (MIT)