package utils

func InvalidRequest() map[string]interface{} {
	return Message("Invalid request")
}

func ErrorOccured() map[string]interface{} {
	return Message("Error occured")
}

func NotFound() map[string]interface{} {
	return Message("Not found")
}

func Created() map[string]interface{} {
	return Message("Created")
}

func Updated() map[string]interface{} {
	return Message("Updated")
}

func Found() map[string]interface{} {
	return Message("Found")
}

func Deleted() map[string]interface{} {
	return Message("Deleted")
}

func NoCookie() map[string]interface{} {
	return Message("No cookie")
}

func CantDelete() map[string]interface{} {
	return Message("Can not delete")
}
