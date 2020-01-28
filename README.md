# Base58

[![Build Status](https://travis-ci.org/wellmart/base58.svg?branch=master)](https://travis-ci.org/wellmart/base58)
[![Go Report Card](https://goreportcard.com/badge/github.com/wellmart/base58)](https://goreportcard.com/report/github.com/wellmart/base58)
[![Coverage Status](https://coveralls.io/repos/github/wellmart/base58/badge.svg?branch=master)](https://coveralls.io/github/wellmart/base58?branch=master)
![Version](https://img.shields.io/badge/version-0.1.0-blue)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wellmart/base58?status.svg)](https://godoc.org/github.com/wellmart/base58)

Represent large integers as alphanumeric text, it is similar to Base 64 but has been modified to avoid both non-alphanumeric characters and letters which might look ambiguous when printed.

## Requirements

Go 1.1 and beyond.

## Installation

Use the go package manager to install Base 58.

```bash
go get github.com/wellmart/base58
```

## Usage

```go
package main

import "github.com/wellmart/base58"

func main() {
    value := base58.EncodeToString(byte[]("Hello World"))
    bytes := base58.DecodeString(value)
}
```

## Staying up to date

To update Base 58 to the latest version, use `go get -u github.com/wellmart/base58`.

## License

[MIT](https://choosealicense.com/licenses/mit/)
