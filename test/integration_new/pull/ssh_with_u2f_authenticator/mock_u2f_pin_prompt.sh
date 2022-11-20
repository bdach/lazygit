#!/bin/bash
# The purpose of this script is to test the behaviour of using an SSH key which is backed by a U2F hardware authenticator.
# It intends to simulate the process of providing credentials to such an SSH key, allowing it to be plugged into git via core.askPass.

read -p "Enter PIN for ECDSA-SK key /home/lazygit/.ssh/id_ecdsa_sk: " -s PIN

if [[ $PIN -ne 1234 ]]; then
    exit 1
fi

# This is where the user is supposed to interact with the key physically, so let's just assume instant input.
printf "\nConfirm user presence for key ECDSA-SK SHA256:insert_fingerprint_here\n"