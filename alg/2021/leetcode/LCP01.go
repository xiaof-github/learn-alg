func game(guess []int, answer []int) int {
    var count int
    for i:=0;i<len(guess);i++ {
        if guess[i] == answer[i] {
            count++
        }
    }
    return count
}