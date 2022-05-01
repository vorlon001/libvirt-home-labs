#!/usr/bin/bash

virsh destroy node100
virsh undefine --domain node100
rm -R ./node100

