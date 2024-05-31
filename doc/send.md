# lightcom - sending messages

You need to know the destination's user id.
```http request
PUT /send
Content-Type: application/json
Authorization: <access token>

{
  "fromUser": "<no matter what, the server will use user id from access token>",
  "toUser": "<destination's user id>",
  "content": "<message, should be encrypted and hex-encoded>"
}
```
On success server will return 201 and:
```json
{
  "sent": true
}
```