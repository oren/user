# User

## API

### register
```
curl -X POST -d '{"email":"dan@gmail.com", "password":"somesecret"}' localhost:3000/register
```
save in db and return JWT

### login
```
curl -X POST -d '{"email":"dan@gmail.com", "password":"somesecret"}' localhost:3000/login
```
return JWT

### forgot-password
```
curl -X POST -d '{"email":"dan@gmail.com", "password":"somesecret"}' localhost:3000/forgot-password
```
send email with link to a page that allow password change
