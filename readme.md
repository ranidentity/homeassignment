1. docker-compose up -d
2. run sql script to create tables - /database/sql/init.sql
3. go mod tidy
4. make a copy of .env from example.env


System flow
- main -> router -> controller -> service -> repo -> model -> database

API documentation - Swagger
http://localhost:8080/swagger/index.html

Sample postman collection is included


dockerize project
docker build -t my-go-app .
docker run --env-file .env -p 8080:8080 my-go-app

1. Build Your Docker Image Inside Minikube
By default, Minikube uses its own Docker daemon. To build images directly accessible to Minikube:
minikube docker-env

Then run the command it suggests, e.g.:
eval $(minikube -p minikube docker-env)

Now, build your image as usual:
docker build -t my-go-app:latest .

2. Write Kubernetes Manifests
a. Deployment (deploys your app)
Create a file called deployment.yaml:

b. Service (exposes your app)
Add to service.yaml:

c. (Optional) Secret for .env
If you want to use your .env file, create a Kubernetes Secret:
kubectl create secret generic my-go-app-env --from-env-file=.env

3. Apply the Manifests
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

4. Access Your App
Get the URL to access your service:
minikube service my-go-app-service