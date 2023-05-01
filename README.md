## Envoy API-Gateway Demo

### Deployment
```bash
docker-compose up --build
```

### Testing HTTP Endpoint
Just Open the OpenAPI docs at http://localhost:8080/v1/docs/

### Testing gRPC Endpoint
1. Get token with /v1/auth/login endpoint
2. Import proto/bookstore/bookstore.proto and proto/cart/cart.proto with postman
3. Choose desired endpoint and go to authorization tab
4. Paste token with **Bearer Token** type