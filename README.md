# pingcontainer_http_go

Http isteği ile çalışan docker ping konteynır

## Kurulum:

```shell
git clone https://github.com/HsmTeknoloji/pingcontainer_http_go.git
cd pingcontainer_http_go/devhsmtekpingcontainer
go mod init github.com/HsmTeknoloji/pingcontainer_http_go/devhsmtekpingcontainer
go get
cd ..
#Dockerfile içindeki tz değeri ayarlanmalı -> "ENV TZ=Europe/Istanbul"
sudo docker build -t devhsmtek/pingcontainerhttpgo .
#docker-compose.yml içindeki port değeri ayarlanmalı -> "PORT: 10001" 
sudo docker-compose up -d
```

## Kullanım:

```
http://127.0.0.1:{port}?ipaddress=192.168.1.1&count=5&timeout=5&v=1
```

## Geliştirici Bilgileri:
<img src="https://github.com/HsmTeknoloji/companyfiles/blob/master/hsmtek-logo.png?raw=true" width="200"/>
Web Site        : www.hsmteknoloji.com <br />
Developer Groups : https://t.me/HsmTeknoloji/ <br />