#cloud-config
hostname: {{.VMNAME}}
manage_etc_hosts: true
preserve_hostname: False
fqdn: {{.VMNAME_FQDN}}

users:
  - name: vorlon
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: users, admin
    home: /home/vorlon
    shell: /bin/bash
    lock_passwd: false
    ssh-authorized-keys:{{range $key, $element := .SSHKeys }}
     - {{ $element }}{{end}}

ssh_pwauth: true
disable_root: false


apt:
  sources_list: |
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}} main restricted
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}}-updates main restricted
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}} universe
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}}-updates universe
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}} multiverse
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}}-updates multiverse
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}}-backports main restricted universe multiverse
    deb {{.Config.NEXUSREPOSEC}} {{.Config.VMREPO}}-security main restricted
    deb {{.Config.NEXUSREPOSEC}} {{.Config.VMREPO}}-security universe
    deb {{.Config.NEXUSREPOSEC}} {{.Config.VMREPO}}-security multiverse

ca-certs:
  remove-defaults: false
  trusted:
  - |{{range $key, $element := .ReadFromFile "TPL/root.iblog.pro.crt" }}
     {{ $element }}{{end}}

chpasswd:
  list: |
    vorlon:123
    root:root
  expire: false


write_files:
 - content: |{{range $key, $element := .SSHKeys }}
    {{ $element }}{{end}}
   path: /root/.ssh/authorized_keys
 - content: |{{range $key, $element := .ReadFromFile "TPL/sshd-config.tpl" }}
    {{ $element }}{{end}}
   path: /etc/ssh/sshd_config
 - content: |{{range $key, $element := .ReadFromFile "TPL/pip-conf.tpl" }}
    {{ $element }}{{end}}
   path: /etc/pip.conf
 - content: |
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}} main
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}}-updates main
    deb {{.Config.NEXUSREPO}} {{.Config.VMREPO}}-backports main
    deb {{.Config.NEXUSREPOSEC}} {{.Config.VMREPO}}-security main
   path: /etc/apt/sources.list

ntp:
  enabled: true
  ntp_client: systemd-timesyncd
  servers:
    - 0.ru.pool.ntp.org
    - 1.ru.pool.ntp.org
    - 2.ru.pool.ntp.org
    - 3.ru.pool.ntp.org

growpart:
  mode: auto
  devices: ['/']

timezone: Asia/Yekaterinburg
package_update: true
package_upgrade: true

packages:{{range $key, $element := .Pgk }}
  - {{ $element }}{{end}}

output:
  all: ">> /var/log/cloud-init.log"

runcmd:{{range $key, $element := .Command }}
  - {{ $element }}{{end}}

final_message: "The system is finally up, after $UPTIME seconds"

power_state:
    delay: now
    mode: poweroff
    message: "shutdown after init"
    timeout: 20
