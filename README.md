#### RssFeed Task

The system is implemented to automatically fetch Rss feed automatically at regular
intervals. 
At the start of the system, to Go routines are started separately.

1. The first go routing starts the server for searching the saved feeds
2. The second is used to retrieve the feed at regular intervals and the
response save in the database. 

### The components

i.  Mongodb database:
A cloud based mongodb database was implemented with the system as
it supports search and no need for installations.


### Methods
##### GetRss()
This method retrieves rss feed from any url provided and return it to a channel
##### ReceiveFromChannel()
This methods receive information from the channel, get the array of feeds loop through and return []interface {}

#### Spider()
It makes use of the above methods by passing multiple urls and passing the retrieved information into the channel


#### StartSpider()
This is the method responsible for running the spider methods at intervals.

#### ConnectDB()
This is the method responsible for connecting the system to the database. A mongodb collection is returned
 
#### SaveToDb()
Saving into the database is done in this method. An unordered form of saving is used in other to allow the check on duplicated and allow the procedure
to continue after the error


### Tests
All tests are placed in a different package called tests.
to run test use go test -v -coverpkg ./... ./...


### Route
The route for the search is "/search"
