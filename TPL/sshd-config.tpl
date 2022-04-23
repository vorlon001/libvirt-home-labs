Include /etc/ssh/sshd_config.d/*.conf
Port 22
ListenAddress 0.0.0.0
SyslogFacility AUTH
LogLevel INFO
PermitRootLogin yes
StrictModes yes
MaxAuthTries 6
ChallengeResponseAuthentication no
UsePAM yes
AllowTcpForwarding yes
X11Forwarding no
PrintMotd no
AcceptEnv LANG LC_*
Subsystem sftp  /usr/lib/openssh/sftp-server
PasswordAuthentication yes
