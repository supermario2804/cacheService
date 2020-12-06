# cacheDataService

It is the **cacheDataService**. Feature of this service is as below : 
 - Store cache with REST call
 - Fetch cache with REST call
 - Support paginated result
 - Data reload on the notification via RabbitMQ
 - Database backup at every 30 minute

# Tech stack
- Go
- Redis 
- RabbitMQ
- Docker
- Docker Compose

# How to Run?
- Open terminal at the folder with **docker- compose.yml**
- run `docker-compose build`
- run `docker-compose up`

# API Endpoints 
 - on local base url is : ( localhost:8090 )
 - `/api/set`
 - `/api/get`
 - `/api/setPage`
 - `/api/getPage`

### 1. ../api/set

#### Method : `Post`
#### URL : `localhost:8090/api/set`
#### Description: The following request sets cache for `Table:emp` Row `PrimaryKey:1` and `columns:FirstName,LastName,Phone`
#### Request JSON :
```json
{
    "table":"emp",
    "pk":"1",
    "data": {
        "FirstName":"Nikhil",
        "LastName":"Patel",
        "Phone":"123456"
    }
}
```

### 2. ../api/get

#### Method : `Post`
#### URL : `localhost:8090/api/get`
#### Description: The following request gets cache for `Table:emp` Row `PrimaryKey:1`
#### Request JSON :
```json
{
    "table":"emp",
    "pk":"1"
}
```

### 3. ../api/setPage

#### Method : `Post` 
#### URL : `localhost:8090/api/setPage`
#### Description: The following request sets page cache for `Page:transaction`,`PageNumber:1` and `sortby:asc`.
#### Request JSON :
```json
{
    "title":"transaction",
    "pageNumber":"1",
    "sortBy":"asc",
    "data": [
        {
            "name":"Amazon",
            "amount":"634",
            "token":"A004"
        },

        {
            "name":"Flipkart",
            "amount":"712",
            "token":"A005"
        }

    ]
}
```

### 4. ../api/getPage

#### Method : `Post` 
#### URL : `localhost:8090/api/getPage`
#### Description: The following request gets page cache for `Page:transaction`,`PageNumber:1` and `sortby:asc`.
#### Request JSON :
```json
{
    "title": "transaction",
    "pageNumber": "1",
    "sortBy": "asc"
}
```

### Reload From Database
- To reload the data from database the rabbit-mq server should be running.
- Assuming server is running on the localhost:5672 
- The sample code to send notification to reload (row with `56` primary key from `Employee` table) 
   service is as below:

```Go

		//Don't do this in production, this is for testing purposes only.
		url = "amqp://guest:guest@localhost:5672"

	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}
	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	// We create a message to be sent to the queue.
	// It has to be an instance of the aqmp publishing struct
	message := amqp.Publishing{
		Body: []byte("Employee_56"),
	}

	// We publish the message to the exahange we created earlier
	err = channel.Publish("events", "random-key", false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}
```


