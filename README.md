# GraphQL abstraction for REST API Endpoints

This GraphQL server serves as the API Gateway to bring several REST Endpoints together.

#### Features
- Fetch data from x number of services, serve it in one place

## How to launch the server for development
To develop this GraphQL server further, follow these steps:

1. Pull the repository
2. Run ``go mod download`` to install the dependencies
3. Run ``go run main.go`` to launch the server
4. Use any GraphQL development kit and set the endpoint to ``http://localhost:4000/graphql``