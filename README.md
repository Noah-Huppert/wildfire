#### Project Status
Wishful thinking  
Probably will never complete  
Other projects have priority for the forseable future

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
		# the target_id == "*" then the recieving node will be sure 
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
		# Body will have status of node, how detailed has not been 
		# decided. Could be as simple as "OK" or "NOT OK" or a detailed 
		# report could be returned
		

NodeAddr:
	id String
	host String
	port Int >= 1 && <= 65535
			
```

Cluster queries:
When a node receives a message and it finds a reference for a piece of data it 
was unaware of it submits a cluster query with the URI of the data it is 
missing. It then receives all the answers and uses the most up to date answer.

---

*Image assets provided by [Icons8](https://icons8.com)*  
<img src="assets/flame.png" alt="flame by icons8" width="16" height="16" />
