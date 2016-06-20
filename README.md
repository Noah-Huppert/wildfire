![flame](assets/flame.png)  

# wildfire
A decentralized gossip protocol written in golang.

# communication
Wildfire makes no assumptions about the communication method.  
Message acknowledgement is not expected out of a communication method.  
Thus UDP was chosen as the default communication method. Although UDP does not guarantee delivery, local network tests showed only 0.00517% (994829572/1000000000) of UDP packets sent were not received.  

One may provide a custom `ComProtocol` interface if a different communication method is desired.

---

*Image assets provided by [Icons8](https://icons8.com)*  
<img src="assets/flame.png" alt="flame by icons8" width="16" height="16" />
