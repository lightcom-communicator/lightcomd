# lightcom - receiving messages

## 1. Get who and how many messages sent you
### Regular request
```http request
GET /newMessages
Authorization: <access token>
```
Response should look like this
```json
{
  "<user id>": <how many messages>,
  "<user id>": <how many messages>,
  ...
}
```
Only users which sent us messages appear there.

### Websocket
If we want connection with the server, which will be informing us when someone send us message, there is need to use websocket.
```http request
WEBSOCKET /newMessagesWS

<from us>
{
  "accessToken": "<your access token>"
}

<first from server>
{
  "<user id>": <how many>,
  ...
}

<rest is from server>
{
  "<user id>": 1
}
...
```

## 2. Fetching messages
To fetch messages from a specified user:
```http request
GET /fetch/<user id>
Authorization: <access token>
```
Response should look like this
```json
[
  {
    "fromUser": "<user id>",
    "toUser": "<your user id>",
    "content": "<content, should be encrypted and hex-encoded>"
  },
  ...
]
```