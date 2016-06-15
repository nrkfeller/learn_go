# MQTT
Open, Industry agnostic. Publish subscribe (one to many) light weight m to m communication protocol. For constrained noetworks on constrained systems. Two side pushing. No polling, timely. Always connected model. session aware (we know when devices go online) lets you know when device breaks/goes off. MQTT over web sockets (two way). deliver the message once and only once (and at least once) - quality of service. Battery frugal. json, xml even bit level encoding. can push about 100x more messages than https as well as guaranteed no loss.
 * port 1883

### Flow
1. Publisher publishes a message on a topic.
2. Message goes into a broker/server
3. Consumers then connect to a broker
4. Whenever the broker gets a message, it pushes it to the consumers

### Namespaces
* Typical data grab         : [country]/[region]/[town]/[address]/energyconsumption
* Wildcard all in town      : [country]/[region]/[town]/+]/energyconsumption
* Wildcard all from address :[country]/[region]/[town]/[address]/#

### Durable vs Non-Durable
* Durable : broker will store your messages when disconnected and give them back when you reconnect
* Non-Durable : Only get messages while you are connected (loose subscription when disconnected)
* The broker can also store the last state and send the subscriber the last known state upon subscription.

### 
