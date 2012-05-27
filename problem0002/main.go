/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package main

import (
    "flag"
    "fmt"
)


/**
 * https://projecteuler.net/problem=2
 */
func main(){
    var max_natural_number uint64
    var sum uint64
    var i uint64
    var prev, curr, next uint64
    var debug bool
    
    flag.BoolVar(&debug, "debug", false, "Display verbose info");
    flag.Uint64Var(&max_natural_number, "max", uint64(0), "Maximum natural number");
    // now that all flags have been assigned, we can parse the CLI parameters
    flag.Parse()
    
    // make sure the CLI input is acceptable; if not error out with a helpful usage message
    if (max_natural_number < uint64(1)) {
        fmt.Println("You must pass in a a maximum (Unsigned 64 integer)");
        return
    } else {
        if debug { fmt.Printf("maximum term: %d\n", max_natural_number) }
    }
    
    prev = uint64(1)
    if debug { fmt.Printf("1: %d\n", prev) }
    curr = uint64(2)
    if debug { fmt.Printf("2: %d\n", curr) }
    sum = prev + curr
    // we have hard-coded terms 1 and 2.  Start the programmatic arithmatic with term 3
    for i = 3; curr <= max_natural_number; i++ {
        next = prev + curr
        if next > max_natural_number {
            break;
        }
        if next % 2 == 0 {
            if debug { fmt.Printf("%d: %d (+)\n", i, next) }
            sum += next
        } else {
            if debug { fmt.Printf("%d: %d\n", i, next) }
        }
        prev = curr
        curr = next
    }
    
    // spit out the answer.
    fmt.Printf("%d\n", sum)
    
    return
}
