iptables -P INPUT ACCEPT
iptables -F
# iptables -A INPUT -i lo -j ACCEPT
# iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
iptables -A INPUT -p tcp --dport 22 -j ACCEPT
iptables -t nat -A PREROUTING -j DNAT \ --to 192.168.0.4
# iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -P OUTPUT ACCEPT
iptables -L -v

iptables -t nat -A PREROUTING -j DNAT --to-destination 192.168.0.4


i2c-dev
batman-adv
r8712u