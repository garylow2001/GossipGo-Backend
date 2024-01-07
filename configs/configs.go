package configs

import "time"

var JWTExpirationTime = time.Hour * 72 // 3 days
var JWTExpirationTimeInSeconds = int(JWTExpirationTime.Seconds())

var CORSAllowedOrigins = []string{"http://localhost:3001/*"}
var CORSAllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
var CORSAllowedHeaders = []string{"*"}
