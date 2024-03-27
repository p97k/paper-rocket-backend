package user

import (
	"net/mail"
)

func HandleSignUpError(userReqData *CreateUserReq) (string, bool) {
	if len(userReqData.Username) < 5 {
		if len(userReqData.Username) < 1 {
			return "please enter a username!", false
		} else {
			return "username must be at least 5 characters", false
		}
	}

	if len(userReqData.Password) < 8 {
		return "password must be at least 8 characters", false
	}

	_, err := mail.ParseAddress(userReqData.Email)
	if err != nil {
		return "email format is not valid", false
	}

	return "", true
}
