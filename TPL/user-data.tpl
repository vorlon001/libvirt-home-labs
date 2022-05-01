#cloud-config
hostname: {{VMNAME}}
manage_etc_hosts: true
preserve_hostname: False
fqdn: {{VMNAME_FQDN}}

users:
  - name: vorlon
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: users, admin
    home: /home/vorlon
    shell: /bin/bash
    lock_passwd: false
    ssh-authorized-keys:
{% for item in SSHKEY[:-1] %}     - {{ item }}
{% endfor %}     - {{SSHKEY[-1:].pop()}}

ssh_pwauth: true
disable_root: false

apt:
  sources_list: |
    deb {{NEXUS_REPO}} {{VM_REPO}} main restricted
    deb {{NEXUS_REPO}} {{VM_REPO}}-updates main restricted
    deb {{NEXUS_REPO}} {{VM_REPO}} universe
    deb {{NEXUS_REPO}} {{VM_REPO}}-updates universe
    deb {{NEXUS_REPO}} {{VM_REPO}} multiverse
    deb {{NEXUS_REPO}} {{VM_REPO}}-updates multiverse
    deb {{NEXUS_REPO}} {{VM_REPO}}-backports main restricted universe multiverse
    deb {{NEXUS_REPO_SEC}} {{VM_REPO}}-security main restricted
    deb {{NEXUS_REPO_SEC}} {{VM_REPO}}-security universe
    deb {{NEXUS_REPO_SEC}} {{VM_REPO}}-security multiverse

ca-certs:
  remove-defaults: false
  trusted:
  - |
{{root_cert_append}}

chpasswd:
  list: |
    vorlon:123
    root:root
  expire: false

write_files:
 - content: |
{% for item in SSHKEY[:-1] %}    {{ item }}
{% endfor %}    {{SSHKEY[-1:].pop()}}
   path: /root/.ssh/authorized_keys
 - content: |
{% for item in sshd_config_append[:-1] %}    {{ item }}
{% endfor %}    {{sshd_config_append[-1:].pop()}}
   path: /etc/ssh/sshd_config
 - content: |
{% for item in pip_conf_append[:-1] %}    {{ item }}
{% endfor %}    {{pip_conf_append[-1:].pop()}}
   path: /etc/pip.conf

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

packages:
{% for item in PKG[:-1] %}  - {{ item }}
{% endfor %}  - {{PKG[-1:].pop()}}

output:
  all: ">> /var/log/cloud-init.log"

runcmd:
{% for item in CMD[:-1] %}  - {{ item }}
{% endfor %}  - {{CMD[-1:].pop()}}

final_message: "The system is finally up, after $UPTIME seconds"

power_state:
    delay: now
    mode: poweroff
    message: "shutdown after init"
    timeout: 20
