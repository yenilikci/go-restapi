package utils

func CheckError(err error) {
	if err != nil {
		err.Error()
	}
}
