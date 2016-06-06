# Notes on Web Programming in GO

### Nested dolls that map into each other
* HTTP
* TCP - Transmission Control Protocol
* IP

### HTTP requests
Top network abstraction layer
Domains map to IP addresses - Translated by DNS routers
* $ ping nicolasfeller.com -- translated to 199.34.228.55
1. Type in to field nicolasfeller.com
2. request goes to dns server
3. DNS server looks up IP address
4. IP Address gets the information and sends it to the requester
* $ dig google.com -- lists IP addresses associated with url
* $ whois -- finds out who owns a domain name

### HTTP headers
see what the response header looks like, check the developer tools on a webpage -- Network/Headers

### curl
* $ curl nicolasfeller.com -- gives back the html

### TCP
TCP is the layer below http, tcp provides reliable connection and puts the packages in order. Sits between the application program and the Internet protocol. Provides host-host connectivity at the transport layer of the Internet. Essentially and abstraction that handles the handshaking.
* wireshark can help you manage and view you TCP activity.
* traceroute - shows route to a website

### Scanner
Not only used for networking, but scanner is essentially a convenient way to read data. It has many methods associated with it such as

### kill ports that busy
$ fuser -k 9000/tcp

### Web Socket
Protocol providing full-duplex communication channel over a TCP connection. To implement web browsers and web servers (client server applications). Independent TCP-based protocol that only uses HTTP to interpret the handshake (http servers only see an upgrade request). Javascript thing ;)
Its like an ongoing conversation - the client does not need to solicit the service. Very good for chat stuff. bi-directional flow between client and server. done over port 80. websockets are supported by all major browsers.

### package strings
Alot of the stuff that is passed between connections if just UTF-8 encoded strings. So we need to play with this stuff a lot!


* * *
28
