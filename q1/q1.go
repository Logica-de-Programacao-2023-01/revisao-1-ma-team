package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	val_desc := 0.0
	soma := 0.0
	for _, valor := range purchaseHistory {
		soma += valor
	}
	if soma/float64(len(purchaseHistory)) > 1000 {
		val_desc = float64(currentPurchase) * 0.20
	}
	if len(purchaseHistory) == 0 {
		val_desc = float64(currentPurchase) * 0.10
	}
	if soma > 1000 && soma/float64(len(purchaseHistory)) <= 1000 {
		val_desc = float64(currentPurchase) * 0.10
	}
	if soma > 500 && soma <= 1000 {
		val_desc = float64(currentPurchase) * 0.05
	}
	if soma > 0 && soma <= 500 {
		val_desc = float64(currentPurchase) * 0.02
	}
	if currentPurchase > 0 {
		return val_desc, nil
	}
	return 0, fmt.Errorf("valor da compra invalido")
}
