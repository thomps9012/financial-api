# Financial Server and API
A golang and mongodb based application program interface for the submission of grant receivables, invoices, and mileage requests.

# Features
The Application Program Interface (API) uses regular expressions to ensure that users are originating accounts from within the organization (for the internal website), and allows anyone with a google account to create an account (using the public facing api).
The API endpoint is hosted on the google cloud platform using a combination of google cloud containers and google cloud run functions.
Automatic deployments are set up on the google cloud console, with custom permissions, the utilization of a service account, and an automated YAML script that creates and deploys a docker image to a cloud based container instance.

# Tech Stack
1. [Google UUID](https://github.com/google/uuid)
  - Unique ID Generation
2. [Mongo Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
  - Database Connection to MongoDB Atlas Database
3. [JWT - Go](https://github.com/golang-jwt/jwt)
  - Token generation and authentication for user routes, and context for created requests
  - [old repository link](https://github.com/dgrijalva/jwt-go)
4. [GoLang](https://golang.google.cn/)
  - General Object Oriented language for the application
5. [Fiber](https://docs.gofiber.io/)
  - General http routing
6. [Google Cloud Run](https://cloud.google.com/run/)
  - Where the application graphql route is hosted to interact with the front end facing website
7. [Google Cloud Container](https://cloud.google.com/containers/)
  - Stores the images created via a DockerFile script to be deployed on the google cloud run instance

# Connected Repository
The front facing web application repository can be found @ https://github.com/thomps9012/finance-api-interface
