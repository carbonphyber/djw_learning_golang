/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package main

import (
    "../src/djw_projecteuler"
    "flag"
    "fmt"
)

var (
    debug bool
)

/**
 * @assert input is a valid uint64
 */
func LargestPrimeFactor(input uint64) (uint64) {
    var remainder uint64
    var i uint64
    var largestFactor uint64
    var largestPrimeFactor uint64
    
    remainder = input
    largestPrimeFactor = 1
    for i = uint64(2); remainder > 1 && i < input; i++ {
        for ;remainder > 1 && remainder % i == uint64(0); {
            remainder /= i
            if debug { fmt.Printf("%d is a factor, %d remainder\n", i, remainder) }
            largestFactor = i
            if debug { fmt.Printf("%d", i) }
            if djw_projecteuler.IsPrime(largestFactor) {
                largestPrimeFactor = largestFactor
                if debug { fmt.Printf(" (is prime)") }
            }
            if debug { fmt.Printf("\n") }
        }
    }
    return largestPrimeFactor
}

/**
 * https://projecteuler.net/problem=3
 */
func main(){
    var natural_number uint64
    var largestPrimeFactor uint64
    
    flag.BoolVar(&debug, "debug", false, "Display verbose info");
    flag.Uint64Var(&natural_number, "number", uint64(0), "A number to find factors for");
    // now that all flags have been assigned, we can parse the CLI parameters
    flag.Parse()
    
    // make sure the CLI input is acceptable; if not error out with a helpful usage message
    if (natural_number < uint64(1)) {
        fmt.Println("You must pass in a a number to factorize (Unsigned 64 integer)");
        return
    } else {
        if debug { fmt.Printf("factorizing this term: %d\n", natural_number) }
    }
    
    largestPrimeFactor = LargestPrimeFactor(natural_number)
    
    // spit out the answer.
    fmt.Printf("%d\n", largestPrimeFactor)
    
    return
}
