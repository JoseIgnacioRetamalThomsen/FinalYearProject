# Compile

Compile is done using maven in id

# Create docker image

navigate to docker folder

`
sudo docker image build -t joseretamal/profiles-dba:1.1 .
`

# run docker image
``
 sudo docker  run -d -p  5777:5777 --restart always --name na joseretamal/profiles-dba:1.1
``

# stop docker image
``
sudo docker container rm --force na
``

# push to docker hub

`
sudo docker push joseretamal/profiles-dba:1.1
` 


# run neo4j service on docker

sudo docker run \
    --name ns \
    -p7474:7474 -p7687:7687 \
    -d \
    -v $HOME/neo4j/data:/data \
    -v $HOME/neo4j/logs:/logs \
    -v $HOME/neo4j/import:/var/lib/neo4j/import \
    -v $HOME/neo4j/plugins:/plugins \
    --env NEO4J_AUTH=xxxx/ \
    neo4j:latest
