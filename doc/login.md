# lightcom - logging in

Prepare your user id and shared secret. Then:
```http request
POST /login
Content-Type: application/json

{
  "userId": "<your user id>",
  "sharedSecret": "<your shared secret hex-encoded>"
}
```
Response should have access token and unix date when it expires.
```json
{
  "accessToken": "<access token>",
  "validUntil": <date in unix format>
}
```