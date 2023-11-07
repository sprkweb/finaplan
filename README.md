<p align="center">
    <img src="https://raw.githubusercontent.com/sprkweb/finaplan-cli/master/icon.svg" alt="Logo" width="150" height="150" />
</p>

<h1 align="center">FinaPlan CLI</h1>

Financial planning using modules.

Inspired by [Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy):

- Each module does exactly one thing and does it well;
- Modules are composed together and work together in a way that is defined by user;
- Modules handle simple streams of numbers.

## Installation

    git clone https://github.com/sprkweb/finaplan-cli.git
    cd finaplan-cli
    make install

## Available modules

    Usage:
    finaplan [command]

    Available Commands:
    init        Initialize an empty plan
    add         Add a certain amount of money to your plan
    invest      Add some interest rate on top of your capital
    print       Print the given plan in a more human-friendly manner

## Usage

Let's say you want to make a financial plan for 3 years,
and you want to know how much money you will have in the end.

    $ finaplan init --intervals 36 --months | \ # plan for 3 years = 36 months
        finaplan add 1000 --each 1 | \          # save 1000$ each month
        finaplan invest 10% --interval 12 | \   # invest your savings at 10% per year = per 12 months
        finaplan inflation 4% --interval 12 | \ # adjust for inflation: 4% per year = per 12 monthw
        finaplan print | tail
    Month 26 | 29994.15
    Month 27 | 31233.33
    Month 28 | 32482.39
    Month 29 | 33741.41
    Month 30 | 35010.46
    Month 31 | 36289.64
    Month 32 | 37579.02
    Month 33 | 38878.68
    Month 34 | 40188.71
    Month 35 | 41509.18

## Development

### Run tests

    make test

### Build

    make build
