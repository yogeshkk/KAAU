package utility

import "io/ioutil"

// UserIsValid export
func UserIsValid(Name, pwd string) bool {
	// DB simulation
	Name, FilePwd, IsValid := "yogesh", "yogesh", false

	if Name == Name && pwd == FilePwd {
		IsValid = true
	} else {
		IsValid = false
	}
	return IsValid
}

// IsEmpty export
func IsEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	}
	return false

}

// LoadFile export
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
