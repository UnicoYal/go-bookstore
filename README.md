<h2>A simple Golang API with Postgres Database<h2>

The code contains only one model BOOK<br />

To try this code just clone this repository, change the database setup and use Postman for requests<br />

<h2>Examples of usage</h2>

<h4>Request</h4>

```
GET localhost:9010/books
```

<h4>Response</h4>

```
[
    {
        "ID": 2,
        "CreatedAt": "2024-02-11T15:41:24.418245+03:00",
        "UpdatedAt": "2024-02-11T15:41:24.418245+03:00",
        "DeletedAt": null,
        "Name": "Good evening",
        "Author": "Igor Alburov",
        "Publication": "11-02-2024"
    },
    {
        "ID": 3,
        "CreatedAt": "2024-02-11T15:41:35.109116+03:00",
        "UpdatedAt": "2024-02-11T15:41:35.109116+03:00",
        "DeletedAt": null,
        "Name": "Good sunday",
        "Author": "Igor Alburov",
        "Publication": "11-02-2024"
    },
    {
        "ID": 4,
        "CreatedAt": "2024-02-11T15:41:45.626289+03:00",
        "UpdatedAt": "2024-02-11T15:41:45.626289+03:00",
        "DeletedAt": null,
        "Name": "Golang cool",
        "Author": "Igor Alburov",
        "Publication": "11-02-2024"
    },
    {
        "ID": 5,
        "CreatedAt": "2024-02-11T15:42:00.110999+03:00",
        "UpdatedAt": "2024-02-11T15:45:56.860548+03:00",
        "DeletedAt": null,
        "Name": "Rails is my love",
        "Author": "Igor Alburov",
        "Publication": "Forever"
    },
    {
        "ID": 6,
        "CreatedAt": "2024-02-11T16:34:34.875054+03:00",
        "UpdatedAt": "2024-02-11T16:34:34.875054+03:00",
        "DeletedAt": null,
        "Name": "Rails is fucking great",
        "Author": "Igor Alburov",
        "Publication": "11-02-2024"
    }
]

```

<h4>Request</h4>
```
  POST localhost:9010/books
  BODY = {
    "name": "Rails is fucking great",
    "author": "Igor Alburov",
    "publication": "11-02-2024"
  }
```
<h4>Response</h4>
```
{
  "ID": 6,
  "CreatedAt": "2024-02-11T16:34:34.875054203+03:00",
  "UpdatedAt": "2024-02-11T16:34:34.875054203+03:00",
  "DeletedAt": null,
  "Name": "Rails is fucking great",
  "Author": "Igor Alburov",
  "Publication": "11-02-2024"
}
```

<h4>Request</h4>
```
  GET localhost:9010/books/:bookID
```
<h4>Response</h4>
```
{
  "ID": 5,
  "CreatedAt": "2024-02-11T15:42:00.110999+03:00",
  "UpdatedAt": "2024-02-11T15:45:56.860548+03:00",
  "DeletedAt": null,
  "Name": "Rails is my love",
  "Author": "Igor Alburov",
  "Publication": "Forever"
}
```

<h4>Request</h4>
```
  PUT localhost:9010/books/:bookID
  BODY = {
    "name": "Rails is my love",
    "publication": "Forever"
  }
```
<h4>Response</h4>
```
{
  "ID": 7,
  "CreatedAt": "2024-02-11T16:34:48.271007387+03:00",
  "UpdatedAt": "2024-02-11T16:34:48.271007387+03:00",
  "DeletedAt": null,
  "Name": "Rails is my love",
  "Author": "",
  "Publication": "Forever"
}
```

<h4>Request</h4>
```
  DELETE localhost:9010/books/:bookID
```
<h4>Response</h4>
```
{
  "ID": 0,
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z",
  "DeletedAt": null,
  "Name": "",
  "Author": "",
  "Publication": ""
}
```
