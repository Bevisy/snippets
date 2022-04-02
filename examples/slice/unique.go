package main

import "fmt"

/*
  利用泛型支持多类型 slice 去重
  参考：https://gosamples.dev/generics-remove-duplicates-slice/
*/
func unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

type FruitRank struct {
	Fruit string
	Rank  uint64
}

func main() {
	fmt.Println(unique([]string{"abc", "cde", "efg", "efg", "abc", "cde"}))
	fmt.Println(unique([]int{1, 1, 2, 2, 3, 3, 4}))

	fruits := []FruitRank{
		{
			Fruit: "Strawberry",
			Rank:  1,
		},
		{
			Fruit: "Raspberry",
			Rank:  2,
		},
		{
			Fruit: "Blueberry",
			Rank:  3,
		},
		{
			Fruit: "Blueberry",
			Rank:  3,
		},
		{
			Fruit: "Strawberry",
			Rank:  1,
		},
	}
	fmt.Println(unique(fruits))
}
