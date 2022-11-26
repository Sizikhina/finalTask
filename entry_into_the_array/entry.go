package entry_into_the_array

func Solution(A []int) int {
	//инициализируем map
	result := make(map[int]bool)
	//в цикле проходимся по слайсу и инвертируем значения в map
	for _, value := range A {
		result[value] = !result[value]
	}
	//так как у нас по условию только одно число без пары, то у него будет значение true
	for key, value := range result {
		if value {
			return key
		}
	}
	return 0
}
