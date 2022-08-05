// Package errcode redesigned error code
package errcode

//
// 	Error code design rule:
// 		Pure numbers represent that different parts represent
// 		different services and different modules.
//
//	Example: 10 00 00
//
//			id				details
//		--------------------------------
// 			10			    service
//			00				module
//			00			    id number, each module can register 100 errors
//

var (
	Success                   = NewError(100001, "Success")
	InternalServerError       = NewError(100002, "Internal server error")
	InvalidParams             = NewError(100003, "Invalid parameter")
	NotFound                  = NewError(100004, "Not found")
	DatabaseError             = NewError(100101, "Database error")
	UnauthorizedAuthNotExist  = NewError(100201, "Authorized files not exist")
	UnauthorizedTokenError    = NewError(100202, "Token error")
	UnauthorizedTokenTimeout  = NewError(100203, "Token timeout")
	UnauthorizedTokenGenerate = NewError(100204, "Failed to generate token")
	TooManyRequests           = NewError(100301, "Too many requests")
)
