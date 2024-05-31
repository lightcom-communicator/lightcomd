# lightcom - access tokens

Each request assigned to your account must contain one of your access tokens which should appear in HTTP headers. Just like this:
```http request
GET /newMessages
Authorization: <access token>
```