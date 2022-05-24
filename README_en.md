# pingcontainer_http_go

Docker ping container working with http request 

## Installation:

```shell
git clone https://github.com/HsmTeknoloji/pingcontainer_http_go.git
cd pingcontainer_http_go/devhsmtekpingcontainer
go mod init github.com/HsmTeknoloji/pingcontainer_http_go/devhsmtekpingcontainer
go get
cd ..
#tz value in the Dockerfile should be set -> "ENV TZ=Europe/Istanbul"
sudo docker build -t devhsmtek/pingcontainerhttpgo .
#port value in the docker-compose.yml should be set -> "PORT: 10001" 
sudo docker-compose up -d
```

## Usage:

```
http://127.0.0.1:{port}?ipaddress=192.168.1.1&count=5&timeout=5&v=1
```

## Developers:
<img src="https://github.com/HsmTeknoloji/companyfiles/blob/master/hsmtek-logo.png?raw=true" width="200"/>
Web Site        : www.hsmteknoloji.com <br />
Developer Groups : https://t.me/HsmTeknoloji/ <br />