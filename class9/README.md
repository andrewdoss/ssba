```
multipass launch -d 20G -m 2G -n ssba3 bionic 
multipass exec ssba3 -- /bin/bash
```

Create copy of filesystem
```
sudo rsync -aAXv / --exclude={"/dev/*","/proc/*","/sys/*","/tmp/*","/run/*","/mnt/*","/media/*","/lost+found","/home/*"} /home/rootfs.img
```

Mount host development directory
```
multipass mount ${HOME}/Projects/ssba ssba3:/home/ubuntu/ssba 
```

Install Golang
```
sudo apt install golang-go
```

Run the container
```
sudo go run main.go run /bin/bash
```

References

[Cloning the filesystem](https://askubuntu.com/questions/1049930/how-to-copy-root-file-system-in-ubuntu)



- What is `clone` vs. `fork` vs. `exec`?
