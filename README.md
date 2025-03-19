# Simple Chat Application

This is a simple chat application written in Python that allows two users to communicate over a network. The application supports text messaging and file transfers between users.

## Features

- **Text Chat**: Send and receive text messages in real-time.
- **File Transfer**: Send files to the other user.
- **Configurable**: Uses a `config.json` file to specify server and port settings.

## Prerequisites

- Python version 3

## Installation

1. Clone the repository or download the source code.
   ```bash
   git clone https://github.com/yourusername/simple-chat-app.git
   cd simple-chat-app
   ```

2. Install the required dependencies listed in `requirements.txt`:
   ```bash
   pip install -r requirements.txt
   ```

3. **Configure the Application**:
   - The project already includes a `config.json` file with the necessary configuration. Here's what each field means:
     ```json
     {
         "server": "localhost",
         "your_port": 8080,
         "friend_port": 8000
     }
     ```
     - `server`: The IP address or hostname of the server. Use `localhost` for local testing.
     - `your_port`: The port on which your server will listen for incoming connections.
     - `friend_port`: The port on which your friend's server is listening. This is where your client will connect.

   - Make sure the `your_port` and `friend_port` values are different for each user to avoid conflicts.

## Usage

1. **Run the Server and Client**:
   - Open two terminal windows.
   - In the first terminal, run the chat application:
     ```bash
     python chat.py
     ```
   - Enter your name when prompted.

2. **Start Chatting**:
   - Type your messages in the terminal and press `Enter` to send them.
   - To exit the chat, type `exit`.

3. **File Transfer**:
   - To send a file, use the following format:
     ```
     <SEND /path/to/your/file>
     ```
   - When a file is sent, the receiving user will see a notification, and the file will be saved in the current directory.

## Code Structure

- `chat.py`: The main script that starts the server and client threads.
- `server.py`: Handles incoming connections and messages from clients.
- `client.py`: Connects to the server and sends messages/files.
- `config.json`: Contains configuration settings for the server and ports.
- `requirements.txt`: Lists the Python library dependencies for the project.

## Example

1. User A starts the chat application:
   ```bash
   python chat.py
   ```
   - Uses `your_port: 8080` and `friend_port: 8000`.
   - Enters name: `Alice`.

2. User B starts the chat application:
   ```bash
   python chat.py
   ```
   - Uses `your_port: 8000` and `friend_port: 8080`.
   - Enters name: `Bob`.

3. Alice and Bob can now chat and send files to each other.
