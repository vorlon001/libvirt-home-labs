version: 2
ethernets:{% for item in INTERFACE %}
    {{ item['name'] }}:
        dhcp4: false  
        match:
            macaddress: {{ item['mac'] }}
        set-name: {{ item['name'] }}{% endfor %}
vlans:{% for key,item in node_ip.items() %}
    {{INTERFACE[0]['name']}}.{{key}}:
        id: {{key}}
        link: {{INTERFACE[0]['name']}}
        addresses:
        - {{item['ipaddress']}}/{{item['netmask']}}
        dhcp4: false
{% if 'gateway4' in item %}        gateway4: {{item['gateway4']}}{% endif %}
{% if 'nameservers' in item %}        nameservers:
            addresses:
            - {{item["nameservers"]["ip"]}}
            search:
            - {{item["nameservers"]["search"]}}{% endif %}{% endfor %}
