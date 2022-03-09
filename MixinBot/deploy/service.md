# Systemd

Deploy with systemd so your application would restart automatically when failed.

Example from [here](https://github.com/MixinNetwork/mixin/blob/master/config/systemd.service):

```service
[Unit]
Description=Mixin Network Kernel Daemon
After=network.target

[Service]
User=one
Type=simple
ExecStart=/home/one/bin/mixin kernel -dir /home/one/data/mixin -port 7239
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

Edit those fields with your application's path and details.

For tutorial of writing service file, visit [here](https://www.shubhamdipt.com/blog/how-to-create-a-systemd-service-in-linux/).