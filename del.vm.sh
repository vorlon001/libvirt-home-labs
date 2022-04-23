#!/usr/bin/bash

virsh destroy node100
virsh undefine --domain node100
rm ./node100-seed.qcow2
rm ./node100.qcow2

