func insertSort(a []int){
    l := len(a)
    for i:=1;i<l;i++{
        tmp := a[i]
        for j:=i-1;j>=0;j--{
            if a[j] > tmp{
                a[j+1] = a[j]
                a[j] = tmp
            } else {
                break
            }
        }
    }
}
