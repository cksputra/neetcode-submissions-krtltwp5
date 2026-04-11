func add(a ,b int) int {
    return a+b
}

func mult(a, b int) int {
    return a * b
}

func subs(a,b int) int {
    return a - b
}

func div(a,b int) int {
    return a / b
}

func evalRPN(tokens []string) int {

    numStack := []int{}

    opMap := map[string]func(int,int) int {
        "-" : subs,
        "+" : add,
        "*" : mult,
        "/" : div,
    }

    for _, v := range tokens {
        token := string(v)

        if _, ok := opMap[token]; ok{
            b := numStack[len(numStack)-1]
            numStack = numStack[:len(numStack)-1]
            a := numStack[len(numStack)-1]
            numStack = numStack[:len(numStack)-1]

            res := opMap[token](a, b)

            numStack = append(numStack, res)
        } else {
            num, _ := strconv.Atoi(token)
            numStack = append(numStack, num)
        }
    }

    return numStack[0]
}


