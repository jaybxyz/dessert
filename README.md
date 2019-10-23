<h1 align="center">🙇 Welcome to Dessert 🙇</h1>

<p>
<img  src="https://img.shields.io/badge/version-0.0.1-blue.svg?cacheSeconds=2592000" />
</p>

**[WARNING]: THIS PROJECT IS CURRENTLY A WORK IN PROGRESS.**

## Overview

This project is a very simple project that I am using as a toy project to hack with [Comsos SDK](https://github.com/cosmos/cosmos-sdk). This enables you to query your account balance, your accumulated rewards, and to create different types of transactions and broadcast them to either gaia testnet or mainnet. 

I called this project `dessert` because it is like dessert that gives you an appetite of how to work with Cosmos SDK. I hope that this will help any developer out there to grasp an understanding of how to hack with Cosmos SDK.

Give ⭐️ if anyone who finds this useful, and feel free to add comments or requests in [issues](https://github.com/kogisin/dessert/issues). I will be working from time to time. 

Enjoy :smile:

## Table of Contents

- [Requirements](#Requirements)
- [Installation](#Installation)
- [Usage](#Usage)
- [Examples](#Examples)
- [Progress](#Progress)
- [Contributing](#Contributing)
- [Support](#Support)⭐️ ️
- [License](#License)

## Requirements

This project assumes that you already know [Go language](https://golang.org/) and know how to use [Go Modules](https://github.com/golang/go/wiki/Modules)

- Require [Go 1.13+](https://golang.org/dl/)
- Prepare endpoints for your own full node and REST server. If you don't have one, ask [Riot](riot.cosmos.network) channel for any available public full node to connect to.

## Installation

Clone this repository using `git`.
```
$ git clone https://github.com/kogisin/dessert.git
```

Dessert has a simple way to create a configuration file.
(Other option: copy `example.toml` file and paste it in the same directory and change the name as `config.toml`
```
$ dessert init
```

Install an executable file `dessert`
```
$ go build

OR

$ make install 
```

## Usage

Dessert has a few subcommands that you can hack around. By putting your mnemonic in `config.toml`, `dessert` has full control over your wallet. I recommend you to use test wallet.

- Init
- Version
- Query 
- Tx 
- Wallet

## Examples

Check your account balance
```
$ dessert q balance
```

## Progress

- Add more queries
- Add unit tests
- Implement REST Server that behaves as Comsos SDK Light Client. 

## Contributing

The best way to submit feedback and report bugs is to [Open a GitHub issue](https://github.com/kogisin/dessert/issues). 

If you want to contribute code, please see [CONTRIBUTING.md](https://github.com/kogisin/dessert/contrib) and read the requirement sections carefully.

## Support ⭐️ 

**Give it a ⭐️star if you find this project useful** 

## License

[Apache License Version 2.0](https://github.com/kogisin/dessert/LICENSE)