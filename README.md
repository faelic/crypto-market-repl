# Crypto Market Watch REPL

A terminal-based REPL built in Go for checking live crypto market data.

This project was built as a personal backend learning project. The goal was not just to build a working CLI tool, but to practice structuring a real Go application that handles command parsing, API communication, JSON decoding, validation, error handling, and testing.

## Features

- `/help` to view available commands
- `/list` to view a market overview of supported coins
- `/price <coin>` to get the live USD price of a supported coin
- `/market <coin>` to get richer market details for a supported coin
- `/exit` to quit the REPL

## Supported Coins

For v1, the app supports:

- `bitcoin`
- `ethereum`
- `solana`

## Sample Usage

```text
> /help
Available commands:
    /help
    /list
    /price <coin>
    /market <coin>
    /exit

    Supported coins:
    bitcoin
    ethereum
    solana

