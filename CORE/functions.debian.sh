function init_network_config_debian_v2() {
cat >$VMNAME/network-config <<EOF
---
version: 1
config:
- type: physical
  name: enp1s0
  mac_address: $MAC_ADDR_200
  subnets:
  - type: static
    address: $IP_ADDR_200
    netmask: 255.255.255.0
    routes:
    - network: 0.0.0.0
      netmask: 0.0.0.0
      gateway: $IP_GW
- type: physical
  name: enp2s0
  mac_address: $MAC_ADDR_400
  subnets:
  - type: static
    address: $IP_ADDR_400
    netmask: 255.255.255.0
- type: nameserver
  address: [$IP_DNS ]
EOF
}

function init_network_config_debian() {
cat >$VMNAME/network-config <<EOF
---
version: 1
config:
- type: physical
  name: eth0
  mac_address: $MAC_ADDR_200
  subnets:
  - type: static
    address: $IP_ADDR_200
    netmask: 255.255.255.0
    routes:
    - network: 0.0.0.0
      netmask: 0.0.0.0
      gateway: $IP_GW
- type: physical
  name: eth1
  mac_address: $MAC_ADDR_400
  subnets:
  - type: static
    address: $IP_ADDR_400
    netmask: 255.255.255.0
- type: nameserver
  address: [$IP_DNS ]
EOF
}



function  create_cloud_config_debian() {
cat >$VMNAME/user-data <<EOF
#cloud-config
hostname: $VMNAME
manage_etc_hosts: true
preserve_hostname: False
#network:
#  config: disabled


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

ntp:
  enabled: true
  ntp_client: chrony

ssh_pwauth: true
disable_root: false

apt:
  sources_list: |
    deb https://deb.debian.org/debian $1 main
    deb https://deb.debian.org/debian-security $1-security main
    deb https://deb.debian.org/debian $1-updates main
    deb https://deb.debian.org/debian $1-backports main

ca-certs:
  remove-defaults: false

  trusted:
  - |
$(sed s/^/\ \ \ \ \ / root.x.y.crt )


packages:
  - apt-transport-https
  - ca-certificates
  - curl
  - gnupg
  - lsb-release
#  - openssh-server
  - net-tools
  - tcpdump
  - mc
  - htop
  - iotop
  - iftop

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
    SyslogFacility AUTH
    LogLevel INFO
    Port 22
    ListenAddress 0.0.0.0
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
 - path: /etc/network/interfaces
   content: |
    source /etc/network/interfaces.d/*

growpart:
  mode: auto
  devices: ['/']

timezone: Europe/Moscow
package_update: true
package_upgrade: true

output:
  all: ">> /var/log/cloud-init.log"
runcmd:
  - systemctl disable systemd-udevd.service
  - ssh-keygen -A
  - ssh-keygen -t rsa -b 4096 -f /root/.ssh/id_rsa  -q -P ""
  - ssh-keygen -t rsa -b 4096 -f /home/vorlon/.ssh/id_rsa  -q -P ""
  - echo "LANGUAGE=ru_RU:ru" > /etc/default/locale
  - echo "LANG=ru_RU.UTF-8" >> /etc/default/locale
  - touch /etc/cloud/cloud-init.disabled
final_message: "The system is finally up, after $UPTIME seconds"

power_state:
    delay: now
    mode: poweroff
    message: "shutdown after init"
    timeout: 20

EOF
}
