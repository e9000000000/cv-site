# cv site
the point of it is to show people my web development skils.

# prepare for runing
- install golang-migrate `go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- create + migrate database `migrate -source file://migrations -database sqlite3://db.sqlite3 up`
- install dependencies `go get`

# run
```bash
# if you wanna add new user / new admin user
go run . -action newuser

# run server
go run . -action server  # server will be ran on 0.0.0.0:8000
```
