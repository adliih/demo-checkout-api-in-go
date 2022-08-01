package graph

import (
	"adliih/demo-checkout-api-in-go/checkout"
	"adliih/demo-checkout-api-in-go/graph/model"
	"adliih/demo-checkout-api-in-go/sku"
	"fmt"

	"github.com/google/uuid"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	ProductStore map[string]*model.Product
	TransactionStore map[string]*model.Transaction
}

func (r *Resolver) init() {
	if (r.ProductStore == nil) {
		r.ProductStore = make(map[string]*model.Product)
	}
	if (r.TransactionStore == nil) {
		r.TransactionStore = make(map[string]*model.Transaction)
	}
}

func (r *Resolver) getProductBySku(sku string) (*model.Product)  {
	r.init()
	return r.ProductStore[sku]
}

func (r *Resolver) saveProduct(product *model.Product) {
	r.init()
	r.ProductStore[product.Sku] = product
}

func (r *Resolver) saveTransaction(transaction *model.Transaction) (*model.Transaction)  {
	r.init()
	transaction.ID = uuid.NewString()
	r.TransactionStore[transaction.ID] = transaction
	return transaction
}

func (r *Resolver) getDiscountValue(input model.CheckoutInput) (discount float64) {
	discount = 0.0
	if isEligible, qty := checkBundlingGoogleHomeRaspberryPiPromo(input); isEligible {
		// free raspberry pi
		discount += r.getProductBySku(sku.RaspberryPi).Price * float64(qty)
	}
	
	if isEligible, qty := checkGoogleHomeDiscountPromo(input); isEligible {
		discount += r.getProductBySku(sku.GoogleHome).Price * float64(qty / 3)
	}
	
	if isEligible, qty := checkAlexaSpeakerDiscountPromo(input); isEligible {
		discount += r.getProductBySku(sku.AlexaSpeaker).Price * float64(qty) * 0.1
	}
	return
}

func checkBundlingGoogleHomeRaspberryPiPromo(input model.CheckoutInput) (bool, int) {
	bundleCount := checkout.CountBundlingItemRequest(input, sku.GoogleHome, sku.RaspberryPi)
	return bundleCount > 0, bundleCount
}

func checkGoogleHomeDiscountPromo(input model.CheckoutInput) (bool, int) {
	discountedItemRequest := checkout.CountItemRequest(input, sku.GoogleHome)
	fmt.Println("Google Home", discountedItemRequest)
	return discountedItemRequest >= 3, discountedItemRequest
}

func checkAlexaSpeakerDiscountPromo(input model.CheckoutInput) (bool, int) {
	discountedItemRequest := checkout.CountItemRequest(input, sku.AlexaSpeaker)
	fmt.Println("Alexa", discountedItemRequest)
	return discountedItemRequest >= 3, discountedItemRequest
}
