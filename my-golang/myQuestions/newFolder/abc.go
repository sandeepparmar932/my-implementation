package newFolder

func A(arr []int, i int, j int) {
	arrLength := len(arr)
	if arrLength == 0 || j > arrLength || i > arrLength || i <= 0 || j <= 0 {
		return
	}

	arr[i], arr[j] = arr[j], arr[i]
}
