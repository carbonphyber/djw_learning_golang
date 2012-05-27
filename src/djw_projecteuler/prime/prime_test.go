/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package prime

import (
    "testing"
)

type primeTest struct{
    in1 uint64
    out1 bool
}



var primeTests = []primeTest {
    primeTest{uint64(0), true},
    primeTest{uint64(1), true},
    primeTest{uint64(2), true},
    primeTest{uint64(3), true},
    primeTest{uint64(4), false},
    primeTest{uint64(5), true},
    primeTest{uint64(6), false},
    primeTest{uint64(7), true},
    primeTest{uint64(8), false},
    primeTest{uint64(9), false},
    primeTest{uint64(10), false},
    primeTest{uint64(11), true},
    primeTest{uint64(12), false},
    primeTest{uint64(13), true},
    primeTest{uint64(14), false},
    primeTest{uint64(15), false},
    primeTest{uint64(16), false},
    primeTest{uint64(17), true},
    primeTest{uint64(18), false},
    primeTest{uint64(19), true},
}

func TestPrime(t *testing.T) {
    for _, dt := range primeTests{
        v := IsPrime(dt.in1)
        if v != dt.out1 {
            t.Errorf("%d failed.  Returned: %b; expected %b", dt.in1, v, dt.out1)
        }
    }
}

func BenchmarkIsPrime(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := uint64(0); i < uint64(b.N); i++ {
        IsPrime(i)
    }
}
