func longestConsecutive(nums []int) int {
    res := 0
    hash := make(map[int]bool)

    for _, v := range nums {
        hash[v] = true
    }

    for _, v := range nums {
        longest := 0
        if hash[v-1] != false {
            continue
        }

        number := v
        for hash[number] {
            longest+=1
            number = number + 1
        }

        if longest > res {
            res = longest
        }
    }

    return res
}
