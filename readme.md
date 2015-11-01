[![Stories in Ready](https://badge.waffle.io/oren/user.png?label=ready&title=Ready)](https://waffle.io/oren/user)
# User

## API

### register
```
curl -X POST -H 'Content-Type: application/json' -d '{"email":"dan@gmail.com", "password":"somesecret"}' localhost:3000/user
```
save in db and return JWT

### delete user
```
curl -X DELETE -H 'Content-Type: application/json' -d '{"email":"dan@gmail.com", "password":"somesecret"}' localhost:3000/user
```

### login
```
curl -X POST -H 'Content-Type: application/json' -d '{"email":"dan@gmail.com", "password":"somesecret"}' localhost:3000/login
```
return JWT

### forgot password
```
curl -X POST -H 'Content-Type: application/json' -d '{"email":"dan@gmail.com", "password":"somesecret"}' localhost:3000/forgot-password
```
send email with link to a page that allow password change

### change password
```
curl -X PUT -H 'Content-Type: application/json' -d '{"email":"dan@gmail.com", "password":"somesecret", "new-password":"somenewsecret"}' localhost:3000/change-password
```

