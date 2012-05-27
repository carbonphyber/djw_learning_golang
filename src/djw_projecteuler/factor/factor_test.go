/**
 *  Copyright (c) 2012 David Wortham. All rights reserved.
 *
 * @author David Wortham <djwortham@gmail.com>
 * @author David Wortham <http://davidwortham.com/>
 * @author David Wortham <https://github.com/carbonphyber/>
 */
package factor

import (
    "fmt"
    "sort"
    "testing"
)

type factorTest struct{
    in1 uint64
    out1 map[uint64]uint64
}


var factorTests = []factorTest {
    factorTest{uint64(0),  map[uint64]uint64{}},
    factorTest{uint64(1),  map[uint64]uint64{}},
    factorTest{uint64(2),  map[uint64]uint64{uint64(2):uint64(1)}},
    factorTest{uint64(3),  map[uint64]uint64{3:1}},
    factorTest{uint64(4),  map[uint64]uint64{2:2}},
    factorTest{uint64(5),  map[uint64]uint64{5:1}},
    factorTest{uint64(6),  map[uint64]uint64{2:1,3:1}},
    factorTest{uint64(7),  map[uint64]uint64{7:1}},
    factorTest{uint64(8),  map[uint64]uint64{2:3}},
    factorTest{uint64(9),  map[uint64]uint64{3:2}},
    factorTest{uint64(10), map[uint64]uint64{2:1,5:1}},
    factorTest{uint64(11), map[uint64]uint64{11:1}},
    factorTest{uint64(12), map[uint64]uint64{2:2,3:1}},
    factorTest{uint64(13), map[uint64]uint64{13:1}},
    factorTest{uint64(14), map[uint64]uint64{2:1,7:1}},
    factorTest{uint64(15), map[uint64]uint64{3:1,5:1}},
    factorTest{uint64(16), map[uint64]uint64{2:4}},
    factorTest{uint64(17), map[uint64]uint64{17:1}},
    factorTest{uint64(18), map[uint64]uint64{2:1,3:2}},
    factorTest{uint64(19), map[uint64]uint64{19:1}},
    factorTest{uint64(64), map[uint64]uint64{2:6}},
}

/**
 * A serialization function that converts a "map[uint64]uint64" into a string that is normalized and easily comparable to other serialized "map[uint64]uint64"s
 */
func FactorsSerialize(theFactorsMap map[uint64]uint64) string {
    var returnValue string = ""
    var sortedFactors []string = make([]string, len(theFactorsMap))
    
    var i int = 0
    for key, val := range theFactorsMap {
        sortedFactors[i] = fmt.Sprintf("%d:%d", key, val)
        i++
    }
    
    // we aren't actually sorted until after this code (despite the name of the variable)
    sort.Strings(sortedFactors)
    
    i = 0
    for i = 0; i < len(sortedFactors); i++ {
        if(i == 0) {
            returnValue = fmt.Sprintf("%s%s", returnValue, sortedFactors[i])
        } else {
            returnValue = fmt.Sprintf("%s,%s", returnValue, sortedFactors[i])
        }
    }
    
    return fmt.Sprintf("{%s}", returnValue)
}

func TestFactorize(t *testing.T) {
    for _, dt := range factorTests{
        v := Factorize(dt.in1)
        serializedOut1 := FactorsSerialize(dt.out1)
        serializedV := FactorsSerialize(v)
        if serializedV != serializedOut1 {
            t.Errorf("%d failed.  Returned: %s; expected %s", dt.in1, serializedV, serializedOut1)
        }
    }
}

func BenchmarkFactorize(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
    for i := uint64(0); i < uint64(b.N); i++ {
        Factorize(i)
    }
}
