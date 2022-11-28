#!/usr/bin/bash


cat <<EOF > private-cloud1.xml
<network>
 <name>cloud1</name>
 <forward mode='bridge'/>
 <bridge name='sw1'/>
  <portgroup name='vlan200' default='yes'>
    <vlan>
      <tag id='200'/>
    </vlan>
  </portgroup>
<virtualport type='openvswitch'/>
</network>
EOF

cat <<EOF > private-cloud2.xml
<network>
 <name>cloud2</name>
 <forward mode='bridge'/>
 <bridge name='sw1'/>
  <portgroup name='vlan400' default='yes'>
    <vlan>
      <tag id='400'/>
    </vlan>
  </portgroup>
<virtualport type='openvswitch'/>
</network>
EOF

virsh net-define --file private-cloud1.xml
virsh net-define --file private-cloud2.xml
virsh net-list
virsh net-list --all
virsh net-start  cloud1
virsh net-autostart cloud1
virsh net-dumpxml cloud1

virsh net-start  cloud2
virsh net-autostart cloud2
virsh net-dumpxml cloud2




cat <<EOF > private-cloud3.xml
<network>
 <name>cloud3</name>
 <forward mode='bridge'/>
 <bridge name='sw1'/>
  <portgroup name='vlan600' default='yes'>
    <vlan>
      <tag id='600'/>
    </vlan>
  </portgroup>
<virtualport type='openvswitch'/>
</network>
EOF

cat <<EOF > private-cloud4.xml
<network>
 <name>cloud4</name>
 <forward mode='bridge'/>
 <bridge name='sw1'/>
  <portgroup name='vlan800' default='yes'>
    <vlan>
      <tag id='800'/>
    </vlan>
  </portgroup>
<virtualport type='openvswitch'/>
</network>
EOF

virsh net-define --file private-cloud3.xml
virsh net-define --file private-cloud4.xml
virsh net-list
virsh net-list --all
virsh net-start  cloud3
virsh net-autostart cloud3
virsh net-dumpxml cloud3

virsh net-start  cloud4
virsh net-autostart cloud4
virsh net-dumpxml cloud4

virsh net-list --all
