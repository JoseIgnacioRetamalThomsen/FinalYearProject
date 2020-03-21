# create docker image
  sudo docker image build -t joseretamal/post-service-dba:1.2 .

  # run

  sudo docker run -d -p 2787:2787 --network="host" --restart always --name postdba joseretamal/post-service-dba:1.2 

  # push
  sudo docker push joseretamal/post-service-dba:1.2

  # pull 
  sudo docker pull joseretamal/post-service-dba:10.
