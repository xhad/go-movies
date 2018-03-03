## Go Movies Training Bot


Mission: 

* Create Sererless API Gateway Endpoint with Go Lambdas
* Post Request of Genre ID should return list of movies


### Dependencies
```
go get github.com/aws/aws-lambda-go/lambda # for handler registration
go get github.com/stretchr/testify # for unit tests
```

Get and API key from TheMoviesDB:
```
export API_KEY=<your key>
```

Run tests:
```
go test
```

### Build executable for linux

```
GOOS=linux go build -o main main.go
```

Zip the package into a deployment package:

```
zip deployment.zip main
```

Use the AWS CCLI to create a new lambda function:
```
aws lambda create-function \
--region us-east-1 \
--function-name DiscoverMovies \
--zip-file fileb://./deployment.zip \
--runtime go1.x \
--role arn:aws:iam::<account-id>:role/<role> \
--handler main
```