func reverseBits(n int) int {
    res :=0
    for i:=range 32 {
        if ((n >> i) & 1 == 1 ){
            res |= 1 << (31-i)
        }
    }
    return res
}
