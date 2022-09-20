
# Introduction

# Table of contents
* Initialize Mongodb image on docker by command line. [here](#mongodb-image-on-docker)
# Getting Started
To run this service
```
make run
```

### MongoDB image on Docker 
More details: 
    https://www.youtube.com/watch?v=DzyC8lqbjC8
#### Step 1: Initialize Mongodb Image
To start Mongo image
```
docker run -p 27017:27017 --name mdb mongo
```

#### Step 2: Execute MongoDB
To execute and enter to terminal mongo
```
Syntax: 
    docker exec -it [Container ID] [Image Type]

Example:
    docker exec -it 397f7a8bbe75 mongo
```

#### Step 3: Operation with MongoDB
Show all database on MongoDB
```bash
> show databases;

or clear screen
> cls
```

Switch to another database, or create db if it's not exist
```bash
> use test

check db
> db
```

Create collection, and insert data.
```bash
Case 1: Insert with empty data
> db.employees.insertMany([{}])

Case 2: Insert with multiple data
> db.employees.insertMany([{name: "Van Thong", "ssn":1995},{name: "Hussein", "ssn":111}])


Case 3: Insert by variable
> a = [
... {
... name : "giao linh",
... ssn : 1994
... }
... ]
[ { "name" : "giao linh", "ssn" : 1994 } ]
> db.employees.insertMany(a)
```

Find all documents in collection
```
> db.employees.find()
```

#### Step 4: Create mongo client
```bash
docker run -p 3000:3000 --name mclient mongoclient/mongoclient
```


