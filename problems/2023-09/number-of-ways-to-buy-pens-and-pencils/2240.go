package number_of_ways_to_buy_pens_and_pencils

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	if cost1 < cost2 {
		return waysToBuyPensPencils(total, cost2, cost1)
	}
	var res, cnt int
	for cnt*cost1 <= total {
		res += (total-cnt*cost1)/cost2 + 1
		cnt++
	}
	return int64(res)
}
