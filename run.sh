#!/bin/bash
if [ "$1" = "dev" ];
then
  echo $1
  echo $2
  echo $3
  ./main
fi

if [ "$1" = "uat" ];
then
  echo $1
  echo $2
  echo $3
  ./main --env=uat
fi

if [ "$1" = "public" ];
then
  echo $1
  echo $2
  echo $3
  php app.php --env=public --url=https://nacos.mabangerp.com/nacos/v1/cs/configs  --node.ip=$2 --node.port=$3  --dataId=mas_amz_fba_queue_data_id --group=prd_erp_group  --tenant=ae6ba4f8-dc91-42fb-b79c-4c4f7bb01482 --server.port=8080 --management.server.port=8081 --log.path=/data/logs
fi

if [ "$1" = "private" ];
then
  echo $1
  echo $2
  echo $3
  php app.php --env=private --url=https://nacos.mabangerp.com/nacos/v1/cs/configs   --node.ip=$2 --node.port=$3  --dataId=mas_amz_fba_queue_data_id --group=prd_erp_group  --tenant=a899e072-8d9c-411f-be51-cb6a467b868c --server.port=8080 --management.server.port=8081 --log.path=/data/logs
fi


