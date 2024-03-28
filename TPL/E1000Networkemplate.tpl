<interface type='network'>
      <source network='cloud_sw1'/>
      <model type='e1000'/>
      <driver name='vhost' txmode='iothread' ioeventfd='on' event_idx='off' queues='5'/>
      <mac address='{{.MacAddress}}'/>
      <address type='pci' domain='0x0000' bus='0x03' slot='{{.Slot}}' function='0x0' multifunction='on'/>
</interface>
