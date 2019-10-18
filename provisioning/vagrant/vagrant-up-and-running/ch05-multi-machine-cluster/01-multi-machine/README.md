# vagrant private network names

When specifying a static IP address with host-only networks on Vagrant, Vagrant by default uses a subnet mask of 255.255.255.0. This means that as long as the first three parts (octets) of the IP address are the same, the machines will be placed on the same network.
