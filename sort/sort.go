package sort
// import "fmt"

type Interface interface {

    Len () int
    Less(a, b int) bool
    Swap(a, b int)
}


func insertionSort(p_values Interface, p_begin, p_end int) {
    for i := p_begin + 1; i < p_end; i++ {
        for j := i; j > p_begin && p_values.Less(j, j - 1); j-- {
            p_values.Swap(j, j - 1)
        }
    }
}


func siftUp(p_values Interface, p_curr, p_begin, p_end int) {

    s := p_curr - p_begin
    t := (s - 1) >> 1
    for ; s > 0; {
        if p_values.Less(p_begin + t, p_begin + s) {
            p_values.Swap(p_begin + t, p_begin + s)
            s = t
            t = (t - 1) >> 1
        } else {
            break
        }
    }
}


func siftDown(p_values Interface, p_curr, p_begin, p_end int) {

    diff := p_end - p_begin
    s := p_curr - p_begin
    t := s << 1 | 1
    for ; t < diff; {
        if (p_begin + t + 1 < p_end) &&
           (p_values.Less(p_begin + t, p_begin + t + 1)) {
            t++;
        }
        if p_values.Less(p_begin + s, p_begin + t) {
            p_values.Swap(p_begin + s, p_begin + t)
            s = t
            t = t << 1 | 1
        } else {
            break
        }
    }
}


func MakeHeap(p_values Interface, p_begin, p_end int) {

    for i := ((p_end - p_begin) >> 1) - 1; i >=0; i-- {
        siftDown(p_values, p_begin + i, p_begin, p_end)
    }
}


func heapSort(p_values Interface, p_begin, p_end int) {

    MakeHeap(p_values, p_begin, p_end)

    for i := p_end - 1; i != p_begin; i-- {
        p_values.Swap(p_begin, i)
        siftDown(p_values, p_begin, p_begin, i)
    }
}

// func MedianInt(p_values []int, p_s, p_m, p_e int) int { return median(IntSlice(p_values), p_s, p_m, p_e) }

func median(p_values Interface, p_s, p_m, p_e int) int {
    if p_values.Less(p_s, p_m) {
        p_values.Swap(p_s, p_m)
        //  fmt.Println(p_s, p_m)
    }
    if p_values.Less(p_e, p_s) {
        if p_values.Less(p_m, p_e) {
            p_values.Swap(p_s, p_e)
        } else {
            p_values.Swap(p_s, p_m)
        }
    }
    return p_s
}


func partition(p_values Interface, p_begin, p_end int) int {

    i := p_begin
    j := p_end - 1
    idx := median(p_values, p_begin, p_begin + ((p_end - p_begin) >> 1), p_end - 1)

    for ; i < j; {
    // fmt.Println(i, j, p_begin, p_end)

        for ; idx < j; {
            if !p_values.Less(j, idx) {
                j--
            } else {
                break;
            }
        }

        if idx < j {
            p_values.Swap(idx, j)
            idx = j
        }

        for ; i < idx; {
            if !p_values.Less(idx, i) {
                i++
            } else {
                break
            }
        }

        if i < idx {
            p_values.Swap(i, idx)
            idx = i
        }
    }
    // fmt.Println(p_begin, p_end, idx)
    return idx
}


func partition3way(p_values Interface, p_begin, p_end int) (int, int) {

    i := p_begin
    ii := p_begin
    j := p_end - 1
    jj := p_end - 1
    idx := median(p_values, p_begin, p_begin + ((p_end - p_begin) >> 1), p_end - 1)

    for ; i < j; {

        for ; idx < j; {
            // fmt.Println("j->", p_values, "i: ", i, "j: ", j, "idx:", idx)
            if p_values.Less(idx, j) {
                j--
            } else if !p_values.Less(idx, j) && !p_values.Less(j, idx) {
                p_values.Swap(jj, j)
                jj--
                j--
            } else {
                p_values.Swap(idx, j)
                idx = j
            }
        }

        for ; i < idx; {
            // fmt.Println("i->", p_values, "i: ", i, "j: ", j, "idx:", idx)
            if p_values.Less(i, idx) {
                i++
            } else if !p_values.Less(i, idx) && !p_values.Less(idx, i) {
                p_values.Swap(ii, i)
                ii++
                i++
            } else {
                p_values.Swap(i, idx)
                idx = i
            }
        }
    }

    for i = p_begin; i < ii; i++ {
        p_values.Swap(i, p_begin + idx - 1 - i)
    }

    for j = p_end - 1; j > jj; j-- {
        p_values.Swap(j, idx + p_end - j)
    }

    return p_begin + idx - ii, idx + p_end - jj
}


//func Part(p_values []int, p_begin, p_end int) int  { return partition(IntSlice(p_values), p_begin, p_end) }

func introSort(p_values Interface, p_begin, p_end, p_depth int) {

    if p_end - p_begin > 12 {

        if p_depth == 0 {
             heapSort(p_values, p_begin, p_end)
             return ;
        }

        p_depth--

        // fmt.Println("debug", p_begin, p_end, p_values)
        mlo, mhi := partition3way(p_values, p_begin, p_end)
        /* for i := 0; i != mlo; i++ {
            if !p_values.Less(i, mlo) {
                fmt.Println(p_values, "p_beign: ", p_begin, "p_end: ", p_end, "mlo: ", mlo, "mhi: ",mhi)
            }
        }
        for i := mhi; i != p_values.Len(); i++ {
            if !p_values.Less(mlo, i) {
                fmt.Println(p_values, "p_beign: ", p_begin, "p_end: ", p_end, "mlo: ", mlo, "mhi: ",mhi)
            }
        }*/
        introSort(p_values, p_begin, mlo, p_depth)
        introSort(p_values, mhi, p_end, p_depth)
    } else {
        insertionSort(p_values, p_begin, p_end)
    }
}


func QuickSort_(p_values []int) {
    QuickSort(IntSlice(p_values), 0, len(p_values))
}


func QuickSort(p_values Interface, p_begin, p_end int) {

    if p_end - p_begin > 1 {
        mid := partition(p_values, p_begin, p_end)
        QuickSort(p_values, p_begin, mid)
        QuickSort(p_values, mid + 1, p_end)
    }
}


func QuickSort3Way_(p_values []int) {
    QuickSort3Way(IntSlice(p_values), 0, len(p_values))
}


func QuickSort3Way(p_values Interface, p_begin, p_end int) {

    if p_end - p_begin > 1 {
        mlo, mhi := partition3way(p_values, p_begin, p_end)
        QuickSort3Way(p_values, p_begin, mlo)
        QuickSort3Way(p_values, mhi, p_end)
    }
}

func maxDepth(p_cnt int) int {
    depth := 0
    for ; p_cnt != 0; p_cnt >>= 1 {
        depth++
    }
    return depth << 1
}


func Sort(p_values Interface) {
    depth := maxDepth(p_values.Len())
    introSort(p_values, 0,  p_values.Len(), depth)
    // insertionSort(p_values, 0, p_values.Len())
    // heapSort(p_values, 0, p_values.Len())
}


func AreSorted(p_values Interface) bool {
    length := p_values.Len() - 1
    for i := 0; i < length; i++ {
        if p_values.Less(i + 1, i) {
            return false;
        }
    }
    return true;
}

type UIntSlice []uint
func (p_arr UIntSlice) Len() int           { return len(p_arr) }
func (p_arr UIntSlice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr UIntSlice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type UInt64Slice []uint64
func (p_arr UInt64Slice) Len() int           { return len(p_arr) }
func (p_arr UInt64Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr UInt64Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type UInt32Slice []uint32
func (p_arr UInt32Slice) Len() int           { return len(p_arr) }
func (p_arr UInt32Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr UInt32Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type UInt16Slice []uint16
func (p_arr UInt16Slice) Len() int           { return len(p_arr) }
func (p_arr UInt16Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr UInt16Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type UInt8Slice []uint8
func (p_arr UInt8Slice) Len() int           { return len(p_arr) }
func (p_arr UInt8Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr UInt8Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type IntSlice []int
func (p_arr IntSlice) Len() int           { return len(p_arr) }
func (p_arr IntSlice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr IntSlice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type Int64Slice []int64
func (p_arr Int64Slice) Len() int           { return len(p_arr) }
func (p_arr Int64Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr Int64Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type Int32Slice []int32
func (p_arr Int32Slice) Len() int           { return len(p_arr) }
func (p_arr Int32Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr Int32Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type Int16Slice []int16
func (p_arr Int16Slice) Len() int           { return len(p_arr) }
func (p_arr Int16Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr Int16Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type Int8Slice []int8
func (p_arr Int8Slice) Len() int           { return len(p_arr) }
func (p_arr Int8Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr Int8Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type Float64Slice []float64
func (p_arr Float64Slice) Len() int           { return len(p_arr) }
func (p_arr Float64Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr Float64Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }

type Float32Slice []float32
func (p_arr Float32Slice) Len() int           { return len(p_arr) }
func (p_arr Float32Slice) Less(a, b int) bool { return p_arr[a] < p_arr[b] }
func (p_arr Float32Slice) Swap(a, b int)      { p_arr[a], p_arr[b] = p_arr[b], p_arr[a] }


func SortInt  (p_arr []int)       { Sort(IntSlice(p_arr)) }
func SortInt64(p_arr []int64)     { Sort(Int64Slice(p_arr)) }
func SortInt32(p_arr []int32)     { Sort(Int32Slice(p_arr)) }
func SortInt16(p_arr []int16)     { Sort(Int16Slice(p_arr)) }
func SortInt8 (p_arr []int8)      { Sort(Int8Slice(p_arr)) }
func SortUInt  (p_arr []uint)     { Sort(UIntSlice(p_arr)) }
func SortUInt64(p_arr []uint64)   { Sort(UInt64Slice(p_arr)) }
func SortUInt32(p_arr []uint32)   { Sort(UInt32Slice(p_arr)) }
func SortUInt16(p_arr []uint16)   { Sort(UInt16Slice(p_arr)) }
func SortUInt8 (p_arr []uint8 )   { Sort(UInt8Slice(p_arr)) }
func SortFloat64(p_arr []float64) { Sort(Float64Slice(p_arr)) }
func SortFloat32(p_arr []float32) { Sort(Float32Slice(p_arr)) }

func AreSortedInt    (p_arr []int)     bool { return AreSorted(IntSlice(p_arr)) }
func AreSortedInt64  (p_arr []int64)   bool { return AreSorted(Int64Slice(p_arr)) }
func AreSortedInt32  (p_arr []int32)   bool { return AreSorted(Int32Slice(p_arr)) }
func AreSortedInt16  (p_arr []int16)   bool { return AreSorted(Int16Slice(p_arr)) }
func AreSortedInt8   (p_arr []int8)    bool { return AreSorted(Int8Slice(p_arr)) }
func AreSortedUInt   (p_arr []uint)    bool { return AreSorted(UIntSlice(p_arr)) }
func AreSortedUInt64 (p_arr []uint64)  bool { return AreSorted(UInt64Slice(p_arr)) }
func AreSortedUInt32 (p_arr []uint32)  bool { return AreSorted(UInt32Slice(p_arr)) }
func AreSortedUInt16 (p_arr []uint16)  bool { return AreSorted(UInt16Slice(p_arr)) }
func AreSortedUInt8  (p_arr []uint8 )  bool { return AreSorted(UInt8Slice(p_arr)) }
func AreSortedFloat64(p_arr []float64) bool { return AreSorted(Float64Slice(p_arr)) }
func AreSortedFloat32(p_arr []float32) bool { return AreSorted(Float32Slice(p_arr)) }
