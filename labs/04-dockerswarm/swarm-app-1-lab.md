#SWARM Labs
Labs  (https://github.com/BretFisher/udemy-docker-mastery/tree/master/swarm-app-1)

## Lab1

```
docker network create --driver overlay  frontend
docker network create --driver overlay  backend

docker service create --replicas 3 --network frontend --publish 80:80  --name vote dockersamples/examplevotingapp_vote:before  
docker service create --replicas 1 --network frontend  --name redis redis:3.2
docker service create --replicas 1 --network frontend --network backend  --name worker dockersamples/examplevotingapp_worker
docker service create --replicas 1 --network backend --mount type=volume,source=db-data,target=/var/lib/postgresql/data    --name db postgres:9.4
docker service create --replicas 1 --network backend  --publish 5001:80 --name result dockersamples/examplevotingapp_result:before


```


$ docker inspect network frontend | jq ".[].Containers"
Error: No such object: network
{
 "92191b11b15ea7a6494b3cbf2f68f03e9b5dee38c7785a820788a92604409278": {
   "Name": "redis.1.dpu32ycgxo0pu1fxk5w8qpu0x",
   "EndpointID": "9f2950e5961b29d49bed5b8124af5a4d69afd8a6097ba0784810399076c395f6",
   "MacAddress": "02:42:0a:00:01:67",
   "IPv4Address": "10.0.1.103/24",
   "IPv6Address": ""
 },
 "ae4871c8b66d6ab067a66a88f870afc003766c3526ecacb0c84738ce3ca8b7b5": {
   "Name": "vote.1.uipzbzpssaqheea7hafyr2qx8",
   "EndpointID": "8da807f648a87ca8d4d1b1c132f60c115a6c94cdf5f6557c25dab89f153000ae",
   "MacAddress": "02:42:0a:00:01:64",
   "IPv4Address": "10.0.1.100/24",
   "IPv6Address": ""
 },
 "lb-frontend": {
   "Name": "frontend-endpoint",
   "EndpointID": "acd90a13187c88d9d47be87ca70826308ba83bd8036d081ac6dbfe7696f4bf2f",
   "MacAddress": "02:42:0a:00:01:65",
   "IPv4Address": "10.0.1.101/24",
   "IPv6Address": ""
 }
}
