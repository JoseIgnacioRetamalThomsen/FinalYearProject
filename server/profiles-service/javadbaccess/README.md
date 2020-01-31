# Compile

Compile is done using maven in id

# Create docker image

navigate to docker folder

`
sudo docker image build -t joseretamal/profiles-dba:1.0 .
`

# run docker image
``
 sudo docker  run -d -p  5077:5077 --restart always --name na joseretamal/profiles-dba:1.0
``

# stop docker image
``
sudo docker container rm --force na
``

# push to docker hub

`
docker push joseretamal/profiles-dba:1.0
` 