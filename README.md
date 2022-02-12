# fileserver

Serve your files on http

## How to install

```
sudo curl -L https://github.com/navilg/fileserver/raw/master/bin/fileserver -o /usr/local/bin/fileserver
sudo chmod +x /usr/local/bin/fileserver
```

## Usage of fileserver

  -path string

        Path to serve (default "Current working directory")

  -port string

        Port to listen on (default "3000")

## Example

```
navi@navi-kubuntu:~$ fileserver --port=5000 --path=/home/navi/fileserver
Serving files under /home/navi/fileserver on port 5000
```

![Screenshot](./Screenshot_20220206_190418.png "Screenshot")