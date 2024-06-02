# rmbcmd

To invoke RMB commands on specific node or twin using the twinID or nodeID. For example

```sh
 ./rmbcmd -cmd 'zos.system.version' -twinID 16
```

or 

```
 ./rmbcmd -cmd 'zos.system.version' -nodeID 8
```


## usage

```
➜  rmbcmd git:(master) ✗ ./rmbcmd -h                                  
Usage of ./rmbcmd:
  -chainUrl string
        chain url (default "wss://tfchain.grid.tf/")
  -cmd string
        rmb cmd
  -gridProxyUrl string
        gridproxy url (default "https://gridproxy.grid.tf")
  -mnemonic string
        mnemonic
  -nodeID uint
        node id
  -relayUrl string
        relay url (default "wss://relay.grid.tf")
  -twinID uint
        node twin id
```

- You can provide `MNEMONIC` env variable
- You can't have twinID and nodeID at the same time

## Example

### Invoke by nodeID

Will invoke `zos.system.version` on node with ID 8

```
➜  rmbcmd git:(master) ✗ ./rmbcmd -cmd 'zos.system.version' -nodeID 8
https://gridproxy.grid.tf/nodes/8
{"level":"debug","url":"wss://tfchain.grid.tf/","time":"2024-06-02T20:13:27+03:00","message":"connecting"}
2024/06/02 20:13:27 Connecting to wss://tfchain.grid.tf/...
{"level":"debug","url":"wss://tfchain.grid.tf/","time":"2024-06-02T20:13:28+03:00","message":"connecting"}
2024/06/02 20:13:28 Connecting to wss://tfchain.grid.tf/...
{"level":"info","twin":4,"session":"test-client","time":"2024-06-02T20:13:28+03:00","message":"starting peer"}
{"level":"debug","url":"wss://relay.grid.tf","time":"2024-06-02T20:13:28+03:00","message":"connecting"}
========
{
  "zinit": "v0.2.11",
  "zos": "3.10.6"
}
========
```

### Invoke by twinID

```
➜  rmbcmd git:(master) ✗ ./rmbcmd -cmd 'zos.system.version' -twinID 16
{"level":"debug","url":"wss://tfchain.grid.tf/","time":"2024-06-02T20:13:50+03:00","message":"connecting"}
2024/06/02 20:13:50 Connecting to wss://tfchain.grid.tf/...
{"level":"debug","url":"wss://tfchain.grid.tf/","time":"2024-06-02T20:13:51+03:00","message":"connecting"}
2024/06/02 20:13:51 Connecting to wss://tfchain.grid.tf/...
{"level":"info","twin":4,"session":"test-client","time":"2024-06-02T20:13:52+03:00","message":"starting peer"}
{"level":"debug","url":"wss://relay.grid.tf","time":"2024-06-02T20:13:52+03:00","message":"connecting"}
========
{
  "zinit": "v0.2.11",
  "zos": "3.10.6"
}
========
```