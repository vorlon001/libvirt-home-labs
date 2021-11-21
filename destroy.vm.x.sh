#!/bin/bash

set -x

NODENAME=$1

virsh destroy ${NODENAME}
virsh undefine --domain ${NODENAME}
rm $NODENAME/$NODENAME-disk2.qcow2
rm $NODENAME/$NODENAME-disk3.qcow2
rm $NODENAME/$NODENAME-disk4.qcow2
rm $NODENAME/$NODENAME-seed.qcow2
rm $NODENAME/$NODENAME.qcow2

zfs destroy  clouds/$NODENAME
zfs destroy  clouds/$NODENAME-disk2
zfs destroy  clouds/$NODENAME-disk3
zfs destroy  clouds/$NODENAME-disk4

rm -R $NODENAME
