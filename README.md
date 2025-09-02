# GoChat

GoChat is a simple command-line chat application written in Go. It allows two users to communicate over a local network or the internet. The application supports both text messaging and file transfers.

## Features

- **Real-time Messaging**: Send and receive messages in real-time.
- **File Transfer**: Send files to your chat partner.
- **Simple Configuration**: Configure the server and ports using a `config.json` file.
- **Cross-Platform**: Works on Windows, macOS, and Linux.

## Getting Started

### Prerequisites

- Go installed on your machine (version 1.16 or higher recommended).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gochat.git
   ```
2. Navigate to the project directory:
   ```bash
   cd gochat
   ```
3. To install dependencies, run the following command:
   ```bash
   go mod tidy
   ```

### Configuration

   - Edit the `config.json` file to set the server address and ports:
      ```json
      {
         "your_server": "0.0.0.0",
         "your_port": 8080,
         "friend_server" : "0.0.0.0"
         "friend_port": 8000
      }
      ```
      - `your_server`: The IP address or hostname of your server. Use `localhost` for local testing.
      - `your_port`: The port on which your server will listen for incoming connections.
      - `friend_server`: The IP address or hostname of friend server. Use `localhost` for local testing.
      - `friend_port`: The port on which your friend's server is listening. This is where your client will connect.

- Make sure the `your_port` and `friend_port` values are different for each user to avoid conflicts.

### Running the Application

**Using command line**:
1. Start the chat application:
   ```bash
   go run chat.go
   ```
2. Enter your name when prompted.
3. The application will attempt to connect to your friend's chat instance. If the connection is successful, you can start chatting.

**Using script**
1. run this command on you terminal
   ```bash
   ./run_chat.sh
   ```
   this command create two terminals and run chat with correct configs on each one of them. Use `ctrl+D `to stpo terminals

### Start Chatting:
   - Type your messages in the terminal and press `Enter` to send them.
   - To exit the chat, type `exit`.

### Sending Files

To send a file, use the following command in the chat:

```
<SEND /path/to/your/file>
```

The file will be sent to your chat partner, and they will receive it in the same directory where their chat application is running.

## Usage

1. User A starts the chat application:
   ```bash
   go run chat.go
   ```
   - Uses `your_port: 8080` and `friend_port: 8000`.
   - Enters name: `Alice`.

2. User B starts the chat application:
   ```bash
   go run chat.go
   ```
   - Uses `your_port: 8000` and `friend_port: 8080`.
   - Enters name: `Bob`.

3. Alice and Bob can now chat and send files to each other.



## Project Structure
```
gochat
├── chat.go
├── client
│   ├── client.go
│   └── colors.go
├── config.json
├── go.mod
├── go.sum
├── main_test.go
├── README.md
├── run_chat.sh
└── server
    └── server.go
```
- **chat.go**: The main entry point of the application. It handles the configuration and starts the client and server.
- **client/client.go**: Contains the client logic for sending messages and files.
- **server/server.go**: Contains the server logic for receiving messages and files.
- **config.json**: Configuration file for setting the server address and ports.
