#!/bin/bash

VARIABLE_NAME="PROTOFILES_PATH"

# Check rights
if [ "$EUID" -ne 0 ]; then
  echo "Please run with sudo"
  exit 1
fi

# Path definition
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

# Variable adding /etc/environment
if grep -q "^$VARIABLE_NAME=" /etc/environment; then
  # Updating if exists
  sudo sed -i "s|^$VARIABLE_NAME=.*|$VARIABLE_NAME=$SCRIPT_DIR|" /etc/environment
else
  # Else adding
  echo "$VARIABLE_NAME=$SCRIPT_DIR" | sudo tee -a /etc/environment > /dev/null
fi

echo "Variable $VARIABLE_NAME = $SCRIPT_DIR is set."