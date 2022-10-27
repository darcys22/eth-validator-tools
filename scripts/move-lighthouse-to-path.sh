#!/bin/bash

sudo systemctl stop lighthousevalidator.service 
sudo systemctl stop lighthousebeacon.service 
sudo cp lighthouse /usr/local/bin/
sudo systemctl start lighthousebeacon.service 
sudo systemctl start lighthousevalidator.service 
rm lighthouse*

