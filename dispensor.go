package gumi

type Dispensor func(have int, lns []Length) []int
func DispensorOrder(have int, lns []Length) []int {
	var res = make([]int, len(lns))
	//
	var minsum, maxsum int
	for _, v := range lns{
		minsum += int(v.Min)
		maxsum += int(v.Max)
	}
	if maxsum <= have{
		for i, v := range lns{
			res[i] = int(v.Max)
		}
	}else if minsum <= have{
		left := have - minsum
		for i, v := range lns{
			if left > int(v.Max) - int(v.Min){
				res[i] = int(v.Max)
				left -= int(v.Max) - int(v.Min)
			}else if left > 0{
				res[i] = int(v.Min) + left
				left = 0
			}else {
				res[i] = int(v.Min)
			}
		}
	}else {
		left := have
		for i, v := range lns{
			if left > int(v.Min){
				res[i] = int(v.Min)
				left -= int(v.Min)
			}else if left > 0{
				res[i] = left
				left = 0
			}else {
				res[i] = 0
			}
		}
	}
	return res
}
func DispensorMinimalize(have int, lns []Length) []int {
	var res = make([]int, len(lns))
	//
	for i, v := range lns{
		if have - int(v.Min) > 0 {
			res[i] = int(v.Min)
			have -= int(v.Min)
		}else if have > 0 {
			res[i] = have
			have = 0
		}else {
			res[i] = 0
		}

	}
	return res
}