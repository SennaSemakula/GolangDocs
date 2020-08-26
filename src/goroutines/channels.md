#Channels
- Channels allow us to send data from one Goroutine to another via a data transfer pipe
- Channels can be used to block the main goroutine until others have finished
- Channels by default are pointers 
- A channel can only support one data type. Not other types can be allowed in that channel
- You can create a channel with: ```make(chan <type>)```
- Sends and receives are blocked until sender and receiver are ready. When data is available in the channel it does not block. Only blocks when nothing is in the channel
**channel operations are blocked by nature. when data is written to a channel, goroutine is blocked until another goroutine reads from that channel**

Sending data through Channels
- ```channel <- value```
    - receives the value from the channel
Reading data from channel:
- ```<- channel```

**Deadlocks**
- deadlocks error happen when there is no available goroutine e.g. all of them are sleeping when the control is passed on.

**Channel buffering**
- by default channels are unbuffered meaning that they will only accept sends ```chan <-``` if there is a corresponding receive ```chan <-``` ready to receive the value