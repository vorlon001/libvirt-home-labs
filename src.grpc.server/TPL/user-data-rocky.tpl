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
