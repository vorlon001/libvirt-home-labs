function init_network_config_ubuntu() {
cat >$VMNAME/network-config <<EOF
ethernets:
    $INTERFACE_200:
        addresses:
        - $IP_ADDR_200/$IP_MASK
        dhcp4: false
        gateway4: $IP_GW
        match:
            macaddress: $MAC_ADDR_200
        nameservers:
            addresses:
            - $IP_DNS
        set-name: $INTERFACE_200
    $INTERFACE_400:
        addresses:
        - $IP_ADDR_400/$IP_MASK
        dhcp4: false
        match:
            macaddress: $MAC_ADDR_400
        set-name: $INTERFACE_400
version: 2
EOF
}

function  create_cloud_config_ubuntu() {

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

apt:
  sources_list: |
    deb https://archive.ubuntu.com/ $1 main restricted
    deb https://archive.ubuntu.com/ $1-updates main restricted
    deb https://archive.ubuntu.com/ $1 universe
    deb https://archive.ubuntu.com/ $1-updates universe
    deb https://archive.ubuntu.com/ $1 multiverse
    deb https://archive.ubuntu.com/ $1-updates multiverse
    deb https://archive.ubuntu.com/ $1-backports main restricted universe multiverse
    deb https://security.ubuntu.com/ $1-security main restricted
    deb https://security.ubuntu.com/ $1-security universe
    deb https://security.ubuntu.com/ $1-security multiverse

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
    Include /etc/ssh/sshd_config.d/*.conf
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

timezone: Europe/Moscow
package_update: true
package_upgrade: true

packages:
  - apt-transport-https
  - ca-certificates
  - curl
  - gnupg
  - lsb-release
  - openssh-server
  - net-tools
  - tcpdump
  - mc
  - htop
  - iotop
  - iftop
  - language-pack-ru

output:
  all: ">> /var/log/cloud-init.log"
runcmd:
  - systemctl disable systemd-udevd.service
  - ssh-keygen -A
  - ssh-keygen -t rsa -b 4096 -f /root/.ssh/id_rsa  -q -P ""
  - ssh-keygen -t rsa -b 4096 -f /home/vorlon/.ssh/id_rsa  -q -P ""
  - cp /etc/netplan/50-cloud-init.yaml /etc/netplan/00-installer-config.yaml
  - echo "LANGUAGE=ru_RU:ru" > /etc/default/locale
  - echo "LANG=ru_RU.UTF-8" >> /etc/default/locale
  - touch /etc/cloud/cloud-init.disabled
  - apt update

final_message: "The system is finally up, after $UPTIME seconds"

power_state:
    delay: now
    mode: poweroff
    message: "shutdown after init"
    timeout: 20

EOF
}
