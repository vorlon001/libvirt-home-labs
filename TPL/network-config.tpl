version: 2
ethernets:{% for item in INTERFACE %}
    {{ item['name'] }}:
        dhcp4: false  
        match:
            macaddress: {{ item['mac'] }}
            name: enp*s0
        set-name: {{ item['name'] }}{% endfor %}
vlans:{% for key,item in node_ip.items() %}
    {{INTERFACE[0]['name']}}.{{key}}:
        id: {{key}}
        link: {{INTERFACE[0]['name']}}
        addresses:
        - {{item['ipaddress']}}/{{item['netmask']}}
{% if 'network6' in item %}        - {{item['network6']}}{{item['octet']}}/{{item['netmask6']}}{% endif %}
        dhcp4: false
        dhcp6: false
{% if 'gateway4' in item %}        gateway4: {{item['gateway4']}}{% endif %}
{% if 'nameservers' in item %}        nameservers:
            addresses:
            - {{item["nameservers"]["ip"]}}
            search:
            - {{item["nameservers"]["search"]}}{% endif %}{% endfor %}
