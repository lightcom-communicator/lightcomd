# lightcom - message protocol

On your lightcom server you can use whatever message protocol you want, but we recommend using this.

## Messages encryption
Every message should be encrypted using AES-GCM and as a key should be shared secret calculated from your private key and destination user's public key.

## Unencrypted message structure
The structure of the message should be:
```json
{
  "content": "<text of the message>",
  "timestamp": <when message was sent in unix format>,
  "mediaUrl": []
}
```
Sending medias will be implemented later, so for now set `mediaUrls` field to an empty array.