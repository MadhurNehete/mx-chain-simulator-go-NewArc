#!/bin/bash

sudo apt update
sudo apt install python3-pip
python3 -m pip install --upgrade pip --break-system-packages
cd ../..
pip3 install --break-system-packages -r examples/requirements.txt
pip3 install --break-system-packages pytest
