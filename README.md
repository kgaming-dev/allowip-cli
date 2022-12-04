# AllowIpAddr CLI

> Command Line Interface to allow an IP Addr or a list of IP Addresses in ufw rule set.


## Usage

```bash
> allowipaddr --file=txtfile.txt 
```

`txtfile.txt` is a standard text file contains list of ip address formatted `ipaddress:comment` for one ip per line. Comment can be omitted from the ip address line, it will executes `ufw allow` and creates the rule without `comment` argument.

Example content of the text file.

```text
173.245.48.0/20:Cloudflare Ip Address
103.21.244.0/22:Cloudflare Ip Address
103.22.200.0/22:Cloudflare Ip Address
103.31.4.0/22:Cloudflare Ip Address
141.101.64.0/18:Cloudflare Ip Address
108.162.192.0/18:Cloudflare Ip Address
190.93.240.0/20:Cloudflare Ip Address
188.114.96.0/20:Cloudflare Ip Address
197.234.240.0/22:Cloudflare Ip Address
198.41.128.0/17:Cloudflare Ip Address
162.158.0.0/15:Cloudflare Ip Address
104.16.0.0/13:Cloudflare Ip Address
104.24.0.0/14:Cloudflare Ip Address
172.64.0.0/13:Cloudflare Ip Address
131.0.72.0/22:Cloudflare Ip Address
```


