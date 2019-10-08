package tekmor

import (
	"errors"
	"github.com/msteinert/pam"
	"os/user"
)

//Identity is the core component for Tekmor used for authentication against PAM
type Identity struct {
	username string
	password string
	details  UserDetails
}

//UserDetails is a struct containing Unix Shell details for an account that has PAM Authenticated
type UserDetails struct {
	username string
	home     string
	group    string
}

//Authenticate handles determining whether or not an Identity is valid per PAM
func (i Identity) Authenticate() (UserDetails, error) {
	t, err := pam.StartFunc("", "", func(s pam.Style, msg string) (string, error) {
		switch s {
		//PAM will request the password with echo off, so we will return the auth Password here.
		case pam.PromptEchoOff:
			return i.password, nil
		//PAM will request the username with prompt on. So we'll return the username here
		case pam.PromptEchoOn:
			return i.username, nil
		//If we receive an error message back from PAM, let's generate an error and return it.
		case pam.ErrorMsg:
			return "", errors.New(msg)
		//For Text Info, we don't really need to return a response to PAM, or anything. Just let it vanish into the ether
		case pam.TextInfo:
			return "", nil
		}
		return "", errors.New("Unknown message style")
	})
	if err != nil {
		return UserDetails{}, err
	}
	err = t.Authenticate(0)
	if err != nil {
		return UserDetails{}, err
	}

	//Successfully authenticated with PAM Let's pull the user details into the struct for use in tokens
	u, err := user.Lookup(i.username)
	if err != nil {
		return UserDetails{}, err
	}

	return UserDetails{username: u.Username, home: u.HomeDir, group: u.Gid}, nil
}