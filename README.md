# Soccer API

Soccer API is a dummy REST API that can create a team, add player to a team, get a team including with the players, and get a player

## How to run

Clone the repository

```
git clone git@github.com:guntoroyk/soccer-api.git
```

Run the prebuilt binary on `./bin` folder based on your OS and machine architecture with the following format :

```
./bin/soccer-api-[OS]-[arch]
```

Example:

```
./bin/soccer-api-darwin-arm64

./bin/soccer-api-linux-arm64
```

The API server will live on default port `8000`

```
➜  soccer-api git:(main) ✗ ./bin/soccer-api-darwin-arm64

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.9.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8000

```

## Build

To build the binary, you can run this command:

```
go mod vendor
make build
```

Or if you want build to specific machine, you can run the command `make build-[OS]-[arch]`

```
make build-linux-amd64
```

To build to all machines, run:

```
make build-all
```

The binary will be stored at folder `./bin`

## Run using go

To run the API using go, run this command inside `soccer-api` folder:

```
go mod vendor
go run main.go
```

## REST API Documentation

The API has 5 endpoints:

1. [Create a Team](#create-a-team)
2. [Add a player to a team](#add-a-player-to-a-team)
3. [Get list of teams](#get-list-of-teams)
4. [Get a team](#get-a-team)
5. [Get a player of a team](#get-a-player-of-a-team)

The API will return response with the following format:

```
{
  "code": number,
  "data": any,
  "error": string
}
```

### Create a Team

#### Request

`POST /api/teams`

```
curl -X POST \
  'localhost:8000/api/teams' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "Persib"
}'
```

#### Response

```
{
  "code": 201,
  "data": {
    "id": 1,
    "name": "Persib",
    "players": []
  }
}
```

### Add a player to a team

#### Request

`POST /api/teams/:id/players`

```
curl -X POST \
  'localhost:8000/api/teams/1/players' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "Christiano Ronaldo"
}'

```

#### Response

```
{
  "code": 201,
  "data": {
    "id": 1,
    "name": "Christiano Ronaldo"
  }
}
```

```
{
  "code": 400,
  "error": "player already in a team"
}
```

```
{
  "code": 404,
  "error": "team not found"
}
```

### Get list of teams

#### Request

`GET /api/teams`

```
curl -X GET \
  'localhost:8000/api/teams' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```

#### Response

```
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "name": "Persib",
      "players": [
        {
          "id": 1,
          "name": "Christiano Ronaldo"
        }
      ]
    }
  ]
}
```

### Get a team

#### Request

`GET /api/teams/:id`

```
curl -X GET \
  'localhost:8000/api/teams/1' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```

#### Response

```
{
  "code": 200,
  "data": {
    "id": 1,
    "name": "Persib",
    "players": [
      {
        "id": 1,
        "name": "Christiano Ronaldo"
      }
    ]
  }
}
```

```
{
  "code": 404,
  "error": "team not found"
}
```

### Get a player of a team

#### Request

`GET /api/teams/:id/players/:playerId`

```
curl -X GET \
  'localhost:8000/api/teams/1/players/1' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```

#### Response

```
{
  "code": 200,
  "data": {
    "id": 1,
    "name": "Christiano Ronaldo"
  }
}
```

```
{
  "code": 404,
  "error": "team not found"
}
```

```
{
  "code": 404,
  "error": "player not found"
}
```
