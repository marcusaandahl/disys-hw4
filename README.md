# disys-hw4 (Hand-in 4: Distributed Mutual Exclusion)

## _Team Bastian_

**Note:** Nodes will automatically attempt to populate on localhost:4XXX, starting from 4000. If a port is already occupied, a node will attempt to increment and initiate. After attempting with 40 different ports, the node will **not** be started.

## Starting the program

Start a node by running the main.go file in the directory.

- The port will automatically be decided and placed in the *usedPorts.txt* file. When the node is no longer in use, it should be removed automatically from the _.txt_ file. If you experience issues, verify that _usedPorts.txt_ is empty upon program start.
- A sequential ID is added to the node upon start up. This ID determines which node gains access in the case of multiple nodes attempting access to the critical section at the same time.



## Commands - accessing critical section

To request access to the critical section, use command `access`. If another node with a greater ID is also requesting access or the critical section is already in use the request will be denied and held until the resource is available again.



To release the critical resource use command `release`



To exit the program use command `exit`