func simpleSelectionSort(a []int){
    l := len(a)

    for i:=0;i<l;i++{
        minLoc := minIdx(a[i:])
        if minLoc == -1{
            return
        } else if minLoc != 0{
            temp := a[i]
            a[i] = a[minLoc+i]
            a[minLoc+i] = temp
        }
    }
}

func minIdx(a []int)(minIdx int){
    l := len(a)
    if l == 0{
        return -1
    }

    min := a[0]
    for i:=0;i<l;i++{
        if a[i] < min{
            min = a[i]
            minIdx = i
        }
    }
    return
}
