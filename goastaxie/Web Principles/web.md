# Principles of the web
### URL flow
1. Write url request
2. Browser program sends host part of url to DNS server to retrieve host IP - mapped IP to name
3. Asks IP to make TCP connection
4. Make HTTP request through TCP connection to host server
5. Host server sends HTTP response in form of HTML
6. Browser builds page and disconnects from server

### HTTP Basics
* HTTP is based on TCP
* Request/Response model
* Client must connect - Servers cannot connect pro-actively
* HTTP is stateless - servers does not know anything about relationship of connections
* Cookies are needed to maintain states

*Request Package*
1. Request line
2. Request header
* blank line
3. Body

*Response Package*
1. Status - 200 OK
2. Web server
3. Date
4. Content type
5. Transfer-Encoding - send in fragments or not
6. Keep connection or not
7. Length
* blank line
8. Body
