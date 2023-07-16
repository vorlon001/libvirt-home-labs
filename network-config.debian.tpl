---
version: 1
config:
{% for item in INTERFACE %}  - type: physical
    name: {{ item['name'] }}
    mac_address: {{ item['mac'] }}
{% endfor %}
{% for key,item in node_ip.items() %}
  - type: vlan
    name: {{INTERFACE[0]['name']}}.{{key}}
    vlan_link: {{INTERFACE[0]['name']}}
    vlan_id: {{key}}
    mtu: 1500
    dhcp4: false
    subnets:
{% if 'network6' in item %}    - type: static
      address: {{item['network6']}}{{item['octet']}}
      netmask: {{item['netmask6']}}{% endif %}
    - type: static
      address: {{item['ipaddress']}}
      netmask: 255.255.255.0
      dns_nameservers:
        - 192.168.1.10
      search:
        - cloud.local
{% if 'gateway4' in item %}      routes:
      - network: 0.0.0.0
        netmask: 0.0.0.0
        gateway: {{item['gateway4']}}{% endif %}
{% endfor %}
  - type: nameserver
    interface: enp1s0.200
    address:
      - 192.168.1.10
    search:
      - cloud.local
