#!/usr/bin/bash

cat <<EOF > private-sw1.xml
<network>
 <name>cloud_sw1</name>
 <forward mode='bridge'/>
 <bridge name='sw1'/>
<virtualport type='openvswitch'/>
</network>
EOF

virsh net-define --file private-sw1.xml
virsh net-list
virsh net-list --all
virsh net-start  cloud_sw1
virsh net-autostart cloud_sw1
virsh net-dumpxml cloud_sw1
virsh net-list --all
