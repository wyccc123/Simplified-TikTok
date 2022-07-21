# Simplified-TikTok
A simplified version of TikTok backend
一个简化版的抖音后端实现

## Requirements - 必要条件
* Linux
* Goland1.16 and up
* Docker
* Kitex

## Framwork - 架构图
![image](https://user-images.githubusercontent.com/73453090/180171492-f2e853c5-d06c-4607-806c-cb1c351c4823.png)

## Running - 运行
### 运行基础依赖
```
docker-compose up -d
```

### 运行user服务
```
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 运行video服务
```
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```

### 运行comment服务
```
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

### 运行api服务
```
cd cmd/api
chmod +x ./run.sh
sh ./run.sh 
```
