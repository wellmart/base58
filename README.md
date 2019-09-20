# Base58

[![Build Status](https://travis-ci.org/wellmart/base58.svg?branch=master)](https://travis-ci.org/wellmart/base58)
[![Go Report Card](https://goreportcard.com/badge/github.com/wellmart/base58)](https://goreportcard.com/report/github.com/wellmart/base58)

Represent large integers as alphanumeric text, it is similar to Base64 but has been modified to avoid both non-alphanumeric characters and letters which might look ambiguous when printed.

## Installation

Use the go package manager to install Base58.

```bash
go get github.com/wellmart/base58
```

## Usage

```go
import "github.com/wellmart/base58"

value := base58.EncodeToString(byte[]("Hello World"))
bytes := base58.DecodeString(value)
```

## Staying up to date

To update Base58 to the latest version, use `go get -u github.com/wellmart/base58`.

## License
[MIT](https://choosealicense.com/licenses/mit/)
