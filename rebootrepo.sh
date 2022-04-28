#!bin/bash

echo "Removing veritone/engines"
sudo rm -rf /home/guido/workspace/engines


echo "Cloning veritone/engines.git and giving chown permissions"
sudo git clone https://github.com/veritone/engines.git /home/guido/workspace/engines/

cd /home/guido/workspace
sudo chown guido:guido engines/ -R

