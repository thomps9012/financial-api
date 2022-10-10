# Financial Server and API
A graphql and go based application program interface for the submission of grant receivables, invoices, and mileage requests.

# Features
The Application Program Interface (API) uses regular expressions to ensure that users are originating accounts from within the organization (for the internal website), and allows anyone with a google account to create an account (using the public facing api).
The API endpoint is hosted on the google cloud platform using a combination of google cloud containers and google cloud run functions.
Automatic deployments are set up on the google cloud console, with custom permissions, the utilization of a service account, and an automated YAML script that creates and deployes a docker image to a cloud based container instance.

# Tech Stack
1. [Google UUID](https://github.com/google/uuid)
  - Unique ID Generation
2. [Go - Chi](https://github.com/go-chi/chi)
  - Routing handler for GraphQL api calls
3. [Mongo Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
  - Database Connection to MongoDB Atlas Database
4. [JWT - Go](https://github.com/golang-jwt/jwt)
  - Token generation and authentication for user routes, and context for created requests
  - [old repository link](https://github.com/dgrijalva/jwt-go)
5. [GoLang](https://golang.google.cn/)
  - General Object Oriented language for the application
6. [GraphQL Go](https://github.com/graphql-go/graphql)
  - Used in the creation of the graphql based schema and handles route requests
7. [Google Cloud Run](https://cloud.google.com/run/)
  - Where the application graphql route is hosted to interact with the front end facing website
8. [Google Cloud Container](https://cloud.google.com/containers/)
  - Stores the images created via a DockerFile script to be deployed on the google cloud run instance

# Connected Repository
The front facing web application repository can be found @ https://github.com/thomps9012/finance-api-interface
