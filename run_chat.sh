#!/bin/bash

# Script to run two instances of the gochat application with swapped ports
# This allows both instances to connect to each other

# Configuration
CONFIG_FILE="config.json"
BACKUP_FILE="config.json.backup"
TERMINAL_EMULATORS=("gnome-terminal" "konsole" "xterm" "terminator")

# Function to find an available terminal emulator
find_terminal() {
    for term in "${TERMINAL_EMULATORS[@]}"; do
        if command -v "$term" &> /dev/null; then
            echo "$term"
            return 0
        fi
    done
    echo "none"
}

# Function to swap ports in config file
swap_ports() {
    # Use Python to safely parse and modify JSON
    python3 - <<EOF
import json
import sys

try:
    with open("$CONFIG_FILE", "r") as f:
        config = json.load(f)
    
    # Swap the ports
    your_port = config["your_port"]
    friend_port = config["friend_port"]
    config["your_port"] = friend_port
    config["friend_port"] = your_port
    
    with open("$CONFIG_FILE", "w") as f:
        json.dump(config, f, indent=2)
    
    print("Swapped ports: your_port is now", config["your_port"], "and friend_port is now", config["friend_port"])
except Exception as e:
    print("Error swapping ports:", e)
    sys.exit(1)
EOF
}

# Function to cleanup on exit
cleanup() {
    echo "Cleaning up..."
    # Restore original config if backup exists
    if [ -f "$BACKUP_FILE" ]; then
        mv "$BACKUP_FILE" "$CONFIG_FILE"
        echo "Restored original configuration"
    fi
    # Kill any background processes
    kill $(jobs -p) 2>/dev/null
    rm gochat
    exit
}

# Set up trap to ensure cleanup happens on script exit
trap cleanup EXIT INT TERM

# Check if config file exists
if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: $CONFIG_FILE not found!"
    exit 1
fi

# Create a backup of the original config
cp "$CONFIG_FILE" "$BACKUP_FILE"
echo "Created backup of configuration file"

# Find an available terminal emulator
TERMINAL=$(find_terminal)
if [ "$TERMINAL" = "none" ]; then
    echo "Warning: No terminal emulator found. Both instances will run in the current terminal."
fi

# Build the application first to avoid multiple simultaneous builds
echo "Building the application..."
go build -o gochat .

# Run the first instance
echo "Starting first instance..."
if [ "$TERMINAL" = "none" ]; then
    # Run in background if no terminal available
    ./gochat &
    FIRST_PID=$!
else
    # Run in a new terminal
    case $TERMINAL in
        "gnome-terminal")
            gnome-terminal --title="GoChat Instance 1" -- ./gochat &
            ;;
        "konsole")
            konsole --new-tab -e ./gochat -p "GoChat Instance 1" &
            ;;
        "xterm")
            xterm -title "GoChat Instance 1" -e ./gochat &
            ;;
        "terminator")
            terminator --title="GoChat Instance 1" -e ./gochat &
            ;;
    esac
    FIRST_PID=$!
fi

# Wait a moment for the first instance to start its server
sleep 3

# Swap ports for the second instance
echo "Swapping ports for second instance..."
swap_ports

# Run the second instance
echo "Starting second instance..."
if [ "$TERMINAL" = "none" ];  then
    # Run in foreground if no terminal available
    ./gochat
else
    # Run in a new terminal
    case $TERMINAL in
        "gnome-terminal")
            gnome-terminal --title="GoChat Instance 2" -- ./gochat
            ;;
        "konsole")
            konsole --new-tab -e ./gochat -p "GoChat Instance 2"
            ;;
        "xterm")
            xterm -title "GoChat Instance 2" -e ./gochat
            ;;
        "terminator")
            terminator --title="GoChat Instance 2" -e ./gochat
            ;;
    esac
fi

# Wait for user to press Ctrl+C
echo "Both instances are running. Press Ctrl+C to stop."
wait