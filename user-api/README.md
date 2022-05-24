# Introduction
This user-api folder is a user microservice. It include some below feautures:
* Understand about user access to DB.
* Understand about RESTfull API
* Apply Singleton design pattern for DB.
* Understand how to authenticate user with JWT

# Questions
* How do you implement CRUB API with RESTful?
* How do you open MongoDB Driver on local machine?
* When you need to generate JWT Token for a user?
* How do you enable logger for GIN Framework? or Server? What the difference of them?
* In building-microservice-go book, Chapter 8: Security, it remind that we never storing passwords in plain text in a datastore. How do you implement this feature? Refer: [here]()

# Project Structure
```
├───common
├───config
├───controllers
├───daos
├───databases
├───models
└───utils
```
Those folders contain:
* **common:** handle common task such as setting logger, filter config ...
* **config:** configuration for project such as logger, mongodb, jwt...
* **controllers:** where to expose RESTful API for client.
* **daos:** where data access to object on mongoDB.
* **databases:** storage for DB.
* **models:** corresponds to all the data-related logic that the user works with.
* **utils:** implement some utilities for server


## Getting Started



## Test server
To understand this service, please test server with APIs below
### Login With User
Authenticate with your User.  
![ảnh](https://user-images.githubusercontent.com/50081052/170007415-33c77eeb-755a-4f99-821d-0fa804dec0c4.png)


### Register a new User
Register a new user.  
![ảnh](https://user-images.githubusercontent.com/50081052/170009455-3d0ec92e-3fb1-42d6-bcfc-845f9f5b2ff6.png)




## Answer Questions
### Hashing password
**bcrypt** is another popular method of hashing passwords.  
1. To hash a password with bcrypt we use the **GenerateFromPassword()** method:
```go
func Hash(password string) (string, error) {
    // hass password using bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
```

2. To check the equality of a bcrypt hash we can not call **GenerateFromPassword()** again with the given password and compare the output to the hash we have stored as **GenerateFromPassword()** will create a different hash every time it is run. To compare equality we need to use the **CompareHashAndPassword()** method:
```go
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
```

3. To avoid SQL injection error, we can implement **Santize()** function
```go
func Santize(data string) string{
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
```
More details hacking with SQL injection. Refer: [here](https://www.meisternote.com/app/note/uMUTsPEJzQbx/sql-injection)
