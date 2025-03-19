# GoChat

GoChat is a simple command-line chat application written in Go. It allows two users to communicate over a local network or the internet. The application supports both text messaging and file transfers.

## Features

- **Real-time Messaging**: Send and receive messages in real-time.
- **File Transfer**: Send files to your chat partner.
- **Simple Configuration**: Configure the server and ports using a `config.json` file.
- **Cross-Platform**: Works on Windows, macOS, and Linux.

## Project Structure

├── chat.go

├── client

│   └── client.go

├── config.json

├── go.mod

├── server

│   └── server.go


- **chat.go**: The main entry point of the application. It handles the configuration and starts the client and server.
- **client/client.go**: Contains the client logic for sending messages and files.
- **server/server.go**: Contains the server logic for receiving messages and files.
- **config.json**: Configuration file for setting the server address and ports.

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

### Configuration

Edit the `config.json` file to set the server address and ports:

```json
{
    "server": "localhost",
    "your_port": 8080,
    "friend_port": 8000
}
```

- `server`: The IP address or hostname of the server.
- `your_port`: The port on which your instance of the chat application will listen.
- `friend_port`: The port on which your friend's instance of the chat application will listen.

### Running the Application

1. Start the chat application:
   ```bash
   go run gochat.go
   ```
2. Enter your name when prompted.
3. The application will attempt to connect to your friend's chat instance. If the connection is successful, you can start chatting.

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
   go run gochat.go
   ```
   - Uses `your_port: 8080` and `friend_port: 8000`.
   - Enters name: `Alice`.

2. User B starts the chat application:
   ```bash
   go run gochat.go
   ```
   - Uses `your_port: 8000` and `friend_port: 8080`.
   - Enters name: `Bob`.

3. Alice and Bob can now chat and send files to each other.
