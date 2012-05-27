/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package palindrome

import (
    "math"
    "strconv"
)

/**
 * IsPalindrome
 * @assert input is a valid uint64
 */
func IsPalindrome(input uint64) (bool) {
    // var i uint64
    var iIsPalindrome bool
    var stringInput string
    var halfInputLength int
    
    if input < uint64(10) {
        return true
    }
    
    iIsPalindrome = true
    stringInput = strconv.FormatUint(input, 10)
    halfInputLength = int( math.Floor( float64(len(stringInput)-1) ) )
    
    // read this carefully: we are assigning the values of i and j in the same assignment (both the first and third statements ofr the "for")
    for i, j := 0, halfInputLength; i < j; i, j = i+1, j-1 {
        if stringInput[i] != stringInput[j] {
            iIsPalindrome = false
            break
        }
    }
    
    // prime if all numbers have already been tested
    return iIsPalindrome
}

