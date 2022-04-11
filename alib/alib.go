package alib

func Average(s []int) int {
	//スライスの合計値
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
