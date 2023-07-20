package services

import (
	"fmt"
	"open-messaging/src/config"
	"open-messaging/src/structs"
	"strconv"
	"time"

	"github.com/mypurecloud/platform-client-sdk-go/v77/platformclientv2"
)

func SendMessage(api *platformclientv2.ConversationsApi, conversation_message_id string, messageData *structs.MessageData) {
	gcConfig := config.SetConfig()

	var body platformclientv2.Opennormalizedmessage

	var sid string = ``
	var svarType string = `Private`
	var smessageId string = messageData.SenderMessageId
	var sIntegrationId string = gcConfig.MessageDeployment
	var snickname string = gcConfig.Nickname
	var sidtype string = `email`
	var sfirstname string = gcConfig.FirstName
	var slastname string = gcConfig.LastName
	var sbodytext string = messageData.Body
	var sfromaddress string = gcConfig.Email
	dtime := time.Now()
	var stype string = `Text`
	var sdirection string = `Inbound`
	//var sPlatform string = `Open`

	var openmessageToRecipient platformclientv2.Openmessagingtorecipient
	openmessageToRecipient.Id = &sIntegrationId

	var openmessageFromRecipient platformclientv2.Openmessagingfromrecipient
	openmessageFromRecipient.Id = &sfromaddress
	openmessageFromRecipient.FirstName = &sfirstname
	openmessageFromRecipient.LastName = &slastname
	openmessageFromRecipient.Nickname = &snickname
	openmessageFromRecipient.IdType = &sidtype

	var schannel platformclientv2.Openmessagingchannel
	schannel.VarType = &svarType
	schannel.MessageId = &smessageId
	//schannel.Id = &smessageId
	schannel.To = &openmessageToRecipient
	schannel.From = &openmessageFromRecipient
	schannel.Time = &dtime

	//schannel.VarType = &svarType
	//schannel.Platform = &sPlatform
	//schannel.MessageId = &smessageId

	body.Id = &sid
	body.Channel = &schannel
	body.VarType = &stype
	body.Text = &sbodytext
	body.Direction = &sdirection

	if *body.Text == "" {

	} else {
		messageData.Transcript = append(messageData.Transcript, map[string]string{
			"id":      strconv.Itoa(messageData.ID),
			"sender":  gcConfig.Nickname,
			"message": *body.Text,
			"purpose": "customer",
		})
	}

	messageData.ID++

	fmt.Println("TRANSCRIPT", messageData.Transcript)

	fmt.Println("BODY TEXT", body.Text)
	//normalizedmessage, apiresponse, apierr :=
	responsemessage, apiresponse, apierror := api.PostConversationsMessagesInboundOpen(body)
	fmt.Printf("Response:\n  Success: %v\n  Status code: %v\n  Correlation ID: %v\n", apiresponse.IsSuccess, apiresponse.StatusCode, apiresponse.CorrelationID)
	if apierror != nil {
		fmt.Printf("Error calling Open Messaging: %v\n", apierror)
	} else {
		fmt.Printf("Open Messaging response, %v\n", &responsemessage.Id)
	}
}

func SendFirstMessage(api *platformclientv2.ConversationsApi, messageData *structs.MessageData) {
	gcConfig := config.SetConfig()

	var body platformclientv2.Opennormalizedmessage

	var sid string = ``
	var svarType string = `Private`
	var smessageId string = messageData.SenderMessageId
	var sIntegrationId string = gcConfig.MessageDeployment
	var snickname string = gcConfig.Nickname
	var sidtype string = `email`
	var sfirstname string = gcConfig.FirstName
	var slastname string = gcConfig.LastName
	var sbodytext string = `Hello World`
	var sfromaddress string = gcConfig.Email
	dtime := time.Now()
	var stype string = `Text`
	var sdirection string = `Inbound`
	//var sPlatform string = `Open`

	var openmessageToRecipient platformclientv2.Openmessagingtorecipient
	openmessageToRecipient.Id = &sIntegrationId

	var openmessageFromRecipient platformclientv2.Openmessagingfromrecipient
	openmessageFromRecipient.Id = &sfromaddress
	openmessageFromRecipient.FirstName = &sfirstname
	openmessageFromRecipient.LastName = &slastname
	openmessageFromRecipient.Nickname = &snickname
	openmessageFromRecipient.IdType = &sidtype

	var schannel platformclientv2.Openmessagingchannel
	schannel.VarType = &svarType
	schannel.MessageId = &smessageId
	//schannel.Id = &smessageId
	schannel.To = &openmessageToRecipient
	schannel.From = &openmessageFromRecipient
	schannel.Time = &dtime

	//schannel.VarType = &svarType
	//schannel.Platform = &sPlatform
	//schannel.MessageId = &smessageId

	body.Id = &sid
	body.Channel = &schannel
	body.VarType = &stype
	body.Text = &sbodytext
	body.Direction = &sdirection

	fmt.Println(body)
	//normalizedmessage, apiresponse, apierr :=

	if *body.Text == "" {

	} else {
		messageData.Transcript = append(messageData.Transcript, map[string]string{
			"id":      strconv.Itoa(messageData.ID),
			"sender":  gcConfig.Nickname,
			"message": *body.Text,
			"purpose": "customer",
		})
	}

	messageData.ID++

	fmt.Println("TRANSCRIPT", messageData.Transcript)

	responsemessage, apiresponse, apierror := api.PostConversationsMessagesInboundOpen(body)
	fmt.Printf("Response:\n  Success: %v\n  Status code: %v\n  Correlation ID: %v\n", apiresponse.IsSuccess, apiresponse.StatusCode, apiresponse.CorrelationID)
	if apierror != nil {
		fmt.Printf("Error calling Open Messaging: %v\n", apierror)
	} else {
		fmt.Printf("Open Messaging response, %v\n", &responsemessage.Id)
	}

}
