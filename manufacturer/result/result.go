package result

const success int = 0
const fail int = 1
const timeout int = 2
const noInstance int = 3
const authFail int = 4
const invalidAuthToken int = 5
const invalidManufacturerName int = 6
const SUCCESS string = "success"
const FAIL string = "fail"
const TIMEOUT string = "timeout"
const NO_INSTANCE string = "has no manufacturer instance"
const AUTH_FAIL string = "auth fail"
const INVALID_AUTH_TOKEN string = "Invalid Auth Token"
const INVALID_MANUFACTURER_NAME string = "invalid manufacturer name"

type Result struct {
	Code    int
	Message string
	Data    string
}

// type manufacturer_service interface {
// 	uploadIcon(file string) Result
// 	uploadPic(file string) Result
// }

func Success(data string) Result {
	return Result{success, SUCCESS, data}
}

func Fail(data string) Result {
	return Result{fail, FAIL, data}
}

func Timeout() Result {
	return Result{timeout, TIMEOUT, ""}
}

func NoInstance() Result {
	return Result{noInstance, NO_INSTANCE, ""}
}

func AuthFail() Result {
	return Result{authFail, AUTH_FAIL, ""}
}

func InvalidAuthToken() Result {
	return Result{invalidAuthToken, INVALID_AUTH_TOKEN, ""}
}

func InvalidManufacturerName() Result {
	return Result{invalidManufacturerName, INVALID_MANUFACTURER_NAME, ""}
}
