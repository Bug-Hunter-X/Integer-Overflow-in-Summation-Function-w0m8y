func myFunc(a []int) int {
    sum := 0
    for _, v := range a {
        sum += v
    }
    return sum
}

func main() {
    a := []int{1,2,3,4,5}
    fmt.Println(myFunc(a))
}