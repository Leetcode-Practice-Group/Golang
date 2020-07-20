func twoSum(numbers []int, target int) []int {
    len := len(numbers)
    for i:=0;i<len-1;i++{
        tmp := target - numbers[i]
        for j:=i+1;j<len;j++{
            if tmp == numbers[j]{
                return []int{i+1, j+1}
            }
        }
    }
    return []int{0,0}
}
