# Matchmaker

Start matchmaker server
```
go run /example/run.go
```

Add player to matchmaker search

```
curl -H "Content-Type: application/json" --data '{ "name": "A", "skill" : 100 }' http://localhost:8080/search
```
