# pingcontainer_http_go

Docker ping container working with http request 

## Installation:

```shell
git clone https://github.com/AfatekDevelopers/pingcontainer_http_go.git
cd pingcontainer_http_go/devafatekpingcontainer
go mod init github.com/AfatekDevelopers/pingcontainer_http_go/devafatekpingcontainer
go get
cd ..
#tz value in the Dockerfile should be set -> "ENV TZ=Europe/Istanbul"
sudo docker build -t devafatek/pingcontainerhttpgo .
#port value in the docker-compose.yml should be set -> "PORT: 10001" 
sudo docker-compose up -d
```

## Usage:

```
http://127.0.0.1:{port}?ipaddress=192.168.1.1&count=5&timeout=5&v=1
```

## Developers:
<img src="https://github.com/AfatekDevelopers/companyfiles/blob/master/afatek-logo.png?raw=true" width="200"/>
Web Site        : www.afatek.com.tr <br />
Developer Groups : https://t.me/Afatek/ <br />