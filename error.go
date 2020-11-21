package helpers

var errorVal bool
var errorMsg string

// StartErrorCheck: Remove anything from the check
func StartErrorCheck() {
	errorVal = false
	errorMsg = ""
}

// CheckError: If error is true will save the message
func CheckError(e bool, msg string) {
	if e {
		errorVal = true
		if !EmptyString(errorMsg) {
			errorMsg += "\n"
		}
		errorMsg += msg
	}
}

// HasError: If error is true, Return true and all the error messages generate
func HasError() (bool, string) {
	if errorVal {
		return true, errorMsg
	}

	return false, ""
}
