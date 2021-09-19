## Install Packages

Run this to install packages first
```bash
go mod tidy
```


## Environment Variable
Fill the empty string with yours
```
ACCOUNT_DATABASE_HOST="127.0.0.1"
ACCOUNT_DATABASE_PORT="3306"
ACCOUNT_DATABASE_USER=""
ACCOUNT_DATABASE_PASSWORD=""
ACCOUNT_DATABASE_NAME=""
SECRET_KEY_JWT=""
ACCOUNT_PORT=":8089"
```
## How to run project
To the root project directory

```bash
go run main.go local
```


## How to run test file
Go to directory where the test file exist, and run on terminal:

```bash
go test
```

