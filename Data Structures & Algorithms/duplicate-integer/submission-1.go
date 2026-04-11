func hasDuplicate(nums []int) bool {
    m := make(map[int]int) 

    for _, v := range nums {
        if _, ok := m[v]; ok {
            fmt.Println(v, m[v])
            return true
        }

        m[v] = 1
    }

    return false
}
