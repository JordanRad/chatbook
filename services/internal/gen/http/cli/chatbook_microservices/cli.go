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
	chatc "github.com/JordanRad/chatbook/services/internal/gen/http/chat/client"
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
chat (get-conversation-history|search-in-conversation|get-conversations-list|add-conversation)
user (register|get-profile|update-profile-names|add-friend|remove-friend)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` auth refresh-token --body '{
      "email": "Natus quisquam.",
      "refreshToken": "Illo omnis nesciunt minus sed."
   }'` + "\n" +
		os.Args[0] + ` info get-info` + "\n" +
		os.Args[0] + ` chat get-conversation-history --id "Laboriosam tempore atque mollitia ut." --limit "Eum ullam eveniet temporibus quis dolore mollitia." --before-timestamp 5540675296674446940` + "\n" +
		os.Args[0] + ` user register --body '{
      "confirmedPassword": "Molestiae similique omnis voluptate pariatur non.",
      "email": "Quidem sapiente ex et sunt earum.",
      "firstName": "Aut facere molestiae cumque quia blanditiis quos.",
      "lastName": "Amet quia vero illum.",
      "password": "Enim vel sapiente."
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

		chatFlags = flag.NewFlagSet("chat", flag.ContinueOnError)

		chatGetConversationHistoryFlags               = flag.NewFlagSet("get-conversation-history", flag.ExitOnError)
		chatGetConversationHistoryIDFlag              = chatGetConversationHistoryFlags.String("id", "REQUIRED", "Conversation ID")
		chatGetConversationHistoryLimitFlag           = chatGetConversationHistoryFlags.String("limit", "200", "")
		chatGetConversationHistoryBeforeTimestampFlag = chatGetConversationHistoryFlags.String("before-timestamp", "1257894000", "")

		chatSearchInConversationFlags           = flag.NewFlagSet("search-in-conversation", flag.ExitOnError)
		chatSearchInConversationIDFlag          = chatSearchInConversationFlags.String("id", "REQUIRED", "Conversation ID")
		chatSearchInConversationLimitFlag       = chatSearchInConversationFlags.String("limit", "200", "")
		chatSearchInConversationSearchInputFlag = chatSearchInConversationFlags.String("search-input", "200", "")

		chatGetConversationsListFlags     = flag.NewFlagSet("get-conversations-list", flag.ExitOnError)
		chatGetConversationsListBodyFlag  = chatGetConversationsListFlags.String("body", "REQUIRED", "")
		chatGetConversationsListLimitFlag = chatGetConversationsListFlags.String("limit", "100", "")

		chatAddConversationFlags    = flag.NewFlagSet("add-conversation", flag.ExitOnError)
		chatAddConversationBodyFlag = chatAddConversationFlags.String("body", "REQUIRED", "")

		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		userRegisterBodyFlag = userRegisterFlags.String("body", "REQUIRED", "")

		userGetProfileFlags = flag.NewFlagSet("get-profile", flag.ExitOnError)

		userUpdateProfileNamesFlags    = flag.NewFlagSet("update-profile-names", flag.ExitOnError)
		userUpdateProfileNamesBodyFlag = userUpdateProfileNamesFlags.String("body", "REQUIRED", "")

		userAddFriendFlags  = flag.NewFlagSet("add-friend", flag.ExitOnError)
		userAddFriendIDFlag = userAddFriendFlags.String("id", "REQUIRED", "User ID to add")

		userRemoveFriendFlags  = flag.NewFlagSet("remove-friend", flag.ExitOnError)
		userRemoveFriendIDFlag = userRemoveFriendFlags.String("id", "REQUIRED", "User ID to delete")
	)
	authFlags.Usage = authUsage
	authRefreshTokenFlags.Usage = authRefreshTokenUsage
	authLoginFlags.Usage = authLoginUsage

	infoFlags.Usage = infoUsage
	infoGetInfoFlags.Usage = infoGetInfoUsage

	chatFlags.Usage = chatUsage
	chatGetConversationHistoryFlags.Usage = chatGetConversationHistoryUsage
	chatSearchInConversationFlags.Usage = chatSearchInConversationUsage
	chatGetConversationsListFlags.Usage = chatGetConversationsListUsage
	chatAddConversationFlags.Usage = chatAddConversationUsage

	userFlags.Usage = userUsage
	userRegisterFlags.Usage = userRegisterUsage
	userGetProfileFlags.Usage = userGetProfileUsage
	userUpdateProfileNamesFlags.Usage = userUpdateProfileNamesUsage
	userAddFriendFlags.Usage = userAddFriendUsage
	userRemoveFriendFlags.Usage = userRemoveFriendUsage

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
		case "chat":
			svcf = chatFlags
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

		case "chat":
			switch epn {
			case "get-conversation-history":
				epf = chatGetConversationHistoryFlags

			case "search-in-conversation":
				epf = chatSearchInConversationFlags

			case "get-conversations-list":
				epf = chatGetConversationsListFlags

			case "add-conversation":
				epf = chatAddConversationFlags

			}

		case "user":
			switch epn {
			case "register":
				epf = userRegisterFlags

			case "get-profile":
				epf = userGetProfileFlags

			case "update-profile-names":
				epf = userUpdateProfileNamesFlags

			case "add-friend":
				epf = userAddFriendFlags

			case "remove-friend":
				epf = userRemoveFriendFlags

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
		case "chat":
			c := chatc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-conversation-history":
				endpoint = c.GetConversationHistory()
				data, err = chatc.BuildGetConversationHistoryPayload(*chatGetConversationHistoryIDFlag, *chatGetConversationHistoryLimitFlag, *chatGetConversationHistoryBeforeTimestampFlag)
			case "search-in-conversation":
				endpoint = c.SearchInConversation()
				data, err = chatc.BuildSearchInConversationPayload(*chatSearchInConversationIDFlag, *chatSearchInConversationLimitFlag, *chatSearchInConversationSearchInputFlag)
			case "get-conversations-list":
				endpoint = c.GetConversationsList()
				data, err = chatc.BuildGetConversationsListPayload(*chatGetConversationsListBodyFlag, *chatGetConversationsListLimitFlag)
			case "add-conversation":
				endpoint = c.AddConversation()
				data, err = chatc.BuildAddConversationPayload(*chatAddConversationBodyFlag)
			}
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = userc.BuildRegisterPayload(*userRegisterBodyFlag)
			case "get-profile":
				endpoint = c.GetProfile()
				data = nil
			case "update-profile-names":
				endpoint = c.UpdateProfileNames()
				data, err = userc.BuildUpdateProfileNamesPayload(*userUpdateProfileNamesBodyFlag)
			case "add-friend":
				endpoint = c.AddFriend()
				data, err = userc.BuildAddFriendPayload(*userAddFriendIDFlag)
			case "remove-friend":
				endpoint = c.RemoveFriend()
				data, err = userc.BuildRemoveFriendPayload(*userRemoveFriendIDFlag)
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
      "email": "Natus quisquam.",
      "refreshToken": "Illo omnis nesciunt minus sed."
   }'
`, os.Args[0])
}

func authLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth login -body JSON

Login implements login.
    -body JSON: 

Example:
    %[1]s auth login --body '{
      "email": "Labore tempora.",
      "password": "Ut ut."
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

// chatUsage displays the usage of the chat command and its subcommands.
func chatUsage() {
	fmt.Fprintf(os.Stderr, `User service is responsible for handling user data and requests
Usage:
    %[1]s [globalflags] chat COMMAND [flags]

COMMAND:
    get-conversation-history: GetConversationHistory implements getConversationHistory.
    search-in-conversation: SearchInConversation implements searchInConversation.
    get-conversations-list: GetConversationsList implements getConversationsList.
    add-conversation: AddConversation implements addConversation.

Additional help:
    %[1]s chat COMMAND --help
`, os.Args[0])
}
func chatGetConversationHistoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat get-conversation-history -id STRING -limit STRING -before-timestamp INT64

GetConversationHistory implements getConversationHistory.
    -id STRING: Conversation ID
    -limit STRING: 
    -before-timestamp INT64: 

Example:
    %[1]s chat get-conversation-history --id "Laboriosam tempore atque mollitia ut." --limit "Eum ullam eveniet temporibus quis dolore mollitia." --before-timestamp 5540675296674446940
`, os.Args[0])
}

func chatSearchInConversationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat search-in-conversation -id STRING -limit STRING -search-input STRING

SearchInConversation implements searchInConversation.
    -id STRING: Conversation ID
    -limit STRING: 
    -search-input STRING: 

Example:
    %[1]s chat search-in-conversation --id "Animi et velit illo." --limit "Eos magni officia." --search-input "ilz"
`, os.Args[0])
}

func chatGetConversationsListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat get-conversations-list -body JSON -limit STRING

GetConversationsList implements getConversationsList.
    -body JSON: 
    -limit STRING: 

Example:
    %[1]s chat get-conversations-list --body '{
      "ID": "Suscipit qui nesciunt consequatur quia repellat."
   }' --limit "Quia excepturi error similique."
`, os.Args[0])
}

func chatAddConversationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat add-conversation -body JSON

AddConversation implements addConversation.
    -body JSON: 

Example:
    %[1]s chat add-conversation --body '{
      "participants": [
         {
            "email": "Sint quasi et.",
            "firstName": "Consequatur voluptatem.",
            "id": "Ab est.",
            "lastName": "Libero inventore in tempore."
         },
         {
            "email": "Sint quasi et.",
            "firstName": "Consequatur voluptatem.",
            "id": "Ab est.",
            "lastName": "Libero inventore in tempore."
         }
      ]
   }'
`, os.Args[0])
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `User service is responsible for handling user data and requests
Usage:
    %[1]s [globalflags] user COMMAND [flags]

COMMAND:
    register: Register implements register.
    get-profile: GetProfile implements getProfile.
    update-profile-names: UpdateProfileNames implements updateProfileNames.
    add-friend: AddFriend implements addFriend.
    remove-friend: RemoveFriend implements removeFriend.

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
      "confirmedPassword": "Molestiae similique omnis voluptate pariatur non.",
      "email": "Quidem sapiente ex et sunt earum.",
      "firstName": "Aut facere molestiae cumque quia blanditiis quos.",
      "lastName": "Amet quia vero illum.",
      "password": "Enim vel sapiente."
   }'
`, os.Args[0])
}

func userGetProfileUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user get-profile

GetProfile implements getProfile.

Example:
    %[1]s user get-profile
`, os.Args[0])
}

func userUpdateProfileNamesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user update-profile-names -body JSON

UpdateProfileNames implements updateProfileNames.
    -body JSON: 

Example:
    %[1]s user update-profile-names --body '{
      "firstName": "Animi aperiam veniam.",
      "lastName": "Quam ad voluptatem dolor quae accusamus deleniti."
   }'
`, os.Args[0])
}

func userAddFriendUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user add-friend -id STRING

AddFriend implements addFriend.
    -id STRING: User ID to add

Example:
    %[1]s user add-friend --id "Perspiciatis quia."
`, os.Args[0])
}

func userRemoveFriendUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] user remove-friend -id STRING

RemoveFriend implements removeFriend.
    -id STRING: User ID to delete

Example:
    %[1]s user remove-friend --id "Velit enim modi consequatur."
`, os.Args[0])
}
