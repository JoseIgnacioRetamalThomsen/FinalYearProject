# Compile to docker folder

``
 sudo go build -o docker/main *.go
``

# Run 

``
docker/main docker/config.json
``

# Create docker image

`` 
cd docker
``

``
sudo docker image build -t joseretamal/auth-service:1.1 .
``

# run docker image
``
 sudo docker container run --publish 50051:50051 --detach --name as joseretamal/auth-service:1.0
``
better for auto restart

`
docker run -d -p 50051:50051 --restart always --name as joseretamal/hash-service:1.1
`



# stop docker image
``
sudo docker container rm --force as
``

# push image to docker hub
``
docker login
``

``
sudo docker push joseretamal/auth-service:1.1
``

# Check logs
`
sudo docker logs as
`

# pull
`
sudo docker pull joseretamal/auth-service
`
