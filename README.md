# coinflip-go

CLI program that simulates a coin flip.

If you ever have a hard time making a decision, don't
fret no more! Let the universe decide for you.

[![Go](https://github.com/MawKKe/coinflip-go/actions/workflows/go.yml/badge.svg)](https://github.com/MawKKe/coinflip-go/actions/workflows/go.yml)

## Install

    $ go install github.com/MawKKe/coinflip-go@latest

## Usage

View help:

    $ coinflip-go help

To flip a coin and get either "heads" or "tails", run:

    $ coinflip-go coinflip
    heads

...or if you want result as either "yes" or "no":

    $ coinflip-go yesno
    no

...or if you want to pick among choices (e.g number 1-5):

    $ coinflip-go choose 1 2 3 4 5
    4


Note: if you installed the binary with `go install`, make sure `$GOPATH/bin`
is included in your shell's `$PATH`

## License

Copyright 2022 Markus Holmström (MawKKe)

The works under this repository are licenced under Apache License 2.0.
See file `LICENSE` for more information.

## Contributing

This project is hosted at https://github.com/MawKKe/coinflip-go

You are welcome to leave bug reports, fixes and feature requests. Thanks!

