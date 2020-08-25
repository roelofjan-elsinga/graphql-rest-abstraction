# GraphQL abstraction for REST API Endpoints

This GraphQL server serves as the API Gateway to bring several REST Endpoints together. This repository is an example for my blog post: ["GraphQL: Centralize existing REST API endpoints for easier development"](https://roelofjanelsinga.com/articles/graphql-centralize-existing-rest-api-endpoints).
#### Features
- Fetch data from a REST API ([https://jsonplaceholder.typicode.com](https://jsonplaceholder.typicode.com/)) and make this available as a GraphQL resource.

## How to launch the server for development
To develop this GraphQL server further, follow these steps:

1. Pull the repository
2. Run ``go mod download`` to install the dependencies
3. Run ``go run main.go`` to launch the server
4. Use any GraphQL development kit and set the endpoint to ``http://localhost:4000/graphql``
