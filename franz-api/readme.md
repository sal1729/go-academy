# franz-api

An api wrapping the franz-datastore.

To run do `go run .`
Then to test the request types

### **GET** (read)
- all: `curl -X GET http://localhost:8080/api`
- by status: `curl -X GET "http://localhost:8080/api?status=To%20Do"`
- by task wasn't implemented in the api

### **POST** (create)
```shell
curl -X POST http://localhost:8080/api \
  -H "Content-Type: application/json" \
  -d '{"action": "create", "task": "Clone the car", "status": "To Do"}'
```

### **PUT** (update)
```shell
curl -X PUT http://localhost:8080/api \
  -H "Content-Type: application/json" \
  -d '{"action": "update", "task": "Clone the car", "status": "Done"}'
```
Only single task updating was implemented for the api

### **DELETE** (delete)
Single task:
```shell
curl -X DELETE http://localhost:8080/api \
  -H "Content-Type: application/json" \
  -d '{"action": "delete", "task": "Clone the car", "status": ""}'
```
By status:
```shell
curl -X DELETE http://localhost:8080/api \
  -H "Content-Type: application/json" \
  -d '{"action": "delete", "task": "all", "status": "To Do"}'
```