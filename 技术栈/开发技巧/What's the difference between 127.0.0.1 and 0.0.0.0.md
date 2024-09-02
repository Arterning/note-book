### What's the difference between 127.0.0.1 and 0.0.0.0?

- `127.0.0.1` is the loopback address (also known as localhost).
    
- `0.0.0.0` is commonly used as a non-routable meta-address used to designate an invalid, unknown or non applicable target (a no particular address placeholder). However this is non-standard and possibly in conflict with [RFC 1122](https://rfc-editor.org/rfc/rfc1122#page-30).
    

In the context of a route entry, it usually means the default route.

In the context of servers, 0.0.0.0 means "all IPv4 addresses on the local machine". If a host has two IP addresses, 192.168.1.1 and 10.1.2.1, and a server running on the host listens on 0.0.0.0, it will be reachable at both of those IPs.