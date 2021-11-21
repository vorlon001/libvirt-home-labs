# Create VM libvirt

![alt text](MadlHatterByTenniel.jpg)


###  this script creates ubuntu 20.04, 21.04, 21.10 , 22.10, debian 10, 11

INIT

need CPU KVM-OK, INTEL-TV

```shell
sudo apt -y install qemu-kvm libvirt-daemon-system libvirt-daemon virtinst bridge-utils libosinfo-bin libguestfs-tools virt-top
sudo apt-get install -y cloud-image-utils pdsh
```


```shell
mkdir IMAGES
git clone https://github.com/vorlon001/Mad-Hatter-or-The-Rabbit-Hole
mv Mad-Hatter-or-The-Rabbit-Hole KVM

cd KVM
cd run.sh

```


IMAGES UBUNTU

```

-rw-r--r--  1 root         root  583687680 Jul 29 12:17 debian-10-openstack-amd64.qcow2
-rw-r--r--  1 libvirt-qemu kvm   325246976 Aug 15 20:31 debian-11-generic-amd64-20210814-734.qcow2
-rw-r--r--  1 root         root  325224960 Jul 29 12:17 debian-11-generic-amd64-daily-20210608-662.qcow2
-rw-r--r--  1 root         root  326434816 Jul 29 12:17 debian-11-generic-amd64-daily-20210621-680.qcow2
-rw-r--r--  1 root         root  325093376 Jul 29 12:17 debian-11-generic-amd64-daily-20210728-717.qcow2

-rw-r--r--  1 root         root  567148544 Oct  1 06:26 focal-server-cloudimg-amd64.img

-rw-r--r--  1 libvirt-qemu kvm   585302016 Sep 28 13:26 hirsute-server-cloudimg-amd64.img
-rw-r--r--  1 libvirt-qemu kvm   568786944 Oct 14 18:02 impish-server-cloudimg-amd64.img
-rw-r--r--  1 libvirt-qemu kvm   566165504 Oct 20 07:48 jammy-server-cloudimg-amd64.img

```

### Need two network: cloud1, cloud2
