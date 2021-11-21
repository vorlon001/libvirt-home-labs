#!/usr/bin/bash


if [ "${10}" = "ON" ]; then
    set -x
fi

echo $1 $2 $3 $4

VMNAME=$1
OCTET=$2
OSVER=$3
CORE=$4
MEMORY=$5
ADDON_DISK=$6
OSvariant=$7
TYPE_DISK=$8
DISKROOTSIZE=$9
DEBUGMODE=${10}
PWD=$(pwd)
IP_SET_200="192.168.200"
IP_SET_400="192.168.201"
IP_MASK=24
IP_GW=192.168.200.1
IP_DNS=192.168.1.10


IMAGEADDONSIZE=50
IMAGES_PATH='IMAGES'

source CORE/functions.sh
source CORE/functions.ubuntu.sh
source CORE/functions.debian.sh
source CORE/functions.utils.sh



mkdir $VMNAME

if [[ $ADDON_DISK = "ON" ]]; then
  if [[ $TYPE_DISK = "FILE" ]]; then
      create_disk_addon_file 2 4
  elif [[ $TYPE_DISK = "ZFS" ]]; then
      create_disk_addon_zfs 2 4
  else
      exit
  fi
fi

if [[ $TYPE_DISK = "FILE" ]]; then
    create_disk_raw  $OSVER
elif [[ $TYPE_DISK = "ZFS" ]]; then
    create_disk_zfs  $OSVER
else
    exit
fi


#cp  ../focal-server-cloudimg-amd64.img $VMNAME.qcow2

init
ssh_key_remove

case $OSvariant in
  'ubuntu20.04')
    init_network_config_ubuntu
    if [[ $OSVER = "21.04" ]]; then
        create_cloud_config_ubuntu "hirsute"
    elif [[ $OSVER = "21.10" ]]; then
        create_cloud_config_ubuntu "impish"
    elif [[ $OSVER = "22.04" ]]; then
        create_cloud_config_ubuntu "jammy"
    elif [[ $OSVER = "20.04" ]]; then
        create_cloud_config_ubuntu "focal"
    else:
        echo "OS not CORRECT POINT"
        exit 1
    fi
    ;;
  'debian10')
    case $OSVER in
        '11')
        init_network_config_debian "bullseye"
        ;;
        '11v2')
        init_network_config_debian_v2 "bullseye"
        ;;
    esac
    create_cloud_config_debian "bullseye"
    ;;
esac

init_cloud_config

case $ADDON_DISK in
  'ON')
    if [[ $TYPE_DISK = "ZFS" ]]; then
        virt_install_addon_disk_zfs
    elif [[ $TYPE_DISK = "FILE" ]]; then
        virt_install_addon_disk_file
    else
        exit
    fi
    ;;
  *)
    if [[ $TYPE_DISK = "ZFS" ]]; then
        virt_install_one_zfs
    elif [[ $TYPE_DISK = "FILE" ]]; then
        virt_install_one_file
    else
        exit
    fi
    ;;
esac


isShutdown
virsh start $VMNAME
isSshEnable

case $OSvariant in
  'ubuntu20.04')
    finale_ubuntu
    ;;
  'debian10')
    finale_ubuntu
    ;;
esac

isShutdown

echo "remote $VMNAME-seed.qcow2"
virsh detach-disk --domain $VMNAME $PWD/$VMNAME/$VMNAME-seed.qcow2 --persistent --config
virsh domblklist  $VMNAME

sleeps 10

echo "start $VMNAME"
virsh --connect qemu:///system start $VMNAME

isSshEnable
