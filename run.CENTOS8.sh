#!/usr/bin/bash

# VMNAME=CENTOS180 CORE=4 RAM=4096 DISKSIZE=20 IP-OCTET=204
set -x

MAC_ADDR_200=$(printf '52:54:00:%02x:%02x:%02x' $((RANDOM%256)) $((RANDOM%256)) $((RANDOM%256)))
MAC_ADDR_400=$(printf '52:54:00:%02x:%02x:%02x' $((RANDOM%256)) $((RANDOM%256)) $((RANDOM%256)))
echo $MAC_ADDR_200 $MAC_ADDR_400


OSvariant=centos8
OSvariant=fedora33

VMNAME=$1 #CSR1
CORE=$2
MEMORY=$3
IMAGES_PATH='IMAGES'
mkdir $VMNAME
DISKROOTSIZE=$4

SSHKEY=$(cat /root/.ssh/id_rsa.pub)
SSHKEY2=$(cat /root/.ssh/authorized_keys | grep node)

IP_SET_200="192.168.200"
IP_SET_400="192.168.201"
IP_MASK=24
IP_GW=192.168.200.1
IP_DNS=192.168.1.10

OCTET=$5
IP_ADDR_200=$IP_SET_200.$OCTET
IP_ADDR_400=$IP_SET_400.$OCTET


function init_network_config_centos() {
cat >$VMNAME/network-config <<EOF
version: 2
ethernets:
  eth0:
     dhcp4: false
     dhcp6: false
     # default libvirt network
     addresses: [ $IP_ADDR_200/$IP_MASK ]
     gateway4: $IP_GW
     nameservers:
       addresses: [ $IP_DNS ]
  eth1:
     dhcp4: false
     dhcp6: false
     # default libvirt network
     addresses: [ $IP_ADDR_400/$IP_MASK ]
EOF
}


function  init_cloud_config() {
  touch $VMNAME/meta-data
  cloud-localds -v --network-config=$VMNAME/network-config $VMNAME/$VMNAME-seed.qcow2 $VMNAME/user-data $VMNAME/meta-data
}


function create_disk_file_initor {
    qemu-img create -F qcow2 -b $(python3 -c 'print(__import__("pathlib").Path(__import__("os").getcwd()).parent)')/$IMAGES_PATH/$1 -f qcow2 $VMNAME/$VMNAME.qcow2 ${DISKROOTSIZE}G
}

function  create_cloud_config_centos8() {

cat >$VMNAME/user-data <<EOF
#cloud-config
hostname: $VMNAME
manage_etc_hosts: true
preserve_hostname: False
fqdn: $VMNAME.cloud.local

users:
  - name: vorlon
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: users, admin
    home: /home/vmadm
    shell: /bin/bash
    lock_passwd: false
    ssh-authorized-keys:
     - $SSHKEY
     - $SSHKEY2

ssh_pwauth: true
disable_root: false

ca-certs:
  remove-defaults: false
  trusted:
  - |
$(sed s/^/\ \ \ \ \ / root.x.y.crt )

chpasswd:
  list: |
    vorlon:123
    root:root
  expire: false

write_files:
 - content: |
       $SSHKEY
       $SSHKEY2
   path: /root/.ssh/authorized_keys
 - path: /etc/ssh/sshd_config
   content: |
    Port 22
    ListenAddress 0.0.0.0
    SyslogFacility AUTH
    LogLevel INFO
    PermitRootLogin yes
    StrictModes yes
    MaxAuthTries 6
    ChallengeResponseAuthentication no
    UsePAM yes
    AllowTcpForwarding yes
    X11Forwarding no
    PrintMotd no
    AcceptEnv LANG LC_*
    Subsystem sftp  /usr/lib/openssh/sftp-server
    PasswordAuthentication yes

growpart:
  mode: auto
  devices: ['/']

timezone: Asia/Yekaterinburg
package_update: true
package_upgrade: true

output:
  all: ">> /var/log/cloud-init.log"
runcmd:
  - systemctl disable systemd-udevd.service
  - ssh-keygen -A
  - ssh-keygen -t rsa -b 4096 -f /root/.ssh/id_rsa  -q -P ""
  - ssh-keygen -t rsa -b 4096 -f /home/vorlon/.ssh/id_rsa  -q -P ""

final_message: "The system is finally up, after $UPTIME seconds"

power_state:
    delay: now
    mode: poweroff
    message: "shutdown after init"
    timeout: 20

EOF
}

#create_disk_file_initor "CentOS-Stream-GenericCloud-8-20210603.0.x86_64.qcow2"
create_disk_file_initor "Fedora-Cloud-Base-34-1.2.x86_64.qcow2"

init_network_config_centos
create_cloud_config_centos8
init_cloud_config


function virt_install_one_file() {
# osinfo-query os
virt-install --connect qemu:///system --virt-type kvm --name $VMNAME --ram $MEMORY --vcpus=$CORE --cpu Westmere --os-type linux --os-variant $OSvariant \
 --disk path=$VMNAME/$VMNAME.qcow2,device=disk \
 --disk path=$VMNAME/$VMNAME-seed.qcow2,device=disk --import \
 --network network=cloud1,mac=$MAC_ADDR_200 \
 --network network=cloud2,model=virtio,mac=$MAC_ADDR_400 --noautoconsole
}

virt_install_one_file
