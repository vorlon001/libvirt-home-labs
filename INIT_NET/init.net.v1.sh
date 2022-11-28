#!/usr/bin/bash
cat <<EOF > private.xml
<network>
  <name>private</name>
  <bridge name="virbr2"/>
  <ip address="192.168.152.1" netmask="255.255.255.0">
    <dhcp>
      <range start="192.168.152.2" end="192.168.152.254"/>
    </dhcp>
  </ip>
</network>
EOF


cat <<EOF > public.xml
<network>
  <name>public</name>
  <forward mode='nat'>
    <nat>
      <port start='1024' end='65535'/>
    </nat>
  </forward>
  <bridge name='virbr10' stp='on' delay='0'/>
  <ip address='11.0.0.1' netmask='255.255.255.0'>
    <dhcp>
      <range start='11.0.0.2' end='11.0.0.254'/>
    </dhcp>
  </ip>
</network>
EOF


virsh net-define --file private.xml
virsh net-list
virsh net-list --all
virsh net-start  private
virsh net-autostart private
virsh net-dumpxml private

virsh net-define --file public.xml
virsh net-list
virsh net-list --all
virsh net-start  public
virsh net-autostart public
virsh net-dumpxml public
