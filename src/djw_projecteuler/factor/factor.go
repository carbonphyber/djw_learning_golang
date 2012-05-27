/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package factor

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
        }
    }
    return factors
}
