#!/usr/bin/bash

virsh net-destroy cloud1
virsh net-undefine cloud1

virsh net-destroy cloud2
virsh net-undefine cloud2
