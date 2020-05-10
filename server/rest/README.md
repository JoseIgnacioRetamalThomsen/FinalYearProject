# Compile to docker folder

``
 sudo go build -o main *.go
``

# Run 

``
./main config.json
``

# Create docker image

`` 
cd docker
``

``
sudo docker image build -t joseretamal/wcity-rest:1.1 .
``

#push to docker hub

```
sudo docker push joseretamal/wcity-rest:1.1
```

# run docker image
``
 sudo docker container run --publish 50051:50051 --detach --name as joseretamal/auth-service:1.0
``


better for auto restart

`
docker run -d -p 9371:9371 --restart always --name wrest joseretamal/wcity-rest:1.0
`



# stop docker image
``
sudo docker container rm --force as
