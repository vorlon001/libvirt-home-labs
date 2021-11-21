function sleeps() {
  echo "sleep $1 sec."
  sleep $1
}



function isShutdown() {
    while [[  $(virsh domstate $VMNAME) != 'shut off' ]]; do
        echo "VM:" $(virsh domstate $VMNAME);
        echo "SSH:" $(nc -zw3 $IP_SET_200.$OCTET 22 && echo "open" || echo "close");
        sleep 5;
    done
    echo "VM $VMNAME is shutdown!"
}


function isSshEnable() {
    while [[  $(nc -zw3 $IP_SET_200.$OCTET 22 && echo "open" || echo "close") != 'open' ]]; do
        echo   "VM:" $(virsh domstate $VMNAME);
        echo "SSH:" $(nc -zw3 $IP_SET_200.$OCTET 22 && echo "open" || echo "close");
        sleep 5;
    done
    echo 'VM $VMNAME is ssh enable!'
}
