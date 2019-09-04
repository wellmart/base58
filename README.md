# Base58

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

## License
[MIT](https://choosealicense.com/licenses/mit/)
