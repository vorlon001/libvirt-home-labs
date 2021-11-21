# Create VM libvirt

![alt text](MadlHatterByTenniel.jpg)


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
