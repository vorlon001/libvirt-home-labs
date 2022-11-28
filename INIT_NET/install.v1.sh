#!/usr/bin/bash

apt -y install qemu-kvm libvirt-daemon-system libvirt-daemon virtinst bridge-utils libosinfo-bin libguestfs-tools virt-top

apt install -y nano mc net-tools tcpdump
apt install -y mc ansible pdsh
apt install -y cloud-utils
apt -y install lm-sensors


wget https://cloud-images.ubuntu.com/focal/20210222/focal-server-cloudimg-amd64.img
mv focal-server-cloudimg-amd64.img /home/vorlon/KVM


DATE=$(date +%Y-%m-%d-%H-%M-%S)
cp /etc/ssh/sshd_config /etc/ssh/sshd_config.old.$DATE
sed "s/#PermitRootLogin prohibit-password/PermitRootLogin yes/g" /etc/ssh/sshd_config.old.$DATE > /etc/ssh/sshd_config
systemctl restart ssh && systemctl status ssh 

cp /etc/sudoers /etc/sudoers.$(date +%Y-%m-%d-%H-%M-%S)
sed "s/# Members of the admin group may gain root privileges/vorlon    ALL=NOPASSWD: ALL/g" /etc/sudoers > /etc/sudoers

#sudo sensors-detect
#sensors
