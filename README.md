# REST API Go With Firebase Cloud Firestore
Simple REST APIs with database firebase cloud firestore
# Build With
- Go 1.23.2 or higher
- JWT Authorization
- [Cloud Firestore](https://firebase.google.com/docs/firestore/quickstart) Database

# How to Run Project

Follow these steps to run the project:

1. Create a Cloud Firestore database by following this [Guide](https://firebase.google.com/docs/firestore/quickstart).

2. Export your service account credentials on the project settings in the Firebase console. Download the service-account JSON file and convert the JSON into a single-line (minified) format and store it in an environment file named `SERVICE_ACCOUNT_FIREBASE`.

3. Run the application using the following command: `go run ./cmd/app.go`

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
        "doc_ref_id": "iCOmqSXxG8bJZfVJ8iuj",
        "first_name": "Gerald",
        "last_name": "Fisher",
        "username": "BettyHolmes",
        "email": "quia@Flipopia.name",
        "cc_number": "4916791396904137",
        "cc_type": "Discover",
        "country": "Uruguay",
        "city": "Santa Monica",
        "currency": "Albania Leke"
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
            "doc_ref_id": "76KbFdZd8kVCawLNBUle",
            "first_name": "Debra",
            "last_name": "Washington",
            "username": "eum_quia",
            "email": "ipsam_ex@Aivee.org",
            "cc_number": "5237378900055256",
            "cc_type": "VISA",
            "country": "Dominican Republic",
            "city": "Orinda",
            "currency": "Jordan Dinars"
        },
        {
            "doc_ref_id": "M0tbZqmFkS3J42xWJ697",
            "first_name": "Ashley",
            "last_name": "Hamilton",
            "username": "hNguyen",
            "email": "cHunter@Zoonoodle.com",
            "cc_number": "379957517654574",
            "cc_type": "MasterCard",
            "country": "Kazakhstan",
            "city": "Ione",
            "currency": "United Kingdom Pounds"
        },
        {
            "doc_ref_id": "XeidVNtUgURCANof8jxc",
            "first_name": "James",
            "last_name": "Powell",
            "username": "mollitia",
            "email": "uOliver@Vipe.biz",
            "cc_number": "5347663249612785",
            "cc_type": "Discover",
            "country": "Niue",
            "city": "Ione",
            "currency": "Bahrain Dinars"
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
    "first_name": "James",
    "last_name": "Powell",
    "Username": "mollitia",
    "Email": "uOliver@Vipe.biz",
    "cc_number": "5347663249612785",
    "cc_type": "Discover",
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
        "doc_ref_id": "DYyDOPZ8nJzZWB7Y965J"
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
    "first_name": "James",
    "last_name": "Powell",
    "username": "mollitia",
    "email": "uOliver@Vipe.biz",
    "cc_number": "5347663249612785",
    "cc_type": "Discover",
    "country": "Niue",
    "city": "Ione",
    "currency": "Bahrain Dinars"
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