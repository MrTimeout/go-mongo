# Learning MongoDB

## Show commands

```
> show databases
> show dbs
# will show databases

> show collections
> show tables
> db.getCollectionNames() (returns an array of its names)
> db.getCollectionInfos() (returns an array with its names and a lot of information)
# will show collections which are in the actual database

> show users
# will show users in the mongodb

> show roles
# will show roles

> db
# Will prompt the actual database

> use {DatabaseName}
# Will change the actual database to another one
```

## How to export and import data

When using a JSON:

- `mongoimport`
  + `mongoimport --uri <URI> --drop dump.json` the drop means dropping all the data in the actual database so we don't get duplicates.
- `mongoexport`
  + `mongoexport --uri <URI> --collection=<collection-name> --out=<filename>.json`

When using a BSON:

- `mongorestore`
  + `mongorestore --uri <URI> --drop dump.bson`
- `mongodump`
  + `mongodump --uri <URI>`

srv in URI connections means secure

`mongodb+srv://<user>:<password>@clusterIP.mongodb.net/database`

[Reference](https://www.mongodb.com/docs/database-tools/mongoimport/#compatibility)

## ObjectID

ObjectID consists of 12 bytes:

- 4-bytes: timestamp, representing the ObjectId's creation, measured in seconds sinde the Unix section.
- 5-bytes: random value generated once per process. This random value is unique to the machine and process.
- 3-bytes: incrementing counter, initialized to a random value.

If an integer value is used to create an ObjectId, the integer replaces the timestamp.

ObjectId() can accept:

- hexadecimal
- integer

Methods:

- getTimestamp()
- toString()
- valueOf()

example: `x = ObjectId(32) // returns 00000020<random-value><counter-from-random-value>`

[Help](https://www.mongodb.com/docs/manual/reference/method/ObjectId/#objectid)