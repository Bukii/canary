#!/bin/bash

PATH=$PATH:/usr/local/go/bin

output=$(make 2>&1)
if [ $? -ne 0 ]; then
  echo "Build failed:"
  echo "$output"
  exit 1
fi

PS3='Please enter your system architecture: '
options=("Linux Amd" "Linux Arm")
select opt in "${options[@]}"
do
    case $opt in
        "Linux Amd")
            sudo cp bin/canary-linux_amd64 /usr/local/bin/canary
            break
            ;;
        "Linux Arm")
            sudo cp bin/canary-linux_arm64 /usr/local/bin/canary
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done