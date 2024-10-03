package finance

func GetEarningsMetrics(revenue, expenses, taxRate float64) (
	ebt float64,
	profit float64,
	ratio float64,
	warning bool) {

	ebt = revenue - expenses
	tax := ebt * (taxRate / 100.0)
	profit = ebt - tax

	warning = false

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
