package configs

import "time"

var JWTExpirationTime = time.Hour * 72 // 3 days
var JWTExpirationTimeInSeconds = int(JWTExpirationTime.Seconds())

var CORSAllowedOrigins = []string{"https://gossipgo.netlify.app"}
var CORSAllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
var CORSAllowedHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
