#!/usr/bin/env bash

set -e

if [ -f "docker-ce_17.12.1~ce-0~ubuntu_amd64.deb" ]; then
    echo "already downloaded docker deb"
else
    wget https://download.docker.com/linux/ubuntu/dists/artful/pool/stable/amd64/docker-ce_17.12.1~ce-0~ubuntu_amd64.deb
fi
sudo apt install -y libltdl7
sudo dpkg -i docker-ce_17.12.1~ce-0~ubuntu_amd64.deb
sudo usermod -aG docker $USER
# FIXME: https://superuser.com/questions/272061/reload-a-linux-users-group-assignments-without-logging-out
# this requires password ...
#su - $USER
# FIXME: without login and logout, group does not take effect in current session
docker version