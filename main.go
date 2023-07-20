package main

import (
	"open-messaging/src/config"
	"open-messaging/src/controller"
	"open-messaging/src/services"

	"html/template"
	"log"
	"net/http"

	"open-messaging/src/structs"

	"github.com/google/uuid"
	"github.com/mypurecloud/platform-client-sdk-go/v77/platformclientv2"
)

func GetConfig() *platformclientv2.ConversationsApi {
	// Create new config
	gcConfig := config.SetConfig()
	config := platformclientv2.NewConfiguration()
	config.BasePath = platformclientv2.USWest2
	// config.AccessToken = "Wm0Bzx-PZbJx4T8uZX0FEp0Q4RpxFlncfE4XzOdVnyEffg-pVZEfe-ollZatP5au-WfIRWNHPnOoPDqdFv_qug"

	// err := config.AuthorizeClientCredentials(os.Getenv("GENESYS_CLOUD_CLIENT_ID"), os.Getenv("GENESYS_CLOUD_CLIENT_SECRET"))
	// if err != nil {
	//     panic(err)
	// }

	config.LoggingConfiguration.LogRequestBody = true
	config.LoggingConfiguration.LogResponseBody = true
	config.LoggingConfiguration.LogLevel = platformclientv2.LTrace
	config.LoggingConfiguration.SetLogFormat(platformclientv2.JSON)
	config.LoggingConfiguration.SetLogToConsole(true)
	config.LoggingConfiguration.SetLogFilePath("golangsdk.log")

	err := config.AuthorizeClientCredentials(gcConfig.ClientID, gcConfig.ClientSecret)
	if err != nil {
		panic(err)
	}

	messagingApi := platformclientv2.NewConversationsApiWithConfig(config)

	return messagingApi

	// Create API instance using config
	// usersAPI := platformclientv2.NewUsersApiWithConfig(config)
}

func main() {
	guid := uuid.New()

	messageData := structs.MessageData{SenderMessageId: guid.String(), NumberOfMessages: 10, OpenMessageApi: GetConfig(), Transcript: []map[string]string{}, ID: 1}

	messageData.NumberOfMessages = 10

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	tmpl := template.Must(template.ParseFiles("./src/html/index.gohtml"))

	log.Println("server started")
	services.SendFirstMessage(messageData.OpenMessageApi, &messageData)
	http.HandleFunc("/messageToGenesys", func(w http.ResponseWriter, r *http.Request) {
		controller.HandleSendMessage(w, r, &messageData, tmpl)
	})
	http.HandleFunc("/messageFromGenesys", func(w http.ResponseWriter, r *http.Request) {
		controller.HandleFromMessage(w, r, &messageData)
	})
	http.HandleFunc("/transcript", func(w http.ResponseWriter, r *http.Request) {
		controller.HandleTranscript(w, r, &messageData)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// func SendFirstMessage(api *platformclientv2.ConversationsApi) {
// 	var body platformclientv2.Opennormalizedmessage

// 	var sid string = ``
// 	var svarType string = `Private`
// 	var smessageId string = sender_message_id
// 	var sIntegrationId string = `ac9d8d2c-847f-43bd-84e9-71cc6a5cd68d`
// 	var snickname string = `Paddington`
// 	var sidtype string = `email`
// 	var sfirstname string = `Paddington`
// 	var slastname string = `Bear`
// 	var sbodytext string = `Hello World`
// 	var sfromaddress string = `paddington@example.com`
// 	dtime := time.Now()
// 	var stype string = `Text`
// 	var sdirection string = `Inbound`
// 	//var sPlatform string = `Open`

// 	var openmessageToRecipient platformclientv2.Openmessagingtorecipient
// 	openmessageToRecipient.Id = &sIntegrationId

// 	var openmessageFromRecipient platformclientv2.Openmessagingfromrecipient
// 	openmessageFromRecipient.Id = &sfromaddress
// 	openmessageFromRecipient.FirstName = &sfirstname
// 	openmessageFromRecipient.LastName = &slastname
// 	openmessageFromRecipient.Nickname = &snickname
// 	openmessageFromRecipient.IdType = &sidtype

// 	var schannel platformclientv2.Openmessagingchannel
// 	schannel.VarType = &svarType
// 	schannel.MessageId = &smessageId
// 	//schannel.Id = &smessageId
// 	schannel.To = &openmessageToRecipient
// 	schannel.From = &openmessageFromRecipient
// 	schannel.Time = &dtime

// 	//schannel.VarType = &svarType
// 	//schannel.Platform = &sPlatform
// 	//schannel.MessageId = &smessageId

// 	body.Id = &sid
// 	body.Channel = &schannel
// 	body.VarType = &stype
// 	body.Text = &sbodytext
// 	body.Direction = &sdirection

// 	fmt.Println(body)
// 	//normalizedmessage, apiresponse, apierr :=

// 	responsemessage, apiresponse, apierror := api.PostConversationsMessagesInboundOpen(body)
// 	fmt.Printf("Response:\n  Success: %v\n  Status code: %v\n  Correlation ID: %v\n", apiresponse.IsSuccess, apiresponse.StatusCode, apiresponse.CorrelationID)
// 	if apierror != nil {
// 		fmt.Printf("Error calling Open Messaging: %v\n", apierror)
// 	} else {
// 		fmt.Printf("Open Messaging response, %v\n", &responsemessage.Id)
// 	}

// }
