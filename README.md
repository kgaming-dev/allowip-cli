# AllowIpAddr CLI

> Command Line Interface to allow an IP Addr or a list of IP Addresses in ufw rule set.


## Usage

```
> allowipaddr --file=txtfile.txt --comment="Comment in here" --allow-to-ip=x.x.x.x --allow-to-port=8483

--file | -f "filepath"           Path to a txtfile with an `ip address` per line or `ipaddress|comment per line`.
--ipaddr | -s "x.x.x.x"        A single IP Address to allow, this args will be discarded if a file is defined.
--comment | -c "Comment"                  Default to '' - Comment for ip(s) allowed. Use only if comment is not defined in the file
--allow-to-ip | -d    Default to `any`, destination Ip Address in our host
--allow-to-port | -p
```




