package builtin

import (
    "fmt"
    "testing"
)

func TestMap(t *testing.T) {

    // Two methods to initialize map.
    // map1 := map[string]string {}
    // map2 := make(map[string]string)

    map1 := map[string]string { 
        "chensi" : "nvshen",
        "xunaihao" : "diaosi",
        "songying" : "diaosi",
        "dingdong" : "fushuai",
    }

    // Traverse the map.
    for key, value := range map1 {
        fmt.Println(key, value)
    }
    // Get the length of map.
    fmt.Println(len(map1))

    key := "zhanglong"
    map2 := make(map[string]string)
    map2[key] = "hehe"

    value, ok := map2[key]
    if ok {
        fmt.Println(key, ": ", value)
    } else {
        fmt.Println("Not exist")
    }
    // Delete the element of the map.
    delete(map2, key)

    value, ok = map2[key]
    if ok {
        fmt.Println(key, ": ", value)
    } else {
        fmt.Println("Not exist")
    }
}

func TestSlice(t *testing.T) {

    // Two methods to initialize slice.
    slice1 := []int {}
    slice2 := make([]int, 10) // must set the length.

    slice1 = append(slice1, 1)
    slice1 = append(slice1, 2)
    slice1 = append(slice1, 3)
    slice1 = append(slice1, 4)
    slice1 = append(slice1, 5)

    fmt.Println(slice1)

    i := 2
    copy(slice1[i:], slice1[i + 1:])
    slice1 = slice1[:len(slice1) - 1]
    fmt.Println(slice1)

    i = 3
    slice2[9] = 10;

}

func TestArray(t *testing.T) {
    
    array1 := [10]int {}
    array2 := make([10]int)

    array1[0] = 0
    array1[1] = 1
    array1[2] = 2
    array1[3] = 3
    array1[4] = 4
    fmt.Println(array1)

    array2[0] = 100
    fmt.Println(array2)
}
