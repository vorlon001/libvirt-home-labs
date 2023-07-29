<disk type="file" device="disk">
        <driver name='qemu' type='qcow2'/>
        <source file="{{.Config.VMPATH}}/{{.VMNAME}}/{{.VMNAME}}-disk{{.DISKID}}.qcow2"></source>
        <target dev='{{.DISKID}}' bus='scsi'/>
</disk>
