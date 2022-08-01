package checkout

import (
	"adliih/demo-checkout-api-in-go/graph/model"
)

func CountBundlingItemRequest(input model.CheckoutInput, sku... string) int  {
	var counter = make(map[string]int)
	var result = 0
	for _, v := range input.Products {
		existing, isExist := counter[v.Sku]
		if !isExist {
			counter[v.Sku] = 0
			existing = counter[v.Sku]
		}
		existing += v.Qty

		if result == 0 {
			result = counter[v.Sku]
		}

		if counter[v.Sku] < result {
			result = counter[v.Sku]
		}
	}
	// check is all sku found in the request
	if (len(counter) != len(sku)) {
		return 0
	}
	return result
}

func CountItemRequest(input model.CheckoutInput, sku string) int  {
	counter := 0
	for _, v := range input.Products {
		if v.Sku != sku {
			continue
		}
		counter += v.Qty
	}
	return counter
}


