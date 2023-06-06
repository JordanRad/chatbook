// Code generated by goa v3.11.3, DO NOT EDIT.
//
// chatbook-microservices HTTP client CLI support package
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	authc "github.com/JordanRad/chatbook/services/internal/gen/http/auth/client"
	infoc "github.com/JordanRad/chatbook/services/internal/gen/http/info/client"
	userc "github.com/JordanRad/chatbook/services/internal/gen/http/user/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `auth (refresh-token|login)
info get-info
user (register|get-user-profile|update-profile-names)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` auth refresh-token --body '{
      "email": "Assumenda autem quo aut.",
      "refreshToken": "Reprehenderit totam."
   }'` + "\n" +
		os.Args[0] + ` info get-info` + "\n" +
		os.Args[0] + ` user register --body '{
      "confirmedPassword": "Magnam facilis incidunt occaecati consequatur ullam.",
      "email": "Ducimus asperiores sunt.",
      "firstName": "Qui impedit tempore sunt optio.",
      "lastName": "Quisquam eos vitae velit quis.",
      "password": "Adipisci qui suscipit ut."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		authFlags = flag.NewFlagSet("auth", flag.ContinueOnError)

		authRefreshTokenFlags    = flag.NewFlagSet("refresh-token", flag.ExitOnError)
		authRefreshTokenBodyFlag = authRefreshTokenFlags.String("body", "REQUIRED", "")

		authLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		authLoginBodyFlag = authLoginFlags.String("body", "REQUIRED", "")

		infoFlags = flag.NewFlagSet("info", flag.ContinueOnError)

		infoGetInfoFlags = flag.NewFlagSet("get-info", flag.ExitOnError)

		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		userRegisterBodyFlag = userRegisterFlags.String("body", "REQUIRED", "")

		userGetUserProfileFlags = flag.NewFlagSet("get-user-profile", flag.ExitOnError)

		userUpdateProfileNamesFlags    = flag.NewFlagSet("update-profile-names", flag.ExitOnError)
		userUpdateProfileNamesBodyFlag = userUpdateProfileNamesFlags.String("body", "REQUIRED", "")
	)
	authFlags.Usage = authUsage
	authRefreshTokenFlags.Usage = authRefreshTokenUsage
	authLoginFlags.Usage = authLoginUsage

	infoFlags.Usage = infoUsage
	infoGetInfoFlags.Usage = infoGetInfoUsage

	userFlags.Usage = userUsage
	userRegisterFlags.Usage = userRegisterUsage
	userGetUserProfileFlags.Usage = userGetUserProfileUsage
	userUpdateProfileNamesFlags.Usage = userUpdateProfileNamesUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "auth":
			svcf = authFlags
		case "info":
			svcf = infoFlags
		case "user":
			svcf = userFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "auth":
			switch epn {
			case "refresh-token":
				epf = authRefreshTokenFlags

			case "login":
				epf = authLoginFlags

			}

		case "info":
			switch epn {
			case "get-info":
				epf = infoGetInfoFlags

			}

		case "user":
			switch epn {
			case "register":
				epf = userRegisterFlags

			case "get-user-profile":
				epf = userGetUserProfileFlags

			case "update-profile-names":
				epf = userUpdateProfileNamesFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "auth":
			c := authc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "refresh-token":
				endpoint = c.RefreshToken()
				data, err = authc.BuildRefreshTokenPayload(*authRefreshTokenBodyFlag)
			case "login":
				endpoint = c.Login()
				data, err = authc.BuildLoginPayload(*authLoginBodyFlag)
			}
		case "info":
			c := infoc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-info":
				endpoint = c.GetInfo()
				data = nil
			}
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = userc.BuildRegisterPayload(*userRegisterBodyFlag)
			case "get-user-profile":
				endpoint = c.GetUserProfile()
				data = nil
			case "update-profile-names":
				endpoint = c.UpdateProfileNames()
				data, err = userc.BuildUpdateProfileNamesPayload(*userUpdateProfileNamesBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// authUsage displays the usage of the auth command and its subcommands.
func authUsage() {
	fmt.Fprintf(os.Stderr, `Authentication service is responsible for handling user data and requests
Usage:
    %[1]s [globalflags] auth COMMAND [flags]

COMMAND:
    refresh-token: RefreshToken implements refreshToken.
    login: Login implements login.

Additional help:
    %[1]s auth COMMAND --help
`, os.Args[0])
}
func authRefreshTokenUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth refresh-token -body JSON

RefreshToken implements refreshToken.
    -body JSON: 

Example:
    %[1]s auth refresh-token --body '{
      "email": "Assumenda autem quo aut.",
      "refreshToken": "Reprehenderit totam."
   }'
`, os.Args[0])
}

func authLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth login -body JSON

Login implements login.
    -body JSON: 

Example:
    %[1]s auth login --body '{
      "email": "Provident harum similique saepe fuga assumenda.",
      "password": "Sequi libero et qui soluta error."
   }'
`, os.Args[0])
}

// infoUsage displays the usage of the info command and its subcommands.
func infoUsage() {
	fmt.Fprintf(os.Stderr, `Application info
Usage:
    %[1]s [globalflags] info COMMAND [flags]

COMMAND:
    get-info: GetInfo implements getInfo.

Additional help:
    %[1]s info COMMAND --help
`, os.Args[0])
}
func infoGetInfoUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] info get-info

GetInfo implements getInfo.

Example:
    %[1]s info get-info
`, os.Args[0])
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `User service is responsible for handling user data and requests
Usage:
    %[1]s [globalflags] user COMMAND [flags]

COMMAND:
    register: Register implements register.
    get-user-profile: GetUserProfile implements getUserProfile.
    update-profile-names: UpdateProfileNames implements updateProfileNames.

Additional help:
    %[1]s user COMMAND --help
`, os.Args[0])
}
func userRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user register -body JSON

Register implements register.
    -body JSON: 

Example:
    %[1]s user register --body '{
      "confirmedPassword": "Magnam facilis incidunt occaecati consequatur ullam.",
      "email": "Ducimus asperiores sunt.",
      "firstName": "Qui impedit tempore sunt optio.",
      "lastName": "Quisquam eos vitae velit quis.",
      "password": "Adipisci qui suscipit ut."
   }'
`, os.Args[0])
}

func userGetUserProfileUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user get-user-profile

GetUserProfile implements getUserProfile.

Example:
    %[1]s user get-user-profile
`, os.Args[0])
}

func userUpdateProfileNamesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user update-profile-names -body JSON

UpdateProfileNames implements updateProfileNames.
    -body JSON: 

Example:
    %[1]s user update-profile-names --body '{
      "firstName": "Quibusdam qui.",
      "lastName": "Tempora perspiciatis ut ut."
   }'
`, os.Args[0])
}