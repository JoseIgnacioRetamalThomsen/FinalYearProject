# compile
go build -o pdba *.go

#run 
./pdba config.json

# Create image
sudo docker image build -t joseretamal/photos-service-dba:1.1 .

#push image

sudo docker push joseretamal/photos-service-dba:1.1


# run

sudo docker run -d -p 7172:7172 --network="host" --restart always --name photodba joseretamal/photos-service-dba:1.0
