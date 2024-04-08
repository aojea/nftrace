# nftrace

Commodity tool to use nftables trace functionality, it just automates the steps described in

https://wiki.nftables.org/wiki-nftables/index.php/Ruleset_debug/tracing

```sh
A thin wrapper for nftables rules debugging
                Complete documentation is available at https://wiki.nftables.org/wiki-nftables/index.php/Ruleset_debug/tracing

Usage:
  nftrace [flags]
  nftrace [command]

Available Commands:
  add         add a trace (empty matches all traffic)
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        list the existing trace
  monitor     Monitor the nftables traces
  remove      remove the trace
  version     Print the version number

Flags:
  -h, --help   help for nftrace

Use "nftrace [command] --help" for more information about a command.
```


### How to use it

1. Set the expression you want to trace (nftables expression)

```sh
./nftrace add ip protocol tcp
```

2. You can check that it installs a new table/chain/rule

```sh
./nftrace list
table inet nftrace-table {
        comment "table for nftrace"
        chain nftrace-chain {
                comment "nftrace chain"
                type filter hook prerouting priority filter - 10; policy accept;
                ip protocol tcp meta nftrace set 1
        }
}
```

3. You can start monitoring the traffic now
```sh
 ./nftrace monitor
trace id 8974f599 inet nftrace-table nftrace-chain packet: iif "eth0" ether saddr 02:42:c0:a8:08:04 ether daddr 02:42:c0:a8:08:03 ip saddr 192.168.8.4 ip daddr 192.168.8.3 ip dscp cs0 ip ecn not-ect ip ttl 64 ip id 51749 ip protocol tcp ip length 141 tcp sport 6443 tcp dport 43062 tcp flags == 0x18 tcp window 2321
trace id 8974f599 inet nftrace-table nftrace-chain rule ip protocol tcp meta nftrace set 1 (verdict continue)
trace id 8974f599 inet nftrace-table nftrace-chain verdict continue
trace id 8974f599 inet nftrace-table nftrace-chain policy accept
trace id 8974f599 ip filter INPUT packet: iif "eth0" ether saddr 02:42:c0:a8:08:04 ether daddr 02:42:c0:a8:08:03 ip saddr 192.168.8.4 ip daddr 192.168.8.3 ip dscp cs0 ip ecn not-ect ip ttl 64 ip id 51749 ip length 141 tcp sport 6443 tcp dport 43062 tcp flags == 0x18 tcp window 2321
trace id 8974f599 ip filter INPUT rule  counter packets 5220005 bytes 2163351087 jump KUBE-NODEPORTS (verdict jump KUBE-NODEPORTS)
trace id 8974f599 ip filter KUBE-NODEPORTS verdict continue
trace id 8974f599 ip filter INPUT rule counter packets 5213526 bytes 2163060304 jump KUBE-FIREWALL (verdict jump KUBE-FIREWALL)
trace id 8974f599 ip filter KUBE-FIREWALL verdict continue
trace id 8974f599 ip filter INPUT verdict continue
trace id 8974f599 ip filter INPUT policy accept
trace id 03e7a3e4 inet nftrace-table nftrace-chain packet: iif "eth0" ether saddr 02:42:c0:a8:08:04 ether daddr 02:42:c0:a8:08:03 ip saddr 192.168.8.4 ip daddr 192.168.8.3 ip dscp cs0 ip ecn not-ect ip ttl 64 ip id 51750 ip protocol tcp ip length 7292 tcp sport 6443 tcp dport 43062 tcp flags == 0x18 tcp window 2321
```

4. Once you are done don't forget to remove the trace
  
```sh
./nftrace remove

```sh
./nftrace list
no traces active
```





