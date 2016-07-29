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
## interface
### Encoder
Responsible for transforming data into network or application readable formats.

- func Encode(data interface{}) string
- func Decode(data string) interface{}

### Transporter
Responsible for interacting with the chosen network stack and sending 
data through it. 

- func Send(data interface{}) Error
- func OnRecieve(data interface{}) Error

## struct
### Data
All data stored is derived from this struct. This provides a means to 
determine if information should be overwritten due to it being outdated. 

- lastUpdate long

### NodeAddr < Data
All information necessary to contact a node directly.  
Compressed string form: `nodeId`@`ip`:`port`

- nodeId string
- ip string
- port int

### MessageHeader
Contains identifying information pertaining to message.

- uuid string
- senderId string
- targetNodeId string 
	- Either a node id or an asterisks
	- Asterisks signifies the message is targeted at all nodes

---

*Image assets provided by [Icons8](https://icons8.com)*  
<img src="assets/flame.png" alt="flame by icons8" width="16" height="16" />
