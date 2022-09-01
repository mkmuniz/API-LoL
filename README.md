# API-LoL

**About**

An api to return data about champions free, items and etc in League of Legends game.

**How to run**

1 - Clone the repository

2 - Open the repository and type `go mod tidy` to avoid errors

3 - Create a file named `config.toml` and add the following env variables:

```
[api]
key=""
port=""
```

4 - Run with the following command `go run main.go`

5 - Acess `localhost:PORT_DECLARED`

6 - You can see the free champions in `/champions/free` route
