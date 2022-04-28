#!bin/bash

echo "Removing veritone/engines"
sudo rm -rf ~/workspace/engines


echo "Cloning veritone/engines.git and giving chown permissions"
sudo git clone https://github.com/veritone/engines.git -C ~/workspace/

cd ~/workspace
sudo chown guido:guido engines/ -R

