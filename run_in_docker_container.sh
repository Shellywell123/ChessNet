#!/bin/bash
echo """
----------------------------------------------------------------------------------------
╔╗ ┌─┐┌┐┌┌─┐   ╔╦╗┌─┐┌─┐┬┌─┌─┐┬─┐   ╔═╗┌┐┌┬  ┬   ╔═╗┌─┐┬─┐┬┌─┐┌┬┐
╠╩╗├┤ │││└─┐─ ──║║│ ││  ├┴┐├┤ ├┬┘───║╣ │││└┐┌┘───╚═╗│  ├┬┘│├─┘ │
╚═╝└─┘┘└┘└─┘   ═╩╝└─┘└─┘┴ ┴└─┘┴└─   ╚═╝┘└┘ └┘    ╚═╝└─┘┴└─┴┴   ┴ .sh
----------------------------------------------------------------------------------------"""

echo "----------------------------------------------------------------------------------------"
echo "--> stopping any current containers..."
docker stop chessnetcontainer

echo "----------------------------------------------------------------------------------------"
echo "--> removing any preexisting containers..."
docker rm -f chessnetcontainer

echo "----------------------------------------------------------------------------------------"
echo "--> removing any preexisting images..."
docker image rm -f chessnet-server

echo "----------------------------------------------------------------------------------------"
echo "--> running up compose."
docker compose up
