
# RestGo.Mud
RestGo.MUD.Core is a package for creating a Multi-User Dungeon (MUD) game server. It provides the necessary tools to create and manage a MUD game, including user commands, objects, and rooms.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

 1. You will need to have Go installed on your machine in order to run
 2. Data storage is Google firestore

RestGo.Mud.


### Installing
To install RestGo.Mud, clone this repository onto your local machine:
```
git clone https://github.com/myrest/RestGo.Mud.git
```

|Folder |File Name  |
|--|--|
| / | ServerConfig.json |
| /Content | Welcome.txt |

##### ServerConfig.json

```
{
    "ListenOnPortNumber": 4444,
    "Firebase": {
        "ProjectID": "myprojectid",
        "CredentialFile": "CredentialFile-f333d9c85c96.json"
    },
    "MaxLength": 38
}
```
  It is configured to listen on port number 4444. It uses Firebase as its database, with the ProjectID "myprojectid" and the CredentialFile "CredentialFile-f333d9c85c96.json". The maximum length of data per line that can be stored in the database is 38 characters.

##### Word document
Folder structure
/Documents/Object
/Documents/Rooms/[World]/[Region]/[Area]/[ROOMS].json


Then, navigate into the RestGo.Mud directory and run the following command:

 ```
 go build main.go
 ```

 This will compile the code into an executable file that can be run on your machine.

 ## Running RestGo.Mud

 To start the server, run the following command:

 ```
 ./main -port <port_number>
 ```

 Where <port_number> is the port number you want to use for the server (default is 4444). Once the server is running, users can connect to it using any telnet client (such as PuTTY).

 ## Built With

 * [Go](https://golang.org/) - The programming language used
 * [Firestore](https://firebase.google.com/products/firestore/) - The programming Data storage.