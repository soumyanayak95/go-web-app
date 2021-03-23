# go-web-app

This is a simple project built to learn the basics of golang web application.

Steps
1. Build the image of the projecting using the dockerfile:
    docker build -t my-go-app .

2. Run the container:
    winpty docker run -p 8000:8000 -it my-go-app
    * -it : for interactive bash terminal

3. http://192.168.99.100:8000/  -> app url
    * 192.168.99.100: ip of the docker-machine it is running on. 
    * docker-machine ls

4. The repo is cloned inside the container.
