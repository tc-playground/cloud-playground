#!/bin/bash

name="htop-base"

vagrant up
vagrant package ${name}
vagrant destroy
