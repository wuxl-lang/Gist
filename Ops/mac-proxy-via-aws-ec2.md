## Mac Proxy via AWS EC2

> Create an EC2 instance in AWS and download key file, such as **proxy.pem**.

> Create a socket to listen to port on local side.

```
chmod 400 proxy.pem
ssh -qTfnN -D 2009 -i proxy.pem ec2@ec2-xxx.compute.amazonaws.com
```

> Set up proxy in terminal

```
export https_proxy=socks5://127.0.0.1:2009
```
