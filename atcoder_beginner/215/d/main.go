// WIP
package main

import (
	"fmt"
	"math"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	list, _ := get_prime(M)
	fmt.Println(list)
	a := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		list = removeInt(list, a)
	}
	fmt.Println(list[0] + 1)
	fmt.Println(1)
	for i := 1; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func removeInt(s_list []int, i int) (tmp []int) {
	var a intSlice
	a = s_list
	pos := a.pos(i)
	if pos == -1 {
		return s_list
	} else {
		return remove(s_list, pos)

	}

}

type intSlice []int

func (slice intSlice) pos(value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

//配列の特定の要素を削除する関数
func remove(s_list []int, index int) (tmp []int) {
	tmp = append(s_list[:index], s_list[(index+1):]...)
	return
}

func get_prime(number int) ([]int, int) {
	//初期化
	prime_list := []int{}
	search_list := []int{}
	//2からnumberまでの数字の配列を作る
	for i := 2; i < number+1; i++ {
		search_list = append(search_list, i)
	}
	//探索リストの先頭値が√numberを超えたら終了
	limit := int(math.Sqrt(float64(number)))
	for search_list[0] <= limit {
		//探索リストの先頭を素数リストに移動
		p_num := search_list[0]
		prime_list = append(prime_list, p_num)
		//探索リストの先頭を削除
		search_list = remove(search_list, 0)
		//p_numの倍数を探索リストから篩落とす
		search_list_length := len(search_list)
		tmp := []int{}
		for i := 0; i < search_list_length; i++ {
			if search_list[i]%p_num != 0 {
				tmp = append(tmp, search_list[i])
			}
		}
		search_list = tmp
	}
	//探索リストの残りを素数リストに追加
	prime_list = append(prime_list, search_list...)
	return prime_list, len(prime_list)
}
