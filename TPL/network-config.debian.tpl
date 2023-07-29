---
version: 1
config:
{{$INTERFACEINIT:=(index .Config.INTERFACEINIT 0)}}{{$NodeId:=.NodeId}}{{$MagicMac:=.Network.MagicMac }}{{range $element := .Config.INTERFACEINIT}}
  - type: physical
    name: {{ $element.Name }}
    mac_address: {{ $element.MacAddress }}
{{end}}
{{range $key, $element := .Network.Block}}
  - type: vlan
    name: {{ $INTERFACEINIT.Name}}.{{$key}}
    vlan_link: {{ $INTERFACEINIT.Name}}
    vlan_id: {{$key}}
    mtu: 1500
    dhcp4: false

    subnets:
{{ $length6 := len $element.Network6 }}{{ if ne $length6 0 }}
    - type: static
      address: {{$element.Network6}}{{$NodeId}}
      netmask: {{$element.Netmask6}}{{end}}
{{ $length4 := len $element.Network }}{{ if ne $length4 0 }}
    - type: static
      address: {{$element.Network}}{{$NodeId}}
      netmask: 255.255.255.0
      dns_nameservers:
        - 192.168.1.10
      search:
        - cloud.local
{{end}}
{{ $lengthgw := len $element.Gateway4 }}{{ if ne $lengthgw 0 }}      routes:
      - network: 0.0.0.0
        netmask: 0.0.0.0
        gateway: {{$element.Gateway4}}{{end}}
{{end}}
  - type: nameserver
    interface: enp1s0.200
    address:
      - 192.168.1.10
    search:
      - cloud.local
