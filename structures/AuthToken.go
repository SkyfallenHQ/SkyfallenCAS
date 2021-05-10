package structures

import "time"

type AuthToken struct {

	Token string
	TokenExpires time.Time
	TokenCreated time.Time

	TokenScope string
	TokenReason string
	TokenUser string

}
