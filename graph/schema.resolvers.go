package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adliih/demo-checkout-api-in-go/graph/generated"
	"adliih/demo-checkout-api-in-go/graph/model"
	"context"
	"fmt"
)

// Checkout is the resolver for the checkout field.
func (r *mutationResolver) Checkout(ctx context.Context, input model.CheckoutInput) (*model.Transaction, error) {
	transaction := &model.Transaction{}
	details := make(map[string]*model.TransactionDetail)

	for _, productInput := range input.Products {
		// fetch the product
		product := r.getProductBySku(productInput.Sku)

		if product == nil {
			return nil, fmt.Errorf("Invalid requested product: %s", productInput.Sku)
		}

		// quantity validation
		if  productInput.Qty > product.Qty {
			// panic(fmt.)
			return nil, fmt.Errorf("Invalid requested product quantity: %s", productInput.Sku)
		}

		detail := &model.TransactionDetail{}
		detail.Product = product
		detail.Price = product.Price
		detail.Qty = 0

		if _, isExist := details[product.Sku]; !isExist {
			details[product.Sku] = detail
		}

		details[product.Sku].Qty += productInput.Qty
	}

	for _, detail := range details {	
		transaction.SubTotal += detail.Price * float64(detail.Qty)
		transaction.Details = append(transaction.Details, detail)
	}

	transaction.Discount = r.getDiscountValue(input)
	transaction.Total = transaction.SubTotal - transaction.Discount
	r.saveTransaction(transaction)

	return transaction, nil
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input []*model.ProductInput) ([]*model.Product, error) {
	var products []*model.Product
	for _, v := range input {
		product := &model.Product{
			Sku: v.Sku,
			Name: v.Name,
			Price: v.Price,
			Qty: v.Qty,
		}
		r.saveProduct(product)
		products = append(products, product)
	}

	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
