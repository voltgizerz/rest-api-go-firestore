# REST API Go With Firebase Cloud Firestore
Simple REST APIs with database firebase cloud firestore
# BUILD WITH
- Go 1.20 or higher
- JWT Authorization
- [Cloud Firestore](https://firebase.google.com/docs/firestore/quickstart) Database

# How to Run Project

Follow these steps to run the project:

1. Create a Cloud Firestore database by following this [Guide](https://firebase.google.com/docs/firestore/quickstart).

2. Export your service account credentials on the project settings in the Firebase console. Download the service-account JSON file and place it in the credentials folder: `rest-api-go-firestore/config/credential/sa-sample.json`.

3. Update the constant value in `rest-api-go-firestore/database/database.go` named `SERVICE_ACCOUNT_CREDENTIAL_FILE_PATH` based on the location of your service account JSON file.

4. Run the application using the following command: `go run ./cmd/app.go`

# GET
Retrive auth token expires in 1 hour
- [api/token?client_id=sample&client_secret=BiquzG0JVY3pWPrh8xiVPkbNXyx20Gmn](localhost:8080/api/token?client_id=sample&client_secret=BiquzG0JVY3pWPrh8xiVPkbNXyx20Gmn)
- Response :

```json 
{
    "status": 200,
    "message": "Token generated successfully",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "type": "Bearer"
    }
}
```

# GET
Retrive specific user data
- [api/users/:docRefID](localhost:8080/api/users/:docRefID)
- Header :
```sh
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
- Response :

```json 
{
    "status": 200,
    "message": "User data retrieved successfully",
    "data": {
        "DocRefID": "iCOmqSXxG8bJZfVJ8iuj",
        "FirstName": "Gerald",
        "LastName": "Fisher",
        "Username": "BettyHolmes",
        "Email": "quia@Flipopia.name",
        "CCNumber": "4916791396904137",
        "CCType": "Discover",
        "Country": "Uruguay",
        "City": "Santa Monica",
        "Currency": "Albania Leke"
    }
}
```

# GET
Retrive all users data
- [api/users](localhost:8080/api/users)
- Header :
```sh
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
- Response :

```json 
{
    "status": 200,
    "message": "Users data retrieved successfully",
    "data": [
        {
            "DocRefID": "76KbFdZd8kVCawLNBUle",
            "FirstName": "Debra",
            "LastName": "Washington",
            "Username": "eum_quia",
            "Email": "ipsam_ex@Aivee.org",
            "CCNumber": "5237378900055256",
            "CCType": "VISA",
            "Country": "Dominican Republic",
            "City": "Orinda",
            "Currency": "Jordan Dinars"
        },
        {
            "DocRefID": "M0tbZqmFkS3J42xWJ697",
            "FirstName": "Ashley",
            "LastName": "Hamilton",
            "Username": "hNguyen",
            "Email": "cHunter@Zoonoodle.com",
            "CCNumber": "379957517654574",
            "CCType": "MasterCard",
            "Country": "Kazakhstan",
            "City": "Ione",
            "Currency": "United Kingdom Pounds"
        },
        {
            "DocRefID": "XeidVNtUgURCANof8jxc",
            "FirstName": "James",
            "LastName": "Powell",
            "Username": "mollitia",
            "Email": "uOliver@Vipe.biz",
            "CCNumber": "5347663249612785",
            "CCType": "Discover",
            "Country": "Niue",
            "City": "Ione",
            "Currency": "Bahrain Dinars"
        }
    ]
}
```

# POST
Create new user
- [api/users](localhost:8080/api/users)
- Header :
```sh
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
- Body :

```json 
{
    "FirstName": "James",
    "LastName": "Powell",
    "Username": "mollitia",
    "Email": "uOliver@Vipe.biz",
    "CCNumber": "5347663249612785",
    "CCType": "Discover",
    "Country": "Niue",
    "City": "Ione",
    "Currency": "Bahrain Dinars"
}
```

- Response :

```json 
{
    "status": 201,
    "message": "User data inserted successfully",
    "data": {
        "DocRefID": "DYyDOPZ8nJzZWB7Y965J"
    }
}
```

# PATCH
Update specific user data
- [api/users/:docRefID](localhost:8080/api/users/:docRefID)
- Header :
```sh
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
- Body :

```json 
{
    "FirstName": "James",
    "LastName": "Powell",
    "Username": "mollitia",
    "Email": "uOliver@Vipe.biz",
    "CCNumber": "5347663249612785",
    "CCType": "Discover",
    "Country": "Niue",
    "City": "Ione",
    "Currency": "Bahrain Dinars"
}
```

- Response :

```json 
{
    "status": 200,
    "message": "User data updated successfully"
}
```


# DELETE
Delete specific user data
- [api/users/:docRefID](localhost:8080/api/users/:docRefID)
- Header :
```sh
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
- Response :

```json 
{
    "status": 200,
    "message": "User data deleted successfully"
}
```