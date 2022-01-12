# GO URL SHORTENER API
A simple url shortener built with GO and Redis

### Endpoint
``POST {{URL}}/create-short-url``

#### Payload
Accepts long url and user id to distinguish urls

```json
{
    "long_url":https://www.google.com",
    "user_id":"abx-qwe-345-ddd"
}
```

#### Result
```json
{
    "message": "short url created successfully",
    "short_url": "http://localhost:9808/2dDEQAS1"
}
```

### Endpoint
``GET {{URL}}/:short-url:``

#### Result
Redirects to initial URL(Long URL)

## RUN TEST
Run this command in terminal `` go test ./... -v``
