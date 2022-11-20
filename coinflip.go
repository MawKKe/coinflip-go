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
	"flag"
	"fmt"
	"math/big"
)

// genRandomBit generates random integer with value 0 or 1.
func genRandomBit() int {
	// Our fate shall be decided by true randomness,
	// not some silly pseudo-random number generator.
	i, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		panic(err)
	}
	return int(i.Int64() % 2)
}

// bitToString returns string representation for 'bit'.
// the argument 'bit' should be either 0 or 1. If it is not, then
// the modulo of the value is used.
func bitToString(bit int, yesno bool) string {
	if yesno {
		return []string{"no", "yes"}[bit%2]
	} else {
		return []string{"tails", "heads"}[bit%2]
	}
}

func main() {
	var yesno bool
	flag.BoolVar(&yesno, "yesno", false, "Result is yes|no instead of heads|tails")

	flag.Parse()

	bit := genRandomBit()

	answer := bitToString(bit, yesno)

	fmt.Println(answer)
}
