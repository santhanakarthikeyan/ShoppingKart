# ShoppingKart

An implementation of ShoppingKart using Golang with GraphQL. Follow the below steps for installation. 

### Getting started
```bash
git clone https://github.com/santhanakarthikeyan/ShoppingKart
go get github.com/graphql-go/graphql
go install src/main
./bin/main
```

## Documentation
### Add to cart to API
Use this API to add a product to cart 
```
http://URL:PORT/add_to_cart?sku=<SKU>
```

### Checkout API
This API list down cart items and total amount to pay
```
http://URL:PORT/checkout
```
