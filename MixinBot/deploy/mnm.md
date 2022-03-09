# mnm

[MNM](https://mnm.sh) is a Mixin Messenger notifier Cedric Fung has built. It will notify you through Mixin Messenger when your application crash. It's really useful when you don't want to use Systemd to manage your service.


I found a way to run mnm in the background with log.

`setsid mnm run "./main" >| "log" &`
