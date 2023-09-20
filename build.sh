#!/bin/bash

docker system prune -a -f

P=`pwd`
REPO=${P##*/}

git pull origin master

ID=`git rev-parse --short HEAD`
AUTHOR=`git log --pretty=format:"%an" $ID -1`

CONTAINERID=$REPO:$AUTHOR-$ID

docker stop $REPO

I=`docker inspect --format='{{.Image}}' $REPO`
IMAGEID=${I##*:}

docker rm -f $REPO

docker rmi $IMAGEID

docker build -t $CONTAINERID .

# 运行时的容器名亦是仓库名
docker run -itd --restart=always -p 924:924  --name $REPO $CONTAINERID

sudo rm -rf ./main