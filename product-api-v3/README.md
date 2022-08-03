# Introduction
This is a product-api-service version 3. In this version, we will port all the old source to new source using GIN Framework Golang.  
There are a lot of content you need to remember.

# Table of Contents
1. [Part 1:](#part-1) CRUD REST Api on product service.


### Part 1
* CRUD REST Api on product service.
	- Understand the difference between... [Refer](https://pkg.go.dev/github.com/gin-gonic/gin#readme-parameters-in-path)
		- Parameter in path. [Solution 1](#parameter-in-path)
		- Querystring parameters
		- Multipart/Urlencoded Form
		- Map as querystring or postform parameters
	- Different between URL, URI, and structure of URL?
### Part 2
* Basic concepts in MongoDB.
	- Understand Indexes in Mongodb. [Refer](https://viblo.asia/p/tim-hieu-mot-so-loai-indexes-trong-mongodb-bWrZn4LQ5xw)
	- What is Aggregator? How to use it in Mongodb. [Refer]
* Useful command on Mongo
# Getting Started
```
docker-compose up -d
```
# Solution
### Parameter in path
