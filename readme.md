# Customer Merchant API

API for customer and merchant

## Installation

First you need to clone this project

```bash
  git clone https://github.com/achmadfff/Bank-API-Golang.git
```

Move to directory in cmd
```bash
  cd Bank-API-Golang
```

Then you need to migrate database first
```bash
  go run ./db/migrate/migrate.go
```
After that You need to seed the table
```bash
  psql -U postgres bank-api < db/seed.sql
```

## Environment Variables

To run this project, you will need to change the following environment variables to your .env file with your own configuration.

`PORT`

`DATABASE_HOST`

`DATABASE_PORT`

`DATABASE_NAME`

`DATABASE_USER`

`DATABASE_PASSWORD`

## API Reference

You can run this API using postman or you can run in apispec.json which i made in root directory.

#### Register Customer

- Method : `POST`
- Endpoint : `/vi/auth/register`
- Header :
    - Content-Type : `application/json`
    - Accept : `application/json`
- Body :

```json
{

    "name": "string",
    "email": "string",
    "password": "string"

}
```

- Response :

```json
{
  "data": null,
  "code": 200,
  "message": "Success Register Customer"
}
```

#### Login Customer

- Method : `POST`
- Endpoint : `/vi/auth/login`
- Header :
    - Content-Type : `application/json`
    - Accept : `application/json`
- Body :

```json
{
    "email" : "email@gmail.com",
    "password" : "password"
}
```

- Response :

```json
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjU0ZGQ3MDdjLThhZmItNDA3Ni1hZGMyLWEwOTFkMGU5NDVmOSIsImVtYWlsIjoic2FsbWFAZ21haWwuY29tIiwiZXhwIjoxNjc2ODY3MzgzfQ.O6u6KjCv-TnKudq1Rx98z7ykyxcGnjly2-ESujOi-AM"
    },
    "code": 200,
    "message": "Success Login"
}
```
#### Logout Customer

- Method : `POST`
- Endpoint : `/vi/auth/logout`
- Header :
    - Content-Type : `application/json`
    - Accept : `application/json`
    - Authorization : `JWT Bearer`
- Body : No need, Just Header Authorization Bearer Token

- Response :

```json
{
    "data": null,
    "code": 200,
    "message": "Success Logout"
}
```

#### Create Transaction/Transfer

Who made transaction is refer to token which sent in authorization.
- Method : `POST`
- Endpoint : `/v1/payment/transaction`
- Header :
    - Content-Type : `application/json`
    - Accept : `application/json`
    - Authorization : `JWT Bearer`
- Body :

```json
{
    "amount" : 150000,
    "merchantId" : "87c43149-8757-4ecb-b99b-e60affd1eeb9"
}
```

- Response :

```json
{
  "data": null,
  "code": 200,
  "message": "Success Create Transaction"
}
```
#### Get Report Transaction

- Method : `GET`
- Endpoint : `/v1/payment/report`
- Header :
    - Content-Type : `application/json`
    - Accept : `application/json`
- Body : No Body
```

- Response :

```json
{
    "data": {
        "data": [
            {
                "ID": "8f6f397f-654d-47ed-8bb6-ef3dca7e13fd",
                "Amount": 100000,
                "UserId": "e00df198-75e6-4785-ba0d-3d4fb5c1c2f8",
                "MerchantId": "bfdd663e-8167-45f3-91b2-b4cb11bf46c6",
                "CreatedAt": "2023-02-19T18:01:00.679263Z",
                "UpdatedAt": "2023-02-19T18:01:00.679263Z"
            }
        ],
        "meta": {
            "page": 1,
            "limit": 10,
            "total": 1
        }
    },
    "code": 200,
    "message": "OK"
}
```