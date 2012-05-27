/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package main

import (
    "../src/djw_projecteuler/palindrome"
    "flag"
    "fmt"
)

var (
    debug bool
)

/**
 * https://projecteuler.net/problem=3
 */
func main(){
    var maximum_natural_number uint64
    var minimum_natural_number uint64
    var num_digits uint64
    var largest_found uint64
    var product uint64
    
    flag.BoolVar(&debug, "debug", false, "Display verbose info");
    flag.Uint64Var(&num_digits, "numdigits", uint64(0), "The number of digits in the two multipliers");
    // now that all flags have been assigned, we can parse the CLI parameters
    flag.Parse()
    
    // make sure the CLI input is acceptable; if not error out with a helpful usage message
    if (num_digits < uint64(1)) {
        fmt.Println("You must pass in a number that is the number of digits in the two multipliers (Unsigned 64bit integer)");
        return
    } else {
        if debug { fmt.Printf("Finding palindrome products of two %d-digit numbers\n", num_digits) }
    }
    
    switch num_digits {
    case uint64(1):
        maximum_natural_number = 9
        minimum_natural_number = 1
        break;
    case uint64(2):
        maximum_natural_number = 99
        minimum_natural_number = 10
        break;
    case uint64(3):
        maximum_natural_number = 999
        minimum_natural_number = 100
        break;
    case uint64(4):
        maximum_natural_number = 9999
        minimum_natural_number = 1000
        break;
    case uint64(5):
        maximum_natural_number = 99999
        minimum_natural_number = 10000
        break;
    case uint64(6):
        maximum_natural_number = 999999
        minimum_natural_number = 100000
        break;
    case uint64(7):
        maximum_natural_number = 9999999
        minimum_natural_number = 1000000
        break;
    case uint64(8):
        maximum_natural_number = 99999999
        minimum_natural_number = 10000000
        break;
    case uint64(9):
        maximum_natural_number = 999999999
        minimum_natural_number = 100000000
        break;
    default:
        fmt.Println("The number of digits you entered was too large.  This edge case is not yet taken care of.");
        return;
        break;
    }
    
    largest_found = uint64(1)
    // testValues := map[uint64]bool {
    //     0: true,
    //     6: true,
    //     11: true,
    //     121: true,
    //     2321: false,
    //     3414: false,
    //     4554: true,
    // }
    
    // difficult logic goes here
    for i := minimum_natural_number; i < maximum_natural_number; i++ {
        for j := minimum_natural_number; j < maximum_natural_number; j++ {
            product = i * j
            if product > largest_found && palindrome.IsPalindrome(product) {
                if debug { fmt.Printf("%d * %d = %d (*)\n", i, j, product) }
                largest_found = product
            }
        }
    }
    
    // spit out the answer.
    fmt.Printf("%d\n", largest_found)
    
    return
}
