# authorization_server

### About
- This service use password grant type
- Jwt token is used as access token
### Start project
- change config in config/config.yaml
  + change dsn to connect to mysql 
- go mod tidy
- go run ./main.go
### Test project
- go to http://localhost:5001/signup to create new user
- go to https://localhost:5001/login to and type username,password
- copy access token and use it to test api_uploader
