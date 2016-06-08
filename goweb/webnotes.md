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

### atomic and mutex
atomic and mutex. mutex is just an atomic boolean in typical implementation. only the processor can access a variable.

### HTTP
Protocol on top of TCP. Servers receive requests(Request line / headers)and send back responses (status line/ headers / content). Verbs (GET, POST, PUT, DELETE, HEAD(GET but no body, could be useful for video)). PUT vs POST, create and update, but they are similar POST when url unknown, PUT when known.
HTTP server just parses requests and sends a response

### ResponseWriter
Interface for http package
1. Header() - returns the header map that will be sent by write header
2. Write([]byte) (int, error) - writes to the connection as part of the http reply
3. WriteHeader(int) - Sends http response with a status code.

### Restful Web Applications
Web software architecture style for convenience when it comes to scalability of dynamic services. uses HTTP verbs to do build upon simple concepts. URL should be pretty/logical/meaningful
* ServeMux - server multiplexer, route signals based on conditions
* Julien Schmidt httprouter beats all the benchamarks, even servemux
* Restful API dont need SDKs!!

### Serving pictures
When you want to serve pictures on a http server
1. os.Open - io.Copy(res, myfile) - long kind of dumb way
2. http.ServeContent(res, req, f.Name(), fi.ModTime(), f) - Good way can give you diff time and some other useful stuff
3. http.ServeFile(res,req,"picname,jpg") - very convenient no hasle, directly chooses files
4. http.Handle("/", http.FileServer(http.Dir("."))) -  serves up all the files in the directory, all ready for embedding!

### Templates
* Embedding in html {{ stuff }}
* 



* * *
28
