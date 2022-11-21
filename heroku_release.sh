#!/bin/bash

./borealis-pg-init-ssh-tunnel.sh && ./makemigrate.sh up
