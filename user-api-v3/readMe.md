# Reference
CRUD RESTful API with Golang + MongoDB + Redis + Gin Gonic
#### Part 1
* API with Golang + MongoDB + Redis + Gin Gonic: Project Setup
    - API with Golang + MongoDB + Redis + Gin Gonic: Project Setup. [Refer](https://codevoweb.com/api-golang-mongodb-gin-gonic-project-setup)
	- Understand about method to design. [Refer](https://blog.logrocket.com/building-microservices-go-gin/)

#### Part 2
* Golang & MongoDB: JWT Authentication and Authorization
    - Golang & MongoDB: JWT Authentication and Authorization. [Refer](https://codevoweb.com/golang-mongodb-jwt-authentication-authorization)
    - [Solution](#solution-for-design-object-id) To create a object ID from Primitive.  [Refer](https://kb.objectrocket.com/mongo-db/how-to-find-a-mongodb-document-by-its-bson-objectid-using-golang-452)
    - [Solution] Design feature login attempt in Golang. [Refer](https://www.stackhawk.com/blog/golang-broken-authentication-guide-examples-and-prevention/)
#### Part 3
* API with Golang + MongoDB: Send HTML Emails with Gomail
    - API with Golang + MongoDB: Send HTML Emails with Gomail. [Refer](https://codevoweb.com/api-golang-mongodb-send-html-emails-gomail)

#### Part 4
* API with Golang, Gin Gonic & MongoDB: Forget/Reset Password
    - API with Golang, Gin Gonic & MongoDB: Forget/Reset Password. [Refer](https://codevoweb.com/api-golang-gin-gonic-mongodb-forget-reset-password)

#### Part 5
* Build Golang gRPC Server and Client: SignUp User & Verify Email
    - Build Golang gRPC Server and Client: SignUp User & Verify Email. [Refer](https://codevoweb.com/golang-grpc-server-and-client-signup-user-verify-email)
    - [Solution](#solution-to-use-a-temporary-to-store-verification-code-from-email) How to use temporary in MongoDB. 

#### Part 6
* Build Golang gRPC Server and Client: Access & Refresh Tokens
    - Build Golang gRPC Server and Client: Access & Refresh Tokens. [Refer](https://codevoweb.com/golang-grpc-server-and-client-access-refresh-tokens)

#### Part 7
* Build CRUD RESTful API Server with Golang, Gin, and MongoDB
    - Build CRUD RESTful API Server with Golang, Gin, and MongoDB. [Refer](https://codevoweb.com/crud-restful-api-server-with-golang-and-mongodb)

#### Part 8
* Build CRUD gRPC Server API & Client with Golang and MongoDB
    - Build CRUD gRPC Server API & Client with Golang and MongoDB. [Refer](https://codevoweb.com/crud-grpc-server-api-client-with-golang-and-mongodb)

#### Part 9
* Google OAuth Authentication React.js, MongoDB and Golang
    - Google OAuth Authentication React.js, MongoDB and Golang.  [Refer](https://codevoweb.com/google-oauth-authentication-react-mongodb-and-golang)

# Getting Started
### Solution for design object ID
At file: [auth.service.impl.go](./services/auth.service.impl.go).  

**Design 1:** If you want to insert document to collection, then insert object ID for your document.
```go
// Step 1: Insert document without objectID
	res, err := uc.collection.InsertOne(uc.ctx, &user)

// Step 2: get response and then insert ID
	query := bson.M{"_id": res.InsertedID}

// Step 3: find your result
	err = uc.collection.FindOne(uc.ctx, query).Decode(&newUser)
```
**Design 2:** If you want to setting objectID in user.model.go, and generate objectID, then insertOne to DB.
```go
// Step 1: Define a new objectID
	user.ID = primitive.NewObjectID()

// Step 2: Insert to db
	_, err := uc.collection.InsertOne(uc.ctx, &user)

// Step 3: make a query command.
	query := bson.M{"_id": user.ID}

// Step 4: find your result
    err = uc.collection.FindOne(uc.ctx, query).Decode(&newUser)
```

### Solution to use a temporary to store verification code from email
#### Requirement
- We need to implement feature to verify email from users once they register a new account.
- We will generate a verification code and send email to users.
- Once you generate a code, you will store it a temporary table in MongoDB, and after user verified successfully, we will delete it.
- Security: need to make sure the verification code is hacked by another source.
#### Design
At file: [auth.controller.go](./controllers/auth.controller.go).  
```go
// Step 1: User register a new account, then we will generate code send to email.
	code := randstr.String(20)
	verificationCode := utils.Encode(code)

// Step 2: After getting code, we will store it in temporary table
    ac.userService.UpdateUserById(newUser.ID.Hex(), "verificationCode", verificationCode)
    
// Step 3: In service, we will UpdateOne() to insert verifycode in table
	userId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: field, Value: value}}}}
	result, err := uc.collection.UpdateOne(uc.ctx, query, update)

// Step 4: Call verifyemail() API to verify email from code, and delete temporary table by using unset keyword.
	verificationCode := utils.Encode(code)

	query := bson.D{{Key: "verificationCode", Value: verificationCode}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "verified", Value: true}}}, {Key: "$unset", Value: bson.D{{Key: "verificationCode", Value: ""}}}}
	result, err := ac.collection.UpdateOne(ac.ctx, query, update)
```

