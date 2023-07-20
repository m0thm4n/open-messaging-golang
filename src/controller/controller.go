package controller

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"open-messaging/src/services"
	"open-messaging/src/structs"
	"strconv"

	// "time"

	"rsc.io/quote"
)

func HandleFromMessage(w http.ResponseWriter, r *http.Request, messageData *structs.MessageData) {
	//read request body buffer to byte object until buffer is empty
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	//convert byte object to string, required to calculate signature hash
	bodyString := string(bodyBytes)
	fmt.Println(`webhook received body string `, bodyString)

	//extract received message signature hash from header
	signature := r.Header.Get("x-hub-signature-256")

	//webhook secret token as set in open message integration
	secretToken := "1234"

	//create keyed sha256 hash of secret token
	h := hmac.New(sha256.New, []byte(secretToken))

	//add the json body to the keyed has
	h.Write([]byte(bodyString))

	//convert signature hash to string and prepend sha256=
	calculatedsignature := "sha256=" + base64.StdEncoding.EncodeToString(h.Sum(nil))

	//debug output
	fmt.Println("calculated signature: ", calculatedsignature)
	fmt.Println("received signature:   ", signature)

	//validate signatures match => trusted sender
	if calculatedsignature != signature {
		http.Error(w, "invalid signature", http.StatusInternalServerError)
	}

	var messagingresponse structs.WebhookResponse
	mrtojsonerror := json.Unmarshal(bodyBytes, &messagingresponse)
	if mrtojsonerror != nil {
		fmt.Println("ERROR", mrtojsonerror)
	}

	if messagingresponse.Text == "" {

	} else {
		messageData.Transcript = append(messageData.Transcript, map[string]string{
			"id":      strconv.Itoa(messageData.ID),
			"sender":  "Agent #1",
			"message": messagingresponse.Text,
			"purpose": "customer",
		})
	}

	messageData.ID++

	fmt.Println("TRANSCRIPT", messageData.Transcript)

	// if r.Method != http.MethodPost {
	// 	tmpl.Execute(w, nil)
	// 	return
	// }

	// messageData.Body = r.FormValue("chat-form")

	// tmpl.Execute(w, struct{ Success bool }{true})

	//GC returns type==receipt as async event to confirm the public api call is processed and type==text for a text message from the agent.
	// if messagingresponse.Type == "Text" {
	// 	if messageData.NumberOfMessages > 0 {
	// 		conversationId := messagingresponse.Id
	// 		fmt.Println("messaging conversation id ", conversationId)
	// 		time.Sleep(30 * time.Second)
	// 		fmt.Println("sending new message back to GC ", quote.Glass())
	// 		services.SendMessage(messageData.OpenMessageApi, quote.Glass(), messageData)
	// 		messageData.NumberOfMessages = messageData.NumberOfMessages - 1
	// 		fmt.Println("number of simulation iterations left ", messageData.NumberOfMessages)
	// 	}
	// 	if messageData.NumberOfMessages == 0 {
	// 		fmt.Println("simulation FINISHED")

	// 	}

	// }

}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "src/html/index.html")
}

func HandleSendMessage(w http.ResponseWriter, r *http.Request, messageData *structs.MessageData, tmpl *template.Template) {
	switch r.Method {
	case "GET":
		// http.ServeFile(w, r, "src/html/index.gohtml")
		if err := tmpl.Execute(w, messageData.Transcript); err != nil {
			fmt.Println("FAILED.")
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		messageData.Body = r.FormValue("message")
		fmt.Fprintf(w, "body = %s\n", messageData.Body)

		services.SendMessage(messageData.OpenMessageApi, quote.Glass(), messageData)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func HandleTranscript(w http.ResponseWriter, r *http.Request, messageData *structs.MessageData) {
	// for i := 0; i < len(messageData.Transcript); i++ {
	// 	jsonStr, err := json.Marshal(messageData.Transcript[i])
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	w.Write(jsonStr)
	// }
	jsonStr, err := json.Marshal(messageData.Transcript)
	if err != nil {
		log.Fatalln(err)
	}

	messageData.Transcript = []map[string]string{}

	w.Write(jsonStr)
}
