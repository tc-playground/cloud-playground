# Vagrant Global State ~/.vagrant.d

* Vagrant puts all global state by default into the ~/.vagrant.d folder, including boxes. This means that when Vagrant manages boxes “globally,” it actually means it manages boxes per user, by default.

* Because boxes can be large (sometimes gigabytes), you can move this global state di‐ rectory by setting the environmental variable VAGRANT_HOME to another directory.
