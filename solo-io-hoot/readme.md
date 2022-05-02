# solo-io/hoot 学习笔记

本目录是记录学习 [solo-io-hoot](https://github.com/solo-io/hoot/) 旅程。

## ubuntu 安装 envoy

```
sudo apt update
sudo apt install apt-transport-https gnupg2 curl lsb-release
curl -sL 'https://deb.dl.getenvoy.io/public/gpg.8115BA8E629CC074.key' | sudo gpg --dearmor -o /usr/share/keyrings/getenvoy-keyring.gpg
echo a077cb587a1b622e03aa4bf2f3689de14658a9497a9af2c427bba5f4cc3c4723 /usr/share/keyrings/getenvoy-keyring.gpg | sha256sum --check
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/getenvoy-keyring.gpg] https://deb.dl.getenvoy.io/public/deb/ubuntu $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/getenvoy.list
sudo apt update
sudo apt install -y getenvoy-envoy
```

## docker 

```
docker run -it  --restart=always  -d  envoyproxy/envoy-dev:25cc9ea48ba18930e1235dfb6a6a236d59c8b08a  bash
```