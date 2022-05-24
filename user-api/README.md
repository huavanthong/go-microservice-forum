# Introduction
This user-api folder is a user microservice. It include some below feautures:
* Understand about user access to DB.
* Understand about RESTfull API
* Apply Singleton design pattern for DB.
* Understand how to authenticate user with JWT


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


# Getting Started



# Test server
To understand this service, please test server with APIs below
### Login With User
Authenticate with your User.  
![ảnh](https://user-images.githubusercontent.com/50081052/170007415-33c77eeb-755a-4f99-821d-0fa804dec0c4.png)


### Register a new User
Register a new user.  
![ảnh](https://user-images.githubusercontent.com/50081052/170009455-3d0ec92e-3fb1-42d6-bcfc-845f9f5b2ff6.png)


