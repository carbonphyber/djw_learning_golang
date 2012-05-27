/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package main

import (
    // "../src/djw_projecteuler/prime"
    "flag"
    "fmt"
)

var (
    debug bool
)

/**
 * @assert input is a valid uint64
 */
func Factorize(input uint64) (map[uint64]uint64) {
    var remainder uint64
    var i uint64
    var factors map[uint64]uint64
    
    factors = make(map[uint64]uint64)
    remainder = input
    for i = uint64(2); remainder > 1 && i <= input; i++ {
        for ;remainder > 1 && remainder % i == uint64(0); {
            remainder /= i
            factors[i] += 1
            if debug { fmt.Printf("%d is a factor, %d remainder\n", i, remainder) }
        }
    }
    return factors
}

/**
 * https://projecteuler.net/problem=3
 */
func main(){
    var natural_number uint64
    var lowest_common_demoninator uint64
    var allFactors map[uint64]uint64
    var theseFactors map[uint64]uint64
    
    flag.BoolVar(&debug, "debug", false, "Display verbose info");
    flag.Uint64Var(&natural_number, "number", uint64(0), "A number to find factors for");
    // now that all flags have been assigned, we can parse the CLI parameters
    flag.Parse()
    
    // make sure the CLI input is acceptable; if not error out with a helpful usage message
    if (natural_number < uint64(1)) {
        fmt.Println("You must pass in a number to factorize (Unsigned 64 integer)");
        return
    } else {
        if debug { fmt.Printf("factorizing terms: 1 .. %d\n", natural_number) }
    }
    
    allFactors = make(map[uint64]uint64)
    lowest_common_demoninator = uint64(1)
    for i := uint64(2); i < natural_number; i++ {
        if debug { fmt.Printf("factorizing this term: %d\n", i) }
        theseFactors = Factorize(i)
        for key, val := range theseFactors {
            if debug { fmt.Printf("%d => %d\n", key, val) }
            if theseFactors[key] > allFactors[key] {
                allFactors[key] = theseFactors[key]
            }
        }
    }
    
    for key, val := range allFactors {
        if debug { fmt.Printf("%d => %d\n", key, val) }
        for j := 0; uint64(j) < val; j++ {
            lowest_common_demoninator *= key
        }
    }
    
    // spit out the answer.
    fmt.Printf("%d\n", lowest_common_demoninator)
    
    return
}
