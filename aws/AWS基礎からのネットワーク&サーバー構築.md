# AWS基礎からのネットワーク&サーバー構築

## Apache HTTP Server を インストールする

## SSH で AWS のインスタンスに接続

`$ ssh -i my-key.pem ec2-user@35.72.139.60`

## install Apach

`$ sudo yum -y install httpd`

-y means that user will not check the instalation detail and install it instantly

## start Apach

`$ sudo service httpd start`

## configurate Apach will be started when the server starts

`$ sudo chkconfig httpd on`

## check if the configulation has been set properly

`$ sudo chkconfig --list httpd`

## check processes

`$ ps -ax | grep httpd`
481 ? Ss 0:00 /usr/sbin/httpd -DFOREGROUND
482 ? Sl 0:00 /usr/sbin/httpd -DFOREGROUND
483 ? Sl 0:00 /usr/sbin/httpd -DFOREGROUND
484 ? Sl 0:00 /usr/sbin/httpd -DFOREGROUND
485 ? Sl 0:00 /usr/sbin/httpd -DFOREGROUND
486 ? Sl 0:00 /usr/sbin/httpd -DFOREGROUND
588 pts/5 S+ 0:00 grep --color=auto httpd

## check network status

`$ sudo lsof -i -n -P`

COMMAND PID USER FD TYPE DEVICE SIZE/OFF NODE NAME
sshd 411 root 3u IPv4 55148 0t0 TCP 10.0.1.10:22->106.154.132.161:49369 (ESTABLISHED)
sshd 429 ec2-user 3u IPv4 55148 0t0 TCP 10.0.1.10:22->106.154.132.161:49369 (ESTABLISHED)
httpd 481 root 4u IPv6 55574 0t0 TCP _:80 (LISTEN)
httpd 482 apache 4u IPv6 55574 0t0 TCP _:80 (LISTEN)
httpd 483 apache 4u IPv6 55574 0t0 TCP _:80 (LISTEN)
httpd 484 apache 4u IPv6 55574 0t0 TCP _:80 (LISTEN)
httpd 485 apache 4u IPv6 55574 0t0 TCP _:80 (LISTEN)
httpd 486 apache 4u IPv6 55574 0t0 TCP _:80 (LISTEN)
chronyd 2498 chrony 5u IPv4 15340 0t0 UDP 127.0.0.1:323
chronyd 2498 chrony 6u IPv6 15341 0t0 UDP [::1]:323
rpcbind 2509 rpc 6u IPv4 15586 0t0 UDP _:111
rpcbind 2509 rpc 7u IPv4 15649 0t0 UDP _:981
rpcbind 2509 rpc 8u IPv4 15650 0t0 TCP _:111 (LISTEN)
rpcbind 2509 rpc 9u IPv6 15651 0t0 UDP _:111
rpcbind 2509 rpc 10u IPv6 15652 0t0 UDP _:981
rpcbind 2509 rpc 11u IPv6 15653 0t0 TCP _:111 (LISTEN)
dhclient 2724 root 6u IPv4 16142 0t0 UDP _:68
dhclient 2815 root 5u IPv6 16430 0t0 UDP [fe80::c1f:daff:fe3b:e493]:546
master 2942 root 13u IPv4 17144 0t0 TCP 127.0.0.1:25 (LISTEN)
sshd 3187 root 3u IPv4 19241 0t0 TCP _:22 (LISTEN)
sshd 3187 root 4u IPv6 19250 0t0 TCP \*:22 (LISTEN)
sshd 3269 root 3u IPv4 20642 0t0 TCP 10.0.1.10:22->106.154.138.58:64877 (ESTABLISHED)
sshd 3287 ec2-user 3u IPv4 20642 0t0 TCP 10.0.1.10:22->106.154.138.58:64877 (ESTABLISHED)
sshd 3331 root 3u IPv4 21212 0t0 TCP 10.0.1.10:22->106.154.138.58:64970 (ESTABLISHED)
sshd 3349 ec2-user 3u IPv4 21212 0t0 TCP 10.0.1.10:22->106.154.138.58:64970 (ESTABLISHED)
sshd 3450 root 3u IPv4 21990 0t0 TCP 10.0.1.10:22->106.154.138.58:65243 (ESTABLISHED)
sshd 3468 ec2-user 3u IPv4 21990 0t0 TCP 10.0.1.10:22->106.154.138.58:65243 (ESTABLISHED)
sshd 3600 root 3u IPv4 23746 0t0 TCP 10.0.1.10:22->106.154.138.58:65495 (ESTABLISHED)
sshd 3618 ec2-user 3u IPv4 23746 0t0 TCP 10.0.1.10:22->106.154.138.58:65495 (ESTABLISHED)
sshd 32700 root 3u IPv4 54323 0t0 TCP 10.0.1.10:22->106.154.138.58:49249 (ESTABLISHED)
sshd 32718 ec2-user 3u IPv4 54323 0t0 TCP 10.0.1.10:22->106.154.138.58:49249 (ESTABLISHED)

## check DNS server

`$ nslookup ec2-35-72-139-60.ap-northeast-1.compute.amazonaws.com`
