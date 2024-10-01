package finance

func GetEarningsMetrics(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	tax := revenue * (taxRate / 100.0)
	profit := ebt - tax
	ratio := ebt / profit

	return ebt, profit, ratio
}
