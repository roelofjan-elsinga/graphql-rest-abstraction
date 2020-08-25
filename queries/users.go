package queries

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/roelofjan-elsinga/graphql-rest-abstraction/models"
)

var companyField = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Company",
	Description: "A company",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the company",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if company, ok := p.Source.(models.Company); ok {
					return company.Name, nil
				}
				return nil, nil
			},
		},
		"catch_phrase": &graphql.Field{
			Type:        graphql.String,
			Description: "The catch phrase of the company",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if company, ok := p.Source.(models.Company); ok {
					return company.CatchPhrase, nil
				}
				return nil, nil
			},
		},
		"bs": &graphql.Field{
			Type:        graphql.String,
			Description: "The bs of the company",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if company, ok := p.Source.(models.Company); ok {
					return company.BS, nil
				}
				return nil, nil
			},
		},
	},
})

var geoField = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Geolocation",
	Description: "This is an geolocation object",
	Fields: graphql.Fields{
		"latitude": &graphql.Field{
			Type:        graphql.String,
			Description: "The latitude of the geolocation",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if geo, ok := p.Source.(models.Geo); ok {
					return geo.Latitude, nil
				}
				return nil, nil
			},
		},
		"longitude": &graphql.Field{
			Type:        graphql.String,
			Description: "The longitude of the geolocation",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if geo, ok := p.Source.(models.Geo); ok {
					return geo.Longitude, nil
				}
				return nil, nil
			},
		},
	},
})

var addressField = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Address",
	Description: "This is an address object",
	Fields: graphql.Fields{
		"street": &graphql.Field{
			Type:        graphql.Int,
			Description: "The street of the address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if address, ok := p.Source.(models.Address); ok {
					return address.Street, nil
				}
				return nil, nil
			},
		},
		"suite": &graphql.Field{
			Type:        graphql.String,
			Description: "The suite of the address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if address, ok := p.Source.(models.Address); ok {
					return address.Suite, nil
				}
				return nil, nil
			},
		},
		"city": &graphql.Field{
			Type:        graphql.String,
			Description: "The city of the address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if address, ok := p.Source.(models.Address); ok {
					return address.City, nil
				}
				return nil, nil
			},
		},
		"zipcode": &graphql.Field{
			Type:        graphql.String,
			Description: "The zipcode of the address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if address, ok := p.Source.(models.Address); ok {
					return address.Zipcode, nil
				}
				return nil, nil
			},
		},
		"geo": &graphql.Field{
			Type:        geoField,
			Description: "The geolocation of the address",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if address, ok := p.Source.(models.Address); ok {
					return address.Geo, nil
				}
				return nil, nil
			},
		},
	},
})

var userObject = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "The user to return from the REST API",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "The id of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Name, nil
				}
				return nil, nil
			},
		},
		"username": &graphql.Field{
			Type:        graphql.String,
			Description: "The username of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Username, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "The email of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Email, nil
				}
				return nil, nil
			},
		},
		"address": &graphql.Field{
			Type:        addressField,
			Description: "The address of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Address, nil
				}
				return nil, nil
			},
		},
		"phone": &graphql.Field{
			Type:        graphql.String,
			Description: "The phone of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Phone, nil
				}
				return nil, nil
			},
		},
		"website": &graphql.Field{
			Type:        graphql.String,
			Description: "The website of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Website, nil
				}
				return nil, nil
			},
		},
		"company": &graphql.Field{
			Type:        companyField,
			Description: "The company of the user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(models.User); ok {
					return user.Company, nil
				}
				return nil, nil
			},
		},
	},
})

// UsersField is a field that returns all users
var UsersField = &graphql.Field{
	Type: graphql.NewList(userObject),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		return fetchUsers(), nil
	},
}

// Fetch all users from the REST API
func fetchUsers() []models.User {

	response, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var users []models.User

	err = json.Unmarshal(responseData, &users)

	if err != nil {
		log.Fatal(err)
	}

	return users
}
