# Alerts generation from test queries

This report captures the unit test queries signals generation coverage.
Here you can learn what queries are supported.

## Table of contents
   1. [Rules with no signals (2)](#rules-with-no-signals-2)
   1. [Rules with the correct signals (48)](#rules-with-the-correct-signals-48)

## Rules with no signals (2)

### Rule 029

Branch count: 1  
Document count: 1  
Index: geneve-ut-029

```python
network where destination.ip == "127.0.0.1" and _meta.index == "test"
```

```python
[{'destination': {'ip': '127.0.0.1'}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 045

Branch count: 1  
Document count: 2  
Index: geneve-ut-045

```python
sequence
        [network where _meta.index == "index1" and source.ip != null] by source.ip
        [network where _meta.index == "index2" and destination.ip != null ] by destination.ip
```

```python
[{"_meta": {"index": "index1"}, "event": {"category": ["network"]}, "source": {"ip": "167.158.207.19"}},
 {"_meta": {"index": "index2"}, "event": {"category": ["network"]}, "destination": {"ip": "167.158.207.19"}}]
```



## Rules with the correct signals (48)

### Rule 000

Branch count: 1  
Document count: 1  
Index: geneve-ut-000

```python
any where true
```

```python
[{'@timestamp': 0}]
```



### Rule 001

Branch count: 1  
Document count: 1  
Index: geneve-ut-001

```python
any where not false
```

```python
[{'@timestamp': 0}]
```



### Rule 002

Branch count: 1  
Document count: 1  
Index: geneve-ut-002

```python
any where not (true and false)
```

```python
[{'@timestamp': 0}]
```



### Rule 003

Branch count: 1  
Document count: 1  
Index: geneve-ut-003

```python
any where not (false or false)
```

```python
[{'@timestamp': 0}]
```



### Rule 004

Branch count: 1  
Document count: 1  
Index: geneve-ut-004

```python
network where source.port > 512 and source.port < 1024
```

```python
[{'source': {'port': 971}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 005

Branch count: 1  
Document count: 1  
Index: geneve-ut-005

```python
network where not (source.port < 512 or source.port > 1024)
```

```python
[{'source': {'port': 999}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 006

Branch count: 1  
Document count: 1  
Index: geneve-ut-006

```python
network where destination.port not in (80, 443)
```

```python
[{'destination': {'port': 65449}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 007

Branch count: 1  
Document count: 1  
Index: geneve-ut-007

```python
network where not destination.port in (80, 443)
```

```python
[{'destination': {'port': 65449}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 008

Branch count: 1  
Document count: 1  
Index: geneve-ut-008

```python
network where destination.port == 22 and destination.port in (80, 443) or destination.port == 25
```

```python
[{'destination': {'port': 25}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 009

Branch count: 1  
Document count: 1  
Index: geneve-ut-009

```python
process where process.name == "regsvr32.exe"
```

```python
[{'process': {'name': 'regsvr32.exe'}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 010

Branch count: 1  
Document count: 1  
Index: geneve-ut-010

```python
process where process.name != "regsvr32.exe"
```

```python
[{'process': {'name': 'ZFy'}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 011

Branch count: 1  
Document count: 1  
Index: geneve-ut-011

```python
process where process.pid != 0
```

```python
[{'process': {'pid': 4289255490}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 012

Branch count: 1  
Document count: 1  
Index: geneve-ut-012

```python
process where process.pid >= 0
```

```python
[{'process': {'pid': 4289255490}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 013

Branch count: 1  
Document count: 1  
Index: geneve-ut-013

```python
process where process.pid > 0
```

```python
[{'process': {'pid': 4289255490}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 014

Branch count: 1  
Document count: 1  
Index: geneve-ut-014

```python
process where process.code_signature.exists == true
```

```python
[{'process': {'code_signature': {'exists': True}}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 015

Branch count: 1  
Document count: 1  
Index: geneve-ut-015

```python
process where process.code_signature.exists != true
```

```python
[{'process': {'code_signature': {'exists': False}}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 016

Branch count: 1  
Document count: 1  
Index: geneve-ut-016

```python
any where network.protocol == "some protocol"
```

```python
[{'network': {'protocol': 'some protocol'}, '@timestamp': 0}]
```



### Rule 017

Branch count: 1  
Document count: 1  
Index: geneve-ut-017

```python
any where process.pid == null
```

```python
[{'@timestamp': 0}]
```



### Rule 018

Branch count: 1  
Document count: 1  
Index: geneve-ut-018

```python
any where not process.pid != null
```

```python
[{'@timestamp': 0}]
```



### Rule 019

Branch count: 1  
Document count: 1  
Index: geneve-ut-019

```python
any where process.pid != null
```

```python
[{'process': {'pid': 4289255490}, '@timestamp': 0}]
```



### Rule 020

Branch count: 1  
Document count: 1  
Index: geneve-ut-020

```python
any where not process.pid == null
```

```python
[{'process': {'pid': 4289255490}, '@timestamp': 0}]
```



### Rule 021

Branch count: 1  
Document count: 1  
Index: geneve-ut-021

```python
process where process.name == "regsvr32.exe" and process.parent.name == "cmd.exe"
```

```python
[{'process': {'name': 'regsvr32.exe', 'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 022

Branch count: 1  
Document count: 1  
Index: geneve-ut-022

```python
process where process.name : ("*.EXE", "*.DLL")
```

```python
[{'process': {'name': 'XIUtkNI.EXE'}, 'event': {'category': ['process']}, '@timestamp': 0}]
```



### Rule 023

Branch count: 1  
Document count: 1  
Index: geneve-ut-023

```python
network where destination.ip == "127.0.0.1"
```

```python
[{'destination': {'ip': '127.0.0.1'}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 024

Branch count: 1  
Document count: 1  
Index: geneve-ut-024

```python
network where cidrMatch(destination.ip, "10.0.0.0/8", "192.168.0.0/16")
```

```python
[{'destination': {'ip': '192.168.214.62'}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 025

Branch count: 1  
Document count: 1  
Index: geneve-ut-025

```python
network where not cidrMatch(destination.ip, "10.0.0.0/8", "192.168.0.0/16")
```

```python
[{'destination': {'ip': '107.31.65.130'}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 026

Branch count: 1  
Document count: 1  
Index: geneve-ut-026

```python
network where destination.ip == "::1"
```

```python
[{'destination': {'ip': '::1'}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 027

Branch count: 1  
Document count: 1  
Index: geneve-ut-027

```python
network where destination.ip == "822e::/16"
```

```python
[{'destination': {'ip': '822e:c14a:e6ea:94e4:e5ac:b58c:1b43:3a53'}, 'event': {'category': ['network']}, '@timestamp': 0}]
```



### Rule 028

Branch count: 1  
Document count: 1  
Index: geneve-ut-028

```python
event.category:network and destination.ip:"822e::/16"
```

```python
[{'event': {'category': ['network']}, 'destination': {'ip': '822e:c14a:e6ea:94e4:e5ac:b58c:1b43:3a53'}, '@timestamp': 0}]
```



### Rule 030

Branch count: 2  
Document count: 2  
Index: geneve-ut-030

```python
network where not (source.port > 512 and source.port < 1024)
```

```python
[{'source': {'port': 488}, 'event': {'category': ['network']}, '@timestamp': 0},
 {'source': {'port': 28447}, 'event': {'category': ['network']}, '@timestamp': 1}]
```



### Rule 031

Branch count: 2  
Document count: 2  
Index: geneve-ut-031

```python
network where source.port > 512 or source.port < 1024
```

```python
[{'source': {'port': 59173}, 'event': {'category': ['network']}, '@timestamp': 0},
 {'source': {'port': 628}, 'event': {'category': ['network']}, '@timestamp': 1}]
```



### Rule 032

Branch count: 2  
Document count: 2  
Index: geneve-ut-032

```python
network where source.port < 2000 and (source.port > 512 or source.port > 1024)
```

```python
[{'source': {'port': 1768}, 'event': {'category': ['network']}, '@timestamp': 0},
 {'source': {'port': 1915}, 'event': {'category': ['network']}, '@timestamp': 1}]
```



### Rule 033

Branch count: 2  
Document count: 2  
Index: geneve-ut-033

```python
network where (source.port > 512 or source.port > 1024) and source.port < 2000
```

```python
[{'source': {'port': 1768}, 'event': {'category': ['network']}, '@timestamp': 0},
 {'source': {'port': 1915}, 'event': {'category': ['network']}, '@timestamp': 1}]
```



### Rule 034

Branch count: 4  
Document count: 4  
Index: geneve-ut-034

```python
network where (source.port > 1024 or source.port < 2000) and (source.port < 4000 or source.port > 512)
```

```python
[{'source': {'port': 3536}, 'event': {'category': ['network']}, '@timestamp': 0},
 {'source': {'port': 58008}, 'event': {'category': ['network']}, '@timestamp': 1},
 {'source': {'port': 975}, 'event': {'category': ['network']}, '@timestamp': 2},
 {'source': {'port': 1369}, 'event': {'category': ['network']}, '@timestamp': 3}]
```



### Rule 035

Branch count: 2  
Document count: 2  
Index: geneve-ut-035

```python
network where destination.port in (80, 443)
```

```python
[{'destination': {'port': 80}, 'event': {'category': ['network']}, '@timestamp': 0},
 {'destination': {'port': 443}, 'event': {'category': ['network']}, '@timestamp': 1}]
```



### Rule 036

Branch count: 2  
Document count: 2  
Index: geneve-ut-036

```python
process where process.name == "regsvr32.exe" or process.parent.name == "cmd.exe"
```

```python
[{'process': {'name': 'regsvr32.exe'}, 'event': {'category': ['process']}, '@timestamp': 0},
 {'process': {'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, '@timestamp': 1}]
```



### Rule 037

Branch count: 3  
Document count: 3  
Index: geneve-ut-037

```python
process where process.name == "regsvr32.exe" or process.name == "cmd.exe" or process.name == "powershell.exe"
```

```python
[{'process': {'name': 'regsvr32.exe'}, 'event': {'category': ['process']}, '@timestamp': 0},
 {'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 1},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, '@timestamp': 2}]
```



### Rule 038

Branch count: 3  
Document count: 3  
Index: geneve-ut-038

```python
process where process.name in ("regsvr32.exe", "cmd.exe", "powershell.exe")
```

```python
[{'process': {'name': 'regsvr32.exe'}, 'event': {'category': ['process']}, '@timestamp': 0},
 {'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 1},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, '@timestamp': 2}]
```



### Rule 039

Branch count: 3  
Document count: 3  
Index: geneve-ut-039

```python
process where process.name in ("regsvr32.exe", "cmd.exe") or process.name == "powershell.exe"
```

```python
[{'process': {'name': 'regsvr32.exe'}, 'event': {'category': ['process']}, '@timestamp': 0},
 {'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 1},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, '@timestamp': 2}]
```



### Rule 040

Branch count: 2  
Document count: 2  
Index: geneve-ut-040

```python
process where event.type in ("start", "process_started") and process.args : "dump-keychain" and process.args : "-d"
```

```python
[{'event': {'type': ['start'], 'category': ['process']}, 'process': {'args': ['dump-keychain', '-d']}, '@timestamp': 0},
 {'event': {'type': ['process_started'], 'category': ['process']}, 'process': {'args': ['dump-keychain', '-d']}, '@timestamp': 1}]
```



### Rule 041

Branch count: 2  
Document count: 2  
Index: geneve-ut-041

```python
event.type:(start or process_started) and (process.args:"dump-keychain" and process.args:"-d")
```

```python
[{'event': {'type': ['start']}, 'process': {'args': ['dump-keychain', '-d']}, '@timestamp': 0},
 {'event': {'type': ['process_started']}, 'process': {'args': ['dump-keychain', '-d']}, '@timestamp': 1}]
```



### Rule 042

Branch count: 1  
Document count: 2  
Index: geneve-ut-042

```python
sequence
        [process where process.name : "cmd.exe"]
        [process where process.parent.name : "cmd.exe"]
```

```python
[{'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 0},
 {'process': {'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, '@timestamp': 1}]
```



### Rule 043

Branch count: 1  
Document count: 2  
Index: geneve-ut-043

```python
sequence by user.id
        [process where process.name : "cmd.exe"]
        [process where process.parent.name : "cmd.exe"]
```

```python
[{'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'ZFy'}, '@timestamp': 0},
 {'process': {'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, 'user': {'id': 'ZFy'}, '@timestamp': 1}]
```



### Rule 044

Branch count: 1  
Document count: 2  
Index: geneve-ut-044

```python
sequence
        [process where process.name : "cmd.exe"] by user.id
        [process where process.parent.name : "cmd.exe"] by user.name
```

```python
[{'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'ZFy'}, '@timestamp': 0},
 {'process': {'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, 'user': {'name': 'ZFy'}, '@timestamp': 1}]
```



### Rule 046

Branch count: 2  
Document count: 4  
Index: geneve-ut-046

```python
sequence
        [process where process.name : "cmd.exe"]
        [process where process.parent.name : "cmd.exe" or process.name : "powershell.exe"]
```

```python
[{'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 0},
 {'process': {'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, '@timestamp': 1},
 {'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 2},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, '@timestamp': 3}]
```



### Rule 047

Branch count: 2  
Document count: 4  
Index: geneve-ut-047

```python
sequence by user.id
        [process where process.name : "cmd.exe"]
        [process where process.parent.name : "cmd.exe" or process.name : "powershell.exe"]
```

```python
[{'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'ZFy'}, '@timestamp': 0},
 {'process': {'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, 'user': {'id': 'ZFy'}, '@timestamp': 1},
 {'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'XIU'}, '@timestamp': 2},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'XIU'}, '@timestamp': 3}]
```



### Rule 048

Branch count: 4  
Document count: 8  
Index: geneve-ut-048

```python
sequence
        [process where process.name in ("cmd.exe", "powershell.exe")] by process.name
        [process where process.name in ("cmd.exe", "powershell.exe")] by process.parent.name
```

```python
[{'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 0},
 {'process': {'name': 'cmd.exe', 'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, '@timestamp': 1},
 {'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, '@timestamp': 2},
 {'process': {'name': 'powershell.exe', 'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, '@timestamp': 3},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, '@timestamp': 4},
 {'process': {'name': 'cmd.exe', 'parent': {'name': 'powershell.exe'}}, 'event': {'category': ['process']}, '@timestamp': 5},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, '@timestamp': 6},
 {'process': {'name': 'powershell.exe', 'parent': {'name': 'powershell.exe'}}, 'event': {'category': ['process']}, '@timestamp': 7}]
```



### Rule 049

Branch count: 4  
Document count: 8  
Index: geneve-ut-049

```python
sequence by user.id
        [process where process.name in ("cmd.exe", "powershell.exe")] by process.name
        [process where process.name in ("cmd.exe", "powershell.exe")] by process.parent.name
```

```python
[{'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'ZFy'}, '@timestamp': 0},
 {'process': {'name': 'cmd.exe', 'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, 'user': {'id': 'ZFy'}, '@timestamp': 1},
 {'process': {'name': 'cmd.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'XIU'}, '@timestamp': 2},
 {'process': {'name': 'powershell.exe', 'parent': {'name': 'cmd.exe'}}, 'event': {'category': ['process']}, 'user': {'id': 'XIU'}, '@timestamp': 3},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'tkN'}, '@timestamp': 4},
 {'process': {'name': 'cmd.exe', 'parent': {'name': 'powershell.exe'}}, 'event': {'category': ['process']}, 'user': {'id': 'tkN'}, '@timestamp': 5},
 {'process': {'name': 'powershell.exe'}, 'event': {'category': ['process']}, 'user': {'id': 'Ioi'}, '@timestamp': 6},
 {'process': {'name': 'powershell.exe', 'parent': {'name': 'powershell.exe'}}, 'event': {'category': ['process']}, 'user': {'id': 'Ioi'}, '@timestamp': 7}]
```
