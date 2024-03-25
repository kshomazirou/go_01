package piscine

func Len(table []int) int {
	count := 0
	for range table {
		count++
	}
	return count
}

func SortIntegerTable(table []int) {
	n := Len(table)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if table[j] > table[j + 1] {
				tmp := table[j]
				table[j] = table[j + 1]
				table[j + 1] = tmp
			}
		}
	}

}
