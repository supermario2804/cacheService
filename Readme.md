--> change the conf file of redis
--> in redis change the directory of backup
--> like dbfilename, and dir
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
> on local base url is : ( localhost:8090 )
>.../api/set
>.../api/get
>.../api/setPage
>.../api/getPage

## 1. ../api/set

### Method : Post 
### URL : localhost:8090/api/set
### Request JSON :
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
