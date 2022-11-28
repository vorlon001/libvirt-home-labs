#cloud-config
vyos_config_commands:
- set service ssh port 22
- set system login user vyos authentication plaintext-password 'vyos'
- set interfaces ethernet eth0 description 'uplink 1'
- set interfaces ethernet eth1 description 'uplink 2'
- set system host-name '{{VMNAME}}'
- set system login banner pre-login 'VyOS router {{VMNAME}}'
{% for key,item in node_ip.items() %}
- set interfaces ethernet eth0 vif {{key}} description 'VLAN {{key}}'
- set interfaces ethernet eth0 vif {{key}} address '{{item['ipaddress']}}/24'
{% endfor %}
