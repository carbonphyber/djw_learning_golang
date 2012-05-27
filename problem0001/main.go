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
 * https://projecteuler.net/problem=1
 */
func main(){
    var max_natural_number *uint64
    var sum uint64
    var i uint64
    
    max_natural_number = flag.Uint64("max", uint64(0), "Maximum natural number");
    // now that all flags have been assigned, we can parse the CLI parameters
    flag.Parse()
    
    // make sure the CLI input is acceptable; if not error out with a helpful usage message
    if (max_natural_number < unit64(1)) {
        fmt.Println("You must pass in a a maximum (Unsigned 64 integer)");
        return
    } else {
        fmt.Printf("max_natural_number: %d\n", *max_natural_number)
    }
    
    sum = 0
    for i = 1; i < *max_natural_number; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            fmt.Printf("multiple of 3 or 5: %d\n", i)
            sum += i
        }
    }
    
    // spit out the answer.
    fmt.Printf("%d\n", sum)
    
    return
}
