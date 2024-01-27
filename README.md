# go-pns

[![Tag](https://img.shields.io/github/tag/pulsedomains/go-pns.svg)](https://github.com/pulsedomains/go-pns/releases/)
[![License](https://img.shields.io/github/license/pulsedomains/go-pns.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/pulsedomains/go-pns?status.svg)](https://godoc.org/github.com/pulsedomains/go-pns)
[![Travis CI](https://img.shields.io/travis/pulsedomains/go-pns.svg)](https://travis-ci.org/pulsedomains/go-pns)
[![codecov.io](https://img.shields.io/codecov/c/github/pulsedomains/go-pns.svg)](https://codecov.io/github/pulsedomains/go-pns)
[![Go Report Card](https://goreportcard.com/badge/github.com/pulsedomains/go-pns)](https://goreportcard.com/report/github.com/pulsedomains/go-pns)

Go module to simplify interacting with the [PulseChain Name Service](https://pulse.domains/) contracts.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-pns` is a standard Go module which can be installed with:

```sh
go get github.com/pulsedomains/go-pns/v3
```

## Usage

`go-pns` provides simple access to the [PulseChain Name Service](https://pulse.domains/) (PNS) contracts.

### Resolution

The most commonly-used feature of PNS is resolution: converting an PNS name to an Ethereum address.  `go-pns` provides a simple call to allow this:

```go
address, err := pns.Resolve(client, domain)
```

where `client` is a connection to an Ethereum client and `domain` is the fully-qualified name you wish to resolve (e.g. `foo.mydomain.pls`) (full examples for using this are given in the [Example](#Example) section below).

The reverse process, converting an address to an PNS name, is just as simple:

```go
domain, err := pns.ReverseResolve(client, address)
```

Note that if the address does not have a reverse resolution this will return "".  If you just want a string version of an address for on-screen display then you can use `pns.Format()`, for example:

```go
fmt.Printf("The address is %s\n", pns.Format(client, address))
```

This will carry out reverse resolution of the address and print the name if present; if not it will print a formatted version of the address.


### Management of names

A top-level name is one that sits directly underneath `.pls`, for example `mydomain.pls`.  Lower-level names, such as `foo.mydomain.pls` are covered in the following section.  `go-pns` provides a simplified `Name` interface to manage top-level, removing the requirement to understand registrars, controllers, _etc._

Starting out with names in `go-pns` is easy:

```go
client, err := ethclient.Dial("https://infura.io/v3/SECRET")
name, err := pns.NewName(client, "mydomain.pls")
```

Addresses can be set and obtained using the address functions, for example to get an address:

```go
COIN_TYPE_PLS := uint64(1028)
address, err := name.Address(COIN_TYPE_PLS)
```

PNS supports addresses for multiple coin types; values of coin types can be found at https://github.com/satoshilabs/slips/blob/master/slip-0044.md

### Registering and extending names

Most operations on a domain will involve setting resolvers and resolver information.


### Management of subdomains

Because subdomains have their own registrars they do not work with the `Name` interface.

### Example

```go
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	pns "github.com/pulsedomains/go-pns/v3"
)

func main() {
	// Replace SECRET with your own access token for this example to work.
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/SECRET")
	if err != nil {
		panic(err)
	}

	// Resolve a name to an address.
	domain := "ethereum.pls"
	address, err := pns.Resolve(client, domain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address of %s is %s\n", domain, address.Hex())

	// Reverse resolve an address to a name.
	reverse, err := pns.ReverseResolve(client, address)
	if err != nil {
		panic(err)
	}
	if reverse == "" {
		fmt.Printf("%s has no reverse lookup\n", address.Hex())
	} else {
		fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
	}
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/pulsedomains/go-pns/issues).

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd
