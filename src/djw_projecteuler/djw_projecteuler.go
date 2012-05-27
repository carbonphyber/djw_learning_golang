/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package djw_projecteuler

var knownPrimeStatus map[uint64]bool
var maxKnownPrimeStatus uint64

/**
 * IsPrime
 * @assert input is a valid uint64
 */
func IsPrime(input uint64) (bool) {
    var i uint64
    var iIsPrime bool
    
    if input < uint64(2) {
        return true
    }
    
    if nil == knownPrimeStatus {
        knownPrimeStatus = make(map[uint64]bool)
        maxKnownPrimeStatus = 0
    }
    if maxKnownPrimeStatus > input {
        return knownPrimeStatus[input]
    }
    
    // easy hack to save some CPUs on large inputs.  Because 2 is the lowest prime, input can not be evenly divisible by anything above/including: (input / 2) + 1
    halfOfInput := input / uint64(2)
    if halfOfInput + 1 > input {
        halfOfInput = input
    }
    
    // if input is evenly divisivble by any number between 1 and input (exclusive), it is not prime
    for i = 2; i < halfOfInput; i++ {
        iIsPrime = input % i == uint64(0)
        // write to our cache if we haven't already recorded a value
        if(i > maxKnownPrimeStatus) {
            knownPrimeStatus[i] = iIsPrime
            maxKnownPrimeStatus = i
        }
        if(iIsPrime) {
            return false
        }
    }
    
    // prime if all numbers have already been tested
    return true
}

