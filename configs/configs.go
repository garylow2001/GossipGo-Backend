package configs

import "time"

var JWTExpirationTime = time.Hour * 72 // 3 days
var JWTExpirationTimeInSeconds = int(JWTExpirationTime.Seconds())
