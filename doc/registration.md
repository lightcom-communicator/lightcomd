# lightcom - registering account

## 1. Obtain server's public key
```http request
GET /publicKey
```

## 2. Key generation
Generate keypair using Curve25519 (X25519) and calculate shared secret between your private and server's public.

## 3. Create account
Send your public key to the server, then server should return random user id which will be assigned to this public key.
```http request
PUT /register
Content-Type: application/json

{
  "publicKey": "<your public key hex-encoded>"
}
```