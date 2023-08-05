# shelly-cert
Tool for certificate upload to a Shelly Gen2 device

## Usage example:

```
shelly-cert.exe -host 192.168.178.84 -file ca.crt -type PutUserCa
```

Allowed methods are PutUserCa, PutTLSClientCert, PutTLSClientKey. 
**Note:** Any previously uploaded certificate files will be deleted.