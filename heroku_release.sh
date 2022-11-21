#!/bin/bash

./borealis-pg-init-ssh-tunnel.sh

sleep 5

./makemigrate.sh up
