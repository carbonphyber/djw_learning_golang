/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package palindrome

import (
    "testing"
)

type palindromeTest struct{
    in1 uint64
    out1 bool
}



var palindromeTests = []palindromeTest {
    palindromeTest{uint64(0), true},
    palindromeTest{uint64(1), true},
    palindromeTest{uint64(01), true},
    palindromeTest{uint64(10), false},
    palindromeTest{uint64(11), true},
    palindromeTest{uint64(12), false},
    palindromeTest{uint64(121), true},
    palindromeTest{uint64(2332), true},
    palindromeTest{uint64(3413), false},
    palindromeTest{uint64(4567654), true},
}

func TestPalindrome(t *testing.T) {
    for _, dt := range palindromeTests{
        v := IsPalindrome(dt.in1)
        if v != dt.out1 {
            t.Errorf("%d failed.  Returned: %b; expected %b", dt.in1, v, dt.out1)
        }
    }
}

func BenchmarkIsPalindrome(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := uint64(0); i < uint64(b.N); i++ {
        IsPalindrome(i)
    }
}
