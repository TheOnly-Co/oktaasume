package main

import (
	"fmt"
	"github.com/hunkeelin/oktalib"
	"github.com/hunkeelin/userprompt"
	"net/http/cookiejar"
	"os"
	"os/user"
)

func main() {
	o, err := oktalib.New(&oktalib.NewInput{
		Org:                 "dev-815627",
		IdentityProviderArn: "arn:aws:iam::216228501626:saml-provider/Okta_2",
		SamlURI:             "/app/amazon_aws/exkawa67iQIlhKIxE4x6/sso/saml",
	})
	if err != nil {
		panic(err)
	}
	o.Username, o.Password, o.CookieJar = getCredentials()
	err = o.LdapLogin()
	if err != nil {
		panic(err)
	}

	if len(o.UserAuth.Embedded.Factors) == 0 || len(o.UserAuth.Embedded.Factors) < 1 {
		panic(fmt.Errorf("Extra verification must be enabled in Okta. Visit https://varomoney.okta.com/enduser/settings."))
	}

	switch {
	case searchAuthMethod(o.UserAuth.Embedded.Factors, oktalib.YubiKey):
		fmt.Println("Congrats on your shiny new Yubikey")
		code, err := userprompt.UserPrompt("Please give it a squeeze", false)
		if err != nil {
			panic(err)
		}
		err = o.OktaMfa(oktalib.YubiKey, code)
		if err != nil {
			panic(err)
		}
	case searchAuthMethod(o.UserAuth.Embedded.Factors, oktalib.MfaPush):
		err = o.OktaMfa(oktalib.MfaPush, "")
		if err != nil {
			panic(err)
		}
	case searchAuthMethod(o.UserAuth.Embedded.Factors, oktalib.MfaCode):
		passcode, err := userprompt.UserPrompt("Enter a token from your mobile authenticator app", false)
		if err != nil {
			panic(err)
		}
		err = o.OktaMfa(oktalib.MfaCode, passcode)
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Errorf("No recognized mfa method found please contact your administrator."))
	}
	out, err := o.GetAwsCredentials(oktalib.GetAwsCredentialsInput{
		RoleArn:    "arn:aws:iam::216228501626:role/devops-admin-role",
		Expiration: 28800,
	})
	if err != nil {
		panic(err)
	}
	id := []byte("aws_access_key_id = " + out.AwsAccessKeyId + " \n")
	key := []byte("aws_secret_access_key = " + out.AwsSecretAccessKey + " \n")
	token := []byte("aws_session_token = " + out.AwsSessionToken + " \n")
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	towrite := []byte("[default] \n")
	towrite = append(towrite, id...)
	towrite = append(towrite, key...)
	towrite = append(towrite, token...)
	if err != nil {
		panic(err)
	}

	err = writeCredToFile(&writeCredToFileInput{
		location: user.HomeDir + "/.aws/credentials",
		towrite:  towrite,
	})

	if err != nil {
		panic(err)
	}
	return
}

type writeCredToFileInput struct {
	location string
	towrite  []byte
}

func writeCredToFile(i *writeCredToFileInput) error {
	f, err := os.Create(i.location)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(i.towrite)
	if err != nil {
		panic(err)
	}
	return nil
}

func searchAuthMethod(sep []oktalib.OktaUserAuthnFactor, s string) bool {
	for _, i := range sep {
		if i.FactorType == s {
			return true
		}
	}
	return false

}

func getCredentials() (string, string, *cookiejar.Jar) {
	currentUser := os.Getenv("USER")
	userName, err := userprompt.UserPromptWithDefault("Enter Okta Username ("+currentUser+")", currentUser, false)
	if err != nil {
		panic(err)
	}
	pass, err := userprompt.UserPrompt("Enter Okta Password", true)
	if err != nil {
		panic(err)
	}
	cJar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return userName, pass, cJar
}
