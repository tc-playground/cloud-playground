# Vagrant Image from Scratch

* Create a virtual machine manually in VirtualBox.

* Configure this virtual machine as required. 
    * configure drive  40 to 100 GB.
    * set default memory ~512MB +
    * disable audio, USB, and other nonessential peripheral controllers.

* Configure NAT device in interface 1. Vagrant requires this for SSH.

* Create 'vagrant' defaut ssh user.
    * You can always use another SSH user by specifying the config.ssh.username variable.
    * Ensure an SSH server is installed and properly configured to run on system boot.
    * Set up the SSH user to authenticate using the insecure private key that ships with Vagrant.
      ```
      $ mkdir /home/vagrant/.ssh
      $ chmod 700 /home/vagrant/.ssh
      $ cd /home/vagrant/.ssh
      $ wget --no-check-certificate 'https://raw.github.com/mitchellh/vagrant/master/ keys/vagrant.pub' -O authorized_keys
      $ chmod 600 /home/vagrant/.ssh/authorized_keys
      $ chown -R vagrant /home/vagrant/.ssh
      ```
    or
    * Configure the 'config.ssh.private_key_path' variable later to point to the proper private key.
    * The SSH user must have permissions to use sudo without a password. Enusre 'requiretty' is disabled.

* Install VirtualBox Guest Additions





