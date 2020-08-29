module comm.api

go 1.14

require (
	comm.pkgs v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.0.0-beta.8 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/go-sql-driver/mysql v1.5.0
	github.com/spf13/viper v1.7.1 // indirect
)

replace comm.pkgs => ../comm.pkgs
