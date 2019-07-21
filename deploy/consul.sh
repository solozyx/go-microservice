#!/usr/bin/env bash

#安装consul-1.5.1集群 3server-1client
#server1
consul agent -server -bootstrap-expect 3 -data-dir /opt/soft/consul_data -node=n1 -bind=192.168.174.134 -ui -config-dir /opt/soft/consul_config -rejoin -join 192.168.174.134 -client 0.0.0.0 &
#server之间replication端口 8300
firewall-cmd --zone=public --add-port=8300/tcp --permanent
#server与client通信端口 8301
firewall-cmd --zone=public --add-port=8301/tcp --permanent
#ui界面监听端口 8500
firewall-cmd --zone=public --add-port=8500/tcp --permanent
systemctl restart firewalld.service

#server2 加入 #server1
consul agent -server -bootstrap-expect 3 -data-dir /opt/soft/consul_data -node=n2 -bind=192.168.174.135 -ui -rejoin -join 192.168.174.134 &
firewall-cmd --zone=public --add-port=8300/tcp --permanent
firewall-cmd --zone=public --add-port=8301/tcp --permanent
firewall-cmd --zone=public --add-port=8500/tcp --permanent
systemctl restart firewalld.service

#server3 加入 #server1
consul agent -server -bootstrap-expect 3 -data-dir /opt/soft/consul_data -node=n3 -bind=192.168.174.136 -ui -rejoin -join 192.168.174.134 &
firewall-cmd --zone=public --add-port=8300/tcp --permanent
firewall-cmd --zone=public --add-port=8301/tcp --permanent
firewall-cmd --zone=public --add-port=8500/tcp --permanent
systemctl restart firewalld.service

#client
consul agent -data-dir /opt/soft/consul_data -node=n4 -bind=192.168.174.137 -config-dir /opt/soft/consul_config -rejoin -join 192.168.110.123 &

#查看集群状态
http://192.168.174.134:8500/ui/dc1/nodes
consul members