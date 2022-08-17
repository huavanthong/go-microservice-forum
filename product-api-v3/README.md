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
		- Map asj

### Part 2 
* Indexes in Mongodb.
	- [Measure time](#measure-time) Measure query time in Mongodb. [Refer](https://viblo.asia/p/tim-hieu-ve-index-trong-mongodb-924lJL4WKPM)
	- [Indexes](#indexes-in-mongodb) How to use Indexes in MongoDB. [Refer](https://viblo.asia/p/su-dung-index-trong-sql-query-1ZnbRlPQR2Xo)


# Designer
### Measure time
* Step 1: Start Mongodb Shell
```bash
> mongosh
Current Mongosh Log ID: 62f770facd6135a38bfd189f
Connecting to:          mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.5.0
Using MongoDB:          5.0.7
Using Mongosh:          1.5.0
......................................

```
* Step 2: Connect to collection, and use your db
```
> show dbs
> use go-microservice
```

* Step 3: Show collection
```
> show collections
```
* Step 4: insert the sample data
```s
> use test     // Use database test to insert the sample data
    for(var i = 0; i < 1000000; i++) {
        db.users.insert({
            i: i,
            username: 'user' + i,
            age: Math.floor(Math.random() * 100)
        });
    }
```
* Step 4: query to find your document
```s
> db.users.find({username: 'user112'}).explain("executionStats")["executionStats"]
{
	"executionSuccess" : true,
	"nReturned" : 1,
	"executionTimeMillis" : 269,
	"totalKeysExamined" : 0,
	"totalDocsExamined" : 1000000,
	"executionStages" : {
		"stage" : "COLLSCAN",
		"filter" : {
			"username" : {
				"$eq" : "user112"
			}
		},
		"nReturned" : 1,
		"executionTimeMillisEstimate" : 211,
		"works" : 1000002,
		"advanced" : 1,
		"needTime" : 1000000,
		"needYield" : 0,
		"saveState" : 7813,
		"restoreState" : 7813,
		"isEOF" : 1,
		"invalidates" : 0,
		"direction" : "forward",
		"docsExamined" : 1000000
	}
}
```
### Indexes in MongoDB.
* Step 1: Create indexes, and check indexes
```s
> db.users.createIndex({username: 1}) // Đánh index cho username theo thứ tự tăng dần, -1 là giảm dần.

# Check your indexes
> db.users.getIndexes() // Xem index nào đã có.
```
* Step 2: Check time after create indexes
```s
> db.users.find({username: 'user112'}).explain("executionStats")["executionStats"]
```