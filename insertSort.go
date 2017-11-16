func insertSort(a []int){
    l := len(a)
    for i:=1;i<l;i++{
        tmp := a[i]
        j := i-1
        for ;j>=0;j--{
            if a[j] > tmp{
                a[j+1] = a[j]
            } else {
                break
            }
        }

        if j+1 != i{
            a[j+1] = tmp
        }
    }
}
