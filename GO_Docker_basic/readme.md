## Dockerize Your Go Application

This is a simple demonstration of how to dockerize your Go application and run it inside a Docker container.

### Commands

1. Build the Docker image:
   ```
   docker build --tag docker-gs-ping .
   ```

2. Run the Docker container:
   ```
   docker run docker-gs-ping
   ```

3. Run the Docker container and publish the port:
   ```
   docker run --publish 8080:8080 docker-gs-ping
   ```

   To run in detach mode:
   ```
   docker run -d -p 8080:8080 docker-gs-ping
   ```

   Now you can use the curl command to interact with your application:
   ```
   curl http://localhost:8080/
   ```

   In the above commands, "docker-gs-ping" is the name of the container, which can be defined as anything in the Dockerfile.

### Resources

For more information, you can refer to the official Docker documentation on building Docker images for Go applications: [Link](https://docs.docker.com/language/golang/build-images/)
