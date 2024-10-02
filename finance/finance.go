package finance

func GetEarningsMetrics(revenue, expenses, taxRate float64) (float64, float64, float64, bool) {
	ebt := revenue - expenses
	tax := ebt * (taxRate / 100.0)
	profit := ebt - tax

	var ratio float64
	warning := false

	if profit == 0 {
		warning = true
		ratio = 0
	} else {
		ratio = ebt / profit

		if profit < 0 {
			ratio = -ratio
		}
	}

	return ebt, profit, ratio, warning
}
