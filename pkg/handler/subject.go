package handler

import (
	"net/http"
)

func getSubjectFromRequest(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// for this case, we choose to skip the error check because just like HeaderAuthUserField
// in the future, we can extend this function to support cookie jwt token validation
// this assumes that the cookie we get already passed the auth middleware
