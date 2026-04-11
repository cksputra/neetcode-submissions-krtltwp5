type MinStack struct {
    Min int
    Stack []int
}

func Constructor() MinStack {
    init := MinStack{
        Min: math.MaxInt,
    }
    return init
}

func (this *MinStack) Push(val int) {
    this.Stack = append(this.Stack, val)
    if val < this.Min {
        this.Min = val
    }
}

func (this *MinStack) Pop() {
    this.Stack = this.Stack[:len(this.Stack)-1]


    if len(this.Stack) > 0 {
        this.Min = this.Stack[0]
        for _,v := range this.Stack {
            if v < this.Min {
                this.Min = v
            }
        }
    } else {
        this.Min = math.MaxInt
    }
}

func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.Min
}
