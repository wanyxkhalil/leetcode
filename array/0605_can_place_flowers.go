package array

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	for i := 0; i < length && n>0; i+=2 {
		if flowerbed[i] ==0 {
			if i+1==length || flowerbed[i+1]==0 {
				n--
			} else {
				i++
			}
		}
	}
	if n==0 {
			return true 
	}
	return false
}

