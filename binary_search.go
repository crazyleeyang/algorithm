func binary_search(a []int, b int)(ok bool, idx int){
    l := len(a)
    if l == 0{ return }

    for low, mid, high := 0, 0, l-1;low <= high;{
        mid = (low+high)/2
        if b == a[mid]{
            return true, mid
        } else if b > a[mid] {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return
}
