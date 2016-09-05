Project Status: Short Term Hold (Apollo & Migrator are top priority)  
![flame](assets/flame.png)  

# wildfire
A leaderless gossip protocol written in golang.

# Transportation layer
By default Wildfire uses TCP to transport messages between nodes. 
TCP provides several desirable features such as congestion prevention, 
message delivery acknowledgment and many others. As an added bonus TCP 
does not require any additional services to connect to nodes on 
external networks (Such as STUN servers when using UDP).

One may provide a custom `Transporter` interface if a different communication method is desired.

# spec
```
Message:
	[
		Header,
		Body
	]

Header:
	uuid String
	sender_id String
	intermediary_ids [(String = id,Date = contacted date)]  
		# When messages must be forwarded between nodes to reach their 
		# final destination, each intermediary appends a tupple onto 
		# this array. This allows for heartbeat requirements to be 
		# fullfilled  by looking at message headers.
		#
		# This also allows for better duplicate message control. If 
		# the sender_id == "*" then the recieving node will be sure 
		# not to send the message back out to any node who has handled 
		# the message. The receiving node can also keep updating this 
		# list internally for future message forwarding 
	sender_protcl_ver String
	target_id String = id | "*"
	sent_at Date
	event String

Body:
	[[Event Data]]	

Event Data:
When Header.event
	= "node.joined"
		# Sent when a node joins.
		# To join a cluster a new node must simply send this event to 
		# a known node. The known node will then spread knowledge of 
		# the new node to the rest of the cluster
		NodeAddr
	= "node.left"
		NodeAddr
	= "health.heartbeat.ping"
		# Sent after a configurable timeout has been surpassed since 
		# last contact with a node.
	= "health.heartbeat.response"
		

NodeAddr:
	id String
	host String
	port Int >= 1 && <= 65535
			
```

Cluster queries:
When a node receives a message and it finds a reference for a piece of data it 
was unaware of it submits a cluster query with the URI of the data it is 
missing. It then receives all the answers and uses the most up to date answer.

----
## interface
### Encoder
Responsible for transforming data into network or application readable formats.

- func Encode(data interface{}) string
- func Decode(data string) interface{}
//typedef Variable Map<String, Encoder.Field>

### Transporter
Responsible for interacting with the chosen network stack and sending 
data through it. 

- func Send(data interface{}) Error
- func OnRecieve(data interface{}) Error

## struct
### Program data
#### MessageHeader
Contains identifying information pertaining to message.

- uuid string
- senderId string
- senderWldfrVer string
	- Sender wildfire version
- targetNodeId string 
	- Either a node id or an asterisks
	- Asterisks signifies the message is targeted at all nodes

#### MessageBody <PayloadType>
The content of the message itself.  

- intent string
- payload PayloadType

#### Message
Data structure in which all information will be transferred between nodes.

- header MessageHeader
- body MessageBody


### Sendable data
#### Data
All data stored is derived from this struct. This provides a means to 
determine if information should be overwritten due to it being outdated. 

Type: `Template("$lastUpdate") | Variable`

- lastUpdate long

#### NodeAddr < Data
All information necessary to contact a node directly.  

Short template: $nodeId@$ip:$port

- nodeId string
- ip string
- port int


---

*Image assets provided by [Icons8](https://icons8.com)*  
<img src="assets/flame.png" alt="flame by icons8" width="16" height="16" />
