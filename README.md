# Open Messaging built using Golang

### To clone via the git CLI
#### Installing Git
Git download URL: https://git-scm.com/downloads

### Installing Golang

#### Download links
##### Windows
https://go.dev/doc/install

#### Clone to PC
**Instructions for Windows**
```
Open powershell.exe as administrator
cd Desktop
git clone https://github.com/m0thm4n/open-messaging-golang.git
cd open-messaging-golang
```

### Configuring the .env file
###### The .env file needs to configured with your own settings
Create a file called ```.env``` inside the root of the open-messaging-golang folder and paste the below into it:
```
clientId = "" # OAuth client ID
clientSecret = "" # OAuth client secret
messageDeployment = "" # Your Open Messaging Integration ID
region = "" # The API region for your org
environment = "" # The region for your test org
nickname = "" # The nickname you want to use
email = "" # The email you want to use
firstName = "" # The first name you want to use
lastName = "" # The last name you want to use
```

#### Building the application
To build the application inside a powershell window run the following command:
```go build```



##### Ngrok for port forwarding

###### https://ngrok.com/

Run ngrok using this command in a powershell window: ```ngrok http 8081```

###### While inside a powershell window run the app using the following command in the directory where you cloned the project

```open-messaging-golang.exe```