
function create_zfs_volume {
    zfs create -V $2gb clouds/$1
    zfs set sync=disabled clouds/$1
    zfs list | grep $1
}

function create_disk_addon_file {
  for ((i=$1;i<=$2;i++)); do
    echo "Iteration $i"
    qemu-img create -f qcow2 $VMNAME/$VMNAME-disk$i.qcow2 ${IMAGEADDONSIZE}G
  done
}


function create_disk_addon_zfs {
  for ((i=$1;i<=$2;i++)); do
    echo "Iteration $i"
    create_zfs_volume $VMNAME-disk$i ${IMAGEADDONSIZE}
  done
}


function create_disk_addon_file {
  for ((i=$1;i<=$2;i++)); do
    echo "Iteration $i"
    qemu-img create -f qcow2 $VMNAME/$VMNAME-disk$i.qcow2 ${IMAGEADDONSIZE}G
  done
}

function create_disk_zfs_initor {
# python3 -c 'print(__import__("pathlib").Path(__import__("os").getcwd()).parent)'
# python3 -c 'os=__import__("os"); print(os.path.abspath(os.path.join(os.getcwd(), os.pardir)))'

    qemu-img create -F qcow2 -b $(python3 -c 'print(__import__("pathlib").Path(__import__("os").getcwd()).parent)')/$IMAGES_PATH/$1 -f qcow2 $VMNAME/$VMNAME.qcow2 5G
    qemu-img convert -f qcow2 $VMNAME/$VMNAME.qcow2 $VMNAME/$VMNAME.raw
    qemu-img info $VMANE/$VMNAME.raw

    create_zfs_volume $VMNAME 5

    qemu-img convert -O raw $VMNAME/$VMNAME.raw  /dev/zvol/clouds/$VMNAME

    zfs set volsize=${DISKROOTSIZE}GB clouds/$VMNAME

    rm $VMNAME/$VMNAME.raw
}


function create_disk_file_initor {
    qemu-img create -F qcow2 -b $(python3 -c 'print(__import__("pathlib").Path(__import__("os").getcwd()).parent)')/$IMAGES_PATH/$1 -f qcow2 $VMNAME/$VMNAME.qcow2 ${DISKROOTSIZE}G
}


function create_disk_zfs {
  if [[ $1 == "20.04" ]]; then
    create_disk_zfs_initor "focal-server-cloudimg-amd64.img"
  elif [[ $1 == "21.04" ]]; then
    create_disk_zfs_initor "hirsute-server-cloudimg-amd64.img"
  elif [[ $1 == "21.10" ]]; then
    create_disk_zfs_initor "impish-server-cloudimg-amd64.img"
  elif [[ $1 == "22.04" ]]; then
    create_disk_zfs_initor "jammy-server-cloudimg-amd64.img"
  elif [[ $1 == "11v2" ]]; then
    create_disk_zfs_initor "debian-11-generic-amd64-daily-20211008-789.custom.qcow2"
  elif [[ $1 == "11" ]]; then
    create_disk_zfs_initor "debian-11-generic-amd64-20210928-779.qcow2"
  fi
}


function create_disk_raw {
  if [[ $1 == "20.04" ]]; then
    create_disk_file_initor "focal-server-cloudimg-amd64.img"
  elif [[ $1 == "21.04" ]]; then
    create_disk_file_initor "hirsute-server-cloudimg-amd64.img"
  elif [[ $1 == "21.10" ]]; then
    create_disk_file_initor "impish-server-cloudimg-amd64.img"
  elif [[ $1 == "22.04" ]]; then
    create_disk_file_initor "jammy-server-cloudimg-amd64.img"
  elif [[ $1 == "11v2" ]]; then
    create_disk_file_initor "debian-11-generic-amd64-daily-20211008-789.custom.qcow2"
  elif [[ $1 == "11" ]]; then
    create_disk_file_initor "debian-11-generic-amd64-20210928-779.qcow2"
  fi
}

function init() {
  MAC_ADDR_200=$(printf '52:54:00:%02x:%02x:%02x' $((RANDOM%256)) $((RANDOM%256)) $((RANDOM%256)))
  MAC_ADDR_400=$(printf '52:54:00:%02x:%02x:%02x' $((RANDOM%256)) $((RANDOM%256)) $((RANDOM%256)))
  echo $MAC_ADDR_200 $MAC_ADDR_400

  case $OSvariant in
    'ubuntu20.04')
       INTERFACE_200=enp1s0
       INTERFACE_400=enp2s0
      ;;
    'debian10')
       INTERFACE_200=eth0
       INTERFACE_400=eth1
      ;;
  esac


  INTERFACE_200=enp1s0
  INTERFACE_400=enp2s0
  IP_ADDR_200=$IP_SET_200.$OCTET
  IP_ADDR_400=$IP_SET_400.$OCTET
  IP_MASK=$IP_MASK
  IP_GW=$IP_GW
  SSHKEY=$(cat /root/.ssh/id_rsa.pub)
  SSHKEY2=$(cat /root/.ssh/authorized_keys | grep node)
}

function ssh_key_remove() {
  ssh-keygen -f "/root/.ssh/known_hosts" -R $IP_ADDR_200
  ssh-keygen -f "/root/.ssh/known_hosts" -R $IP_ADDR_400
}


function init_network_config_ubuntu() {
cat >$VMNAME/network-config <<EOF
ethernets:
    $INTERFACE_200:
        addresses:
        - $IP_ADDR_200/$IP_MASK
        dhcp4: false
        gateway4: $IP_GW
        match:
            macaddress: $MAC_ADDR_200
        nameservers:
            addresses:
            - $IP_DNS
        set-name: $INTERFACE_200
    $INTERFACE_400:
        addresses:
        - $IP_ADDR_400/$IP_MASK
        dhcp4: false
        match:
            macaddress: $MAC_ADDR_400
        set-name: $INTERFACE_400
version: 2
EOF
}


function  init_cloud_config() {
  touch $VMNAME/meta-data
  cloud-localds -v --network-config=$VMNAME/network-config $VMNAME/$VMNAME-seed.qcow2 $VMNAME/user-data $VMNAME/meta-data
}


function virt_install_one_zfs() {

virt-install --connect qemu:///system --accelerate --virt-type kvm --name $VMNAME --ram $MEMORY --vcpus=$CORE --cpu Westmere --os-type linux --os-variant $OSvariant \
 --boot hd  --disk path=/dev/zvol/clouds/$VMNAME,bus=scsi,discard=unmap \
 --disk path=$VMNAME/$VMNAME-seed.qcow2,device=disk --import \
 --controller type=scsi,model=virtio-scsi \
 --network network=cloud1,model=virtio,mac=$MAC_ADDR_200 \
 --network network=cloud2,model=virtio,mac=$MAC_ADDR_400 --noautoconsole

}


function virt_install_addon_disk_zfs() {

virt-install --connect qemu:///system --accelerate --virt-type kvm --name $VMNAME --ram $MEMORY --vcpus=$CORE --cpu Westmere --os-type linux --os-variant $OSvariant \
 --boot hd  --disk path=/dev/zvol/clouds/$VMNAME,bus=scsi,discard=unmap \
 --controller type=scsi,model=virtio-scsi \
 --disk path=/dev/zvol/clouds/$VMNAME-disk2,bus=scsi,discard=unmap \
 --disk path=/dev/zvol/clouds/$VMNAME-disk3,bus=scsi,discard=unmap \
 --disk path=/dev/zvol/clouds/$VMNAME-disk4,bus=scsi,discard=unmap \
 --disk path=$VMNAME/$VMNAME-seed.qcow2,device=disk --import \
 --network network=cloud1,model=virtio,mac=$MAC_ADDR_200 \
 --network network=cloud2,model=virtio,mac=$MAC_ADDR_400 --noautoconsole
}


function virt_install_addon_disk_file() {

# osinfo-query os
virt-install --connect qemu:///system --accelerate --virt-type kvm --name $VMNAME --ram $MEMORY --vcpus=$CORE --cpu Westmere --os-type linux --os-variant $OSvariant \
 --disk path=$VMNAME/$VMNAME.qcow2,device=disk \
 --disk path=$VMNAME/$VMNAME-disk2.qcow2,device=disk \
 --disk path=$VMNAME/$VMNAME-disk3.qcow2,device=disk \
 --disk path=$VMNAME/$VMNAME-disk4.qcow2,device=disk \
 --disk path=$VMNAME/$VMNAME-seed.qcow2,device=disk --import \
 --network network=cloud1,model=virtio,mac=$MAC_ADDR_200 \
 --network network=cloud2,model=virtio,mac=$MAC_ADDR_400 --noautoconsole
}


function virt_install_one_file() {
# osinfo-query os
virt-install --connect qemu:///system --accelerate --virt-type kvm --name $VMNAME --ram $MEMORY --vcpus=$CORE --cpu Westmere --os-type linux --os-variant $OSvariant \
 --disk path=$VMNAME/$VMNAME.qcow2,device=disk \
 --disk path=$VMNAME/$VMNAME-seed.qcow2,device=disk --import \
 --network network=cloud1,model=virtio,mac=$MAC_ADDR_200 \
 --network network=cloud2,model=virtio,mac=$MAC_ADDR_400 --noautoconsole
}



function finale_ubuntu() {
  PDSH_SSH_ARGS_APPEND="-q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null" pdsh -R ssh -w root@$IP_ADDR_200 "echo 'datasource_list: [ None ]' | sudo -s tee /etc/cloud/cloud.cfg.d/90_dpkg.cfg"
  PDSH_SSH_ARGS_APPEND="-q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null" pdsh -R ssh -w root@$IP_ADDR_200 "sudo apt-get purge -y cloud-init"
  PDSH_SSH_ARGS_APPEND="-q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null" pdsh -R ssh -w root@$IP_ADDR_200 "sudo rm -rf /etc/cloud/; sudo rm -rf /var/lib/cloud/"
  PDSH_SSH_ARGS_APPEND="-q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null" pdsh -R ssh -w root@$IP_ADDR_200 "poweroff"
}

function finale_debian() {
  PDSH_SSH_ARGS_APPEND="-q -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null" pdsh -R ssh -w root@$IP_SET.$OCTET "poweroff"
}
