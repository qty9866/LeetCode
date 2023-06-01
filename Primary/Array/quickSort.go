package main

func quickSort(array []int, left int, right int) {
	if left < right {
		loc := partition(array, left, right)
		quickSort(array, left, loc-1)
		quickSort(array, loc+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	i := left + 1
	j := right
	for i < j {
		if arr[i] > arr[left] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else {
			i++
		}
	}
	if arr[i] >= arr[left] {
		i--
	}
	arr[i], arr[left] = arr[left], arr[i]
	return i
}

//func main() {
//	x := []int{8, 2, 6, 4, 7, 9, 1}
//	quickSort(x, 0, len(x)-1)
//
//	fmt.Println(x)
//}
