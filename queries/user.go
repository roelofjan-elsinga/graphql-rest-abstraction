package queries

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/roelofjan-elsinga/graphql-rest-abstraction/models"
)

// UserField is a field that returns a single user
var UserField = &graphql.Field{
	Type: userObject,
	Args: graphql.FieldConfigArgument{
		"user_id": &graphql.ArgumentConfig{
			Description: "The ID of the user you want to retrieve",
			Type:        graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		return fetchUsers(), nil
	},
}

// Find a single user for the given userID
func fetchUser(userID int) models.User {

	response, err := http.Get("https://jsonplaceholder.typicode.com/users/" + strconv.Itoa(userID))

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	err = json.Unmarshal(responseData, &user)

	if err != nil {
		log.Fatal(err)
	}

	return user

}
