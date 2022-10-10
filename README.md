# Financial Server and API
A graphql and go based application program interface for the submission of grant receivables, invoices, and mileage requests.

# Features
The Application Program Interface (API) uses regular expressions to ensure that users are originating accounts from within the organization (for the internal website), and allows anyone with a google account to create an account (using the public facing api).
The API endpoint is hosted on the google cloud platform using a combination of google cloud containers and google cloud run functions. Automatic deployments are set up with some prior authoriztions on the google cloud console, utilization of a service account, and an automated YAML script that creates a docker image on a cloud based container instance.

# Tech Stack
1. Google UUID
  - Unique ID Generation
  - https://github.com/google/uuid
2. Go - Chi
  - Routing handler for GraphQL api calls
  - https://github.com/go-chi/chi
3. Mongo Drive
  - Database Connection to MongoDB Atlas Database
  - https://pkg.go.dev/go.mongodb.org/mongo-driver
4. JWT - Go
  - Token generation and authentication for user routes, and context for created requests
  - old repository https://github.com/dgrijalva/jwt-go
  - new repositor https://github.com/golang-jwt/jwt
5. GoLang
  - General Object Oriented language for the application
  - https://golang.google.cn/
6. GraphQL Go
  - Used in the creation of the graphql based schema and handles route requests
  - https://github.com/graphql-go/graphql
7. Google Cloud Run
  - Where the application graphql route is hosted to interact with the front end facing website
  - https://cloud.google.com/run/
8. Google Cloud Container
  - Stores the images created via a DockerFile script to be deployed on the google cloud run instance
  - https://cloud.google.com/containers/

# Connected Repository
The front facing web application repository can be found @ https://github.com/thomps9012/finance-api-interface
