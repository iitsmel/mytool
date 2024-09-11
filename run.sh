#!/bin/bash
user_name=$(whoami)

# List of formulae to check and install
#FORMULAE=("sqlmap" "nmap")

# Function to check and install a formula
#check_and_install_formula() {
#  local formula=$1
#  if brew list --formula | grep -q "^${formula}$"; then
#    echo "${formula} is already installed."
#  else
#    echo "${formula} is not installed. Installing now..."
#    brew install "${formula}"
#  fi
#}

# Loop through the list of formulae
#for formula in "${FORMULAE[@]}"; do
#  check_and_install_formula "$formula"
#done

go build
cp mytool /usr/local/bin/mytool
rm mytool