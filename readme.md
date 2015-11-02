# User

[![Stories in Ready](https://badge.waffle.io/oren/user.png?label=ready&title=Ready)](https://waffle.io/oren/user)

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

## Development

### Auto Reload Script

The run.sh script located on the root directory of the repository will allow you to use CompileDaemon to auto compile and reload the web application as soon as you make changes. This allows for faster development. To get started, you need to get the CompileDaemon tool first.

    go get -u github.com/githubnemo/CompileDaemon

Then you can run the run.sh script and it will auto detect changes, compile the app and run it.

    ./run.sh
