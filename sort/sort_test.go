package sort_test

import(
    stdsort "sort"
    "github.com/BismarckDD/golang/sort"
    "fmt"
    "testing"
    "time"
    "math/rand"
)

const (
    MAXN = 1000000
)

var arr []int = make([]int, MAXN, MAXN)
var arr1 []int = make([]int, MAXN, MAXN)
var arr2 []int = make([]int, MAXN, MAXN)
var arr3 []int = make([]int, MAXN, MAXN)
var st time.Time
var elapse time.Duration

func TestRandomData(t *testing.T) {

    for i := 0; i < MAXN; i++ {
        arr[i] = rand.Intn(MAXN)
        arr1[i] = arr[i]
        arr2[i] = arr[i]
        arr3[i] = arr[i]
    }

    st = time.Now()
    stdsort.Ints(arr)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr))
    fmt.Println("Standard Library time cost: ", elapse)

    st = time.Now()
    sort.SortInt(arr1)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr1))
    fmt.Println("Intro time cost: ", elapse)

    st = time.Now()
    sort.QuickSort_(arr2)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr2))
    fmt.Println("Normal QuickSort time cost: ", elapse)

    st = time.Now()
    sort.QuickSort3Way_(arr3)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr3))
    fmt.Println("3Way QuickSort time cost: ", elapse)
}

func TestSoretedData(t *testing.T) {

    for i := 0; i < MAXN; i++ {
        arr[i] = i
        arr1[i] = i
        arr2[i] = i
        arr3[i] = i
    }

    st = time.Now()
    stdsort.Ints(arr)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr))
    fmt.Println("Standard Library time cost: ", elapse)

    st = time.Now()
    sort.SortInt(arr1)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr1))
    fmt.Println("Intro time cost: ", elapse)

    st = time.Now()
    sort.QuickSort_(arr2)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr2))
    fmt.Println("Normal QuickSort time cost: ", elapse)

    st = time.Now()
    sort.QuickSort3Way_(arr3)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr3))
    fmt.Println("3Way QuickSort time cost: ", elapse)

}

func TestEqualData(t *testing.T) {

    for i := 0; i < MAXN; i++ {
        arr[i] = rand.Intn(MAXN) / 100
        arr1[i] = arr[i]
        arr2[i] = arr[i]
        arr3[i] = arr[i]
    }

    st = time.Now()
    stdsort.Ints(arr)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr))
    fmt.Println("Standard Library time cost: ", elapse)

    st = time.Now()
    sort.SortInt(arr1)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr1))
    fmt.Println("Intro time cost: ", elapse)

    st = time.Now()
    sort.QuickSort_(arr2)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr2))
    fmt.Println("Normal QuickSort time cost: ", elapse)

    st = time.Now()
    sort.QuickSort3Way_(arr3)
    elapse = time.Since(st)
    fmt.Println("Is Successful? ", sort.AreSortedInt(arr3))
    fmt.Println("3Way QuickSort time cost: ", elapse)

}

func TestMedian(t *testing.T) {


}







// a,b := sort.Part3Way(arr, 0, len(arr))
    // fmt.Println(arr)
    //fmt.Println(a, b)

    /* t := 10
    arr = []int {1, 2, 3}
    t = sort.MedianInt(arr, 0, 1, 2)
    fmt.Println(arr,":", t)
    arr = []int {1, 3, 2}
    t = sort.MedianInt(arr, 0, 1, 2)
    fmt.Println(arr,":", t)
    arr = []int {2, 1, 3}
    t = sort.MedianInt(arr, 0, 1, 2)
    fmt.Println(arr,":", t)
    arr = []int {2, 3, 1}
    t = sort.MedianInt(arr, 0, 1, 2)
    fmt.Println(arr,":", t)
    arr = []int {3, 1, 2}
    t = sort.MedianInt(arr, 0, 1, 2)
    fmt.Println(arr,":", t)
    arr = []int {3, 2, 1}
    t = sort.MedianInt(arr, 0, 1, 2)
    fmt.Println(arr,":", t) */

