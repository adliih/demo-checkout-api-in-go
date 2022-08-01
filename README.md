# demo-checkout-api-in-go
Simple Checkout Graphql API that has item promotion calculation


# Run the server 
```sh
go run server.go
```
Then Open the Graphql Playground (http://localhost:8080/)

# Graphql API(s)
## Init the data
```graphql
mutation createProduct{
  createProduct(input: [
     {
        sku: "120P90"
        name: "Google Home"
        price: 49.99
        qty: 10
    },
    {
        sku: "43N23P"
        name: "Macbook Pro"
        price: 5399.99
        qty: 5
    },
    {
        sku: "A304SD"
        name: "Alexa Speaker"
        price: 109.50
        qty: 10
    },
    {
        sku: "234234"
        name: "Raspberry Pi"
        price: 30.00
        qty: 2
    }
  ]) {
    sku
    name
    price
    qty
  }
}
```


## Checkout
```graphql
mutation checkout {
  checkout(input: {
    products: [
      {
        sku: "120P90"
        qty: 2
      },
      {
        sku: "43N23P"
        qty: 1
      }
    ]
  }) {
    id
    subTotal
    total
    discount
    details {
      qty
      price
      product {
        sku
        name
      }
    }
  }
}
```
