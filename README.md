# mobile-distributed-system
<!-- ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/SimonBerghem/mobile-distributed-system) -->
![Github top language](https://img.shields.io/github/languages/top/SimonBerghem/mobile-distributed-system)
![GitHub language count](https://img.shields.io/github/languages/count/SimonBerghem/mobile-distributed-system)
![Lines of code](https://img.shields.io/tokei/lines/github/SimonBerghem/mobile-distributed-system)
![GitHub repo size](https://img.shields.io/github/repo-size/SimonBerghem/mobile-distributed-system)
![GitHub license](https://img.shields.io/github/license/SimonBerghem/mobile-distributed-system)

This is a distributed system that uses the Kademlia algorithm as its communication coordinator between nodes. 

## Spin up the network
1. Open a terminal and move into the project root folder, `/mobile-distributed-system`
2. Run `sudo sh ./deploy.sh` to spin up the network of Kademlia nodes and show node status.

## Test communication (Not Kademlia commmunication)
In the same terminal as the network was spun up, run `sudo docker exec -it [NODE_1_NAME] ping [NODE_2_NAME]`. Replace the `[]`-objects with actual node names from the list of active nodes, the list is shown when running `sudo docker ps`.