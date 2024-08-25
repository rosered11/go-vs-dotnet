package services

type Heavy struct{}

func NewHeavy() *Heavy {
	return &Heavy{}
}

func (factorial *Heavy) CPUBurn(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= i
	}
	return result
}

func (factorial *Heavy) MemoryBurn(n int) int {
	largeSlice := make([]uint32, n)
	for i := 0; i < n; i++ {
		largeSlice[i] = uint32(i)
	}
	var sum uint32 = 0
	for _, num := range largeSlice {
		sum += num
	}
	return int(sum)
}
