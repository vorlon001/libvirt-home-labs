version: 2
ethernets:{{$INTERFACEINIT:=(index .Config.INTERFACEINIT 0)}}{{$NodeId:=.NodeId}}{{$MagicMac:=.Network.MagicMac }}{{range $element := .Config.INTERFACEINIT}}
    {{ $element.Name }}:
        dhcp4: false
        dhcp6: false
        match:
            macaddress: {{ $element.MacAddress }}
            name: enp*s0
        set-name: {{ $element.Name }}{{end}}
vlans:
{{range $key, $element := .Network.Block}}
    {{ $INTERFACEINIT.Name}}.{{$key}}:
        id: {{$key}}
        link: {{ $INTERFACEINIT.Name}}{{ $length4 := len $element.Network }}{{ $length6 := len $element.Network6 }}{{if or ( ne $length4 0) ( ne $length6 0 ) }}
        addresses:
{{ $length4 := len $element.Network }}{{ if ne $length4 0 }}        - {{$element.Network}}{{$NodeId}}/{{$element.Netmask}}{{end}}
{{ $length6 := len $element.Network6 }}{{ if ne $length6 0 }}        - {{$element.Network6}}{{$NodeId}}/{{$element.Netmask6}}{{end}}
        dhcp4: false
        dhcp6: false
{{ $lengthgw := len $element.Gateway4 }}{{ if ne $lengthgw 0 }}        gateway4: {{$element.Gateway4}}{{end}}
{{ $lengthdns := len $element.Nameservers.IP }}{{ if ne $lengthdns 0 }}        nameservers:
            addresses:
            - {{$element.Nameservers.IP}}
            search:
            - {{$element.Nameservers.Search}}{{end}}{{end}}{{end}}
