<p align="center">
    <img src="https://raw.githubusercontent.com/sprkweb/finaplan-cli/master/icon.svg" alt="Logo" width="150" height="150" />
</p>

<h1 align="center">FinaPlan</h1>

Financial planning using modules.

Inspired by [Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy):

- Each module does exactly one thing and does it well;
- Modules are composed together and work together in a way that is defined by user;
- Modules handle simple streams of numbers.

## Installation

    go install github.com/sprkweb/finaplan@v0.0.3

## Available modules

    Usage:
    finaplan [command]

    Available Commands:
    init        Initialize an empty plan
    add         Add a certain amount of money to your plan
    invest      Add some interest rate on top of your capital
    inflation   Adjust all the previous plan for inflation
    print       Print the given plan in a more human-friendly manner

## Usage

Let's say you want to make a financial plan for 3 years,
and you want to know how much money you will have in the end.

    $ finaplan init --intervals 36 --months | \ # plan for 3 years = 36 months
        finaplan add 1000 --each 1 | \          # save 1000$ each month
        finaplan invest 10% --interval 12 | \   # invest your savings at 10% per year = per 12 months
        finaplan inflation 4% --interval 12 | \ # adjust for inflation: 4% per year = per 12 monthw
        finaplan print | tail
    Month 26 | 27550.60
    Month 27 | 28595.21
    Month 28 | 29641.72
    Month 29 | 30690.17
    Month 30 | 31740.56
    Month 31 | 32792.91
    Month 32 | 33847.24
    Month 33 | 34903.57
    Month 34 | 35961.92
    Month 35 | 37022.31

## Development

### Run tests

    make test

### Build

    make build
