
#!/bin/bash

set -x

NODENAME=$1


./Cobra LibVirt MachineDestroy --id $NODENAME
rm -R /cloud/KVM/$NODENAME*
