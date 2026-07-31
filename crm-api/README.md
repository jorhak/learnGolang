# 1 Ejecutar levamtar servicio
```
go run ./cmd/api
```
# 2 Ejecutar endpoints
## 2.1 GetCustomers
```
curl -i localhost:8080/customers
```
## 2.1 CreateCustomers
```
curl -X POST localhost:8080/customers \
-d '{"id":2,"name":"Saludos","email":"saludos@gmail.com"}'
```
## 2.2 GetByID
```
curl -i localhost:8080/customers/3
```
## 2.3 DeleteCustomer 
```
curl -i "localhost:8080/delete?id=3"
```
o 
```
curl -i -X DELETE "localhost:8080/delete?id=3"
```
```
```
```

```


