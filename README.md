# MemoryStore

This is the key value store serving through the rest api.

The application saves the current information of the memorystore under the ./savedfiles (can be changed from the config file) folder every 10 minutes (can be changed from the config file). Then when the app is started, it updates the memory from the most recent file.

If the memorystore is flushed, it clears the memory and moves the backup files under ./savedfiles/archive. After the flush, the files are not read and stored again under the archive, they are only used for backing up.

## Install

    git clone https://github.com/mhgenc/MemoryStore.git

## Run the app

    go run main.go

## Run the tests

    go test -v

## Online endpoint

You can access the application deployed to heroku from the link below.
    
    https://memorystore.herokuapp.com

# REST API

The REST API to the memory store is described below.

## Create a new key value

### Request

`POST /api/set`

    curl --location --request POST 'https://memorystore.herokuapp.com/api/set' \
          --header 'Content-Type: application/json' \
          --data-raw '{
                        "key" : "firstname",
                        "value": "mehmet"
                      }'

### Response

    HTTP/1.1 200 OK
   
    {
    "status": 200,
    "data": {
        "key": "firstname",
        "value": "mehmet"
    }
}

## Get value

### Request

`GET /api/get`

    curl --location --request GET 'https://memorystore.herokuapp.com/api/get?key=firstname'

### Response

    HTTP/1.1 200 OK

    {
    "status": 200,
    "data": {
        "key": "firstname",
        "value": "mehmet"
    }
}

## Flush memory store

### Request

`DELETE /flush`

    curl --location --request DELETE 'https://memorystore.herokuapp.com/api/flush'

### Response

    HTTP/1.1 200 OK
    
    {
    "status": 200,
    "data": {}
    }

## Get a non-existent key

### Request

`GET /api/get?key=foo`

    curl --location --request GET 'https://memorystore.herokuapp.com/api/get?key=foo'

### Response

    HTTP/1.1 404 Not Found
    
    {
    "status": 404,
    "error": "key not found"
    }
