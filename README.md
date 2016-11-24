## Run the server
```
git clone https://github.com/incu6us/graphQlGoTest.git
cd graphQlGoTest/src
go test
go run {main,graphql,models,handler}.go
```

## Examples to execute
### Time API
```
curl -XPOST http://localhost:8081/graphql -H 'Content-Type: application/json' -d '{"query": "{timeQuery}"}'
```

### Date API
```
curl -XPOST http://localhost:8081/graphql -H 'Content-Type: application/json' -d '{"query": "{dateQuery}"}'
```

### or with GET Method
```
curl -XGET http://localhost:8081/graphql?query={timeQuery}
```


