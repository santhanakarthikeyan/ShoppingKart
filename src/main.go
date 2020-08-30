package main

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "github.com/graphql-go/graphql"
)

type Product struct {
	SKU string
	Name string
	Price float64
	Quantity int
}

func populate() []Product {
	product1 := Product{
		SKU: "120P90",
		Name: "Google Home",
		Price: 49.99,
		Quantity: 10,
	}
	product2 := Product{
		SKU: "43N23P",
		Name: "MacBook Pro",
		Price: 5399.99,
		Quantity: 5,
	}
	product3 := Product{
		SKU: "A304SD",
		Name: "Alexa Speaker",
		Price: 109.50,
		Quantity: 10,
	}
	product4 := Product{
		SKU: "234234",
		Name: "Raspberry Pi B",
		Price: 30.00,
		Quantity: 2,
	}
	var products []Product
	products = append(products, product1)
	products = append(products, product2)
	products = append(products, product3)
	products = append(products, product4)
	return products
}

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"Sku": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				 Type: graphql.String,
			 },
			 "Price": &graphql.Field{
				 Type: graphql.Float,
			 },
			 "Quantity": &graphql.Field{
				 Type: graphql.Int,
			 },
		},
	},
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func addToCart(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Added product to cart")
	// Implement bussiness logic
}

func checkout(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Scanned items:\n Total:\n")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/add_to_cart", addToCart)
    http.HandleFunc("/checkout", checkout)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	products := populate()

	fields := graphql.Fields{
		"products": &graphql.Field{
			Type: productType,
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sku, ok := p.Args["sku"].(string)
				if ok {
					for _, product := range products {
						if string(product.SKU) == sku {
							return product, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(productType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return products, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := `
	{
		products(sku: "120P90") {
			Name
		}
	}`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	query = `
	{
		list {
			Name
		}
	}`

	params = graphql.Params{Schema: schema, RequestString: query}
	r = graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ = json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

	//handleRequests()
}
