---
version: 1
config:
{% for item in INTERFACE %}  - type: physical
    name: {{ item['name'] }}
    mac_address: {{ item['mac'] }}
{% endfor %}
