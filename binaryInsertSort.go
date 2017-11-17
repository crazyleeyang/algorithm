func binaryInsertSort(a []int){
    l := len(a)
    for i := 1; i < l; i++{
        tmp := a[i]
        low, high := 0, i-1
        for ; low<=high;{
            mid := (low+high)/2
            if tmp < a[mid]{
                high = mid-1
            } else {
                low = mid+1
            }
        }

        for j:=i;j>low;j--{
            a[j] = a[j-1]
        }

        a[low] = tmp
    }
}
