// Copyright 2022 Markus Holmström (MawKKe) markus@mawkke.fi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// coinflip: CLI program that simulates a coin flip
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

// genRandomBit generates random integer with value 0 or 1.
func genRandomUpto(lim int) (int, error) {
	// Our fate shall be decided by true randomness,
	// not some silly pseudo-random number generator.
	i, err := rand.Int(rand.Reader, big.NewInt(int64(lim)))
	if err != nil {
		return 0, err
	}
	return int(i.Int64()), nil
}

func choose(choices []string, gen func(int) (int, error)) (string, error) {
	if len(choices) < 2 {
		return "", fmt.Errorf("expected at least 2 arguments to choose(), got %d", len(choices))
	}
	i, err := gen(len(choices))
	if err != nil {
		return "", fmt.Errorf("problem generating random number: %w", err)
	}
	return choices[i], nil
}

func pickRandomly(choices []string) (string, error) {
	return choose(choices, genRandomUpto)
}

func printUsage(ec int) {
	fmt.Printf("Usage: %v coinflip|yesno|choose\n", os.Args[0])
	os.Exit(ec)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("error: not enough arguments\n\n")
		printUsage(1)
	}
	cmd, args := os.Args[1], os.Args[2:]

	if cmd == "help" {
		printUsage(0)
	}

	res, err := func(cmd string, extraArgs []string) (string, error) {
		switch cmd {
		case "coinflip":
			return pickRandomly([]string{"tails", "heads"})
		case "yesno":
			return pickRandomly([]string{"no", "yes"})
		case "choose":
			return pickRandomly(extraArgs)
		default:
			return "", fmt.Errorf("unknown subcommand '%v'", cmd)
		}
	}(cmd, args)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(res)
}
