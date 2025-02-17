func myFunc(a []int) int64 {
    sum := int64(0)
    for _, v := range a {
        sum += int64(v)
    }
    return sum
}

func main() {
    a := []int{1,2,3,4,5}
    fmt.Println(myFunc(a))
}    
