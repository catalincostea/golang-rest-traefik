# Overview 

In total, there are four tasks.

# Task 1 
In the `/api` folder, you will see a very basic Go lang API service. 
You can start the application running `go run main.go` (you need Go installed on your machine).

The aim of this challenge is to have a glimpse of your programming skills. 
There is a method, `getAllMovies`, that needs to be updated. 

Please make a call for this public API: https://swapi.dev/api/films/ so that your API will return a list of all the Star Wars movie titles. You can dismiss any other data from that call.

# Task 2 

This challenge aims to test some of your Docker skills. 

Please create a working Dockerfile for your `API` application. 
You should be able to run your application as a container and specify the port under which it should be exposed (with a minor code modification). 

# Task 3 

Here, we will test your docker-compose and networking skills (basic Traefik proxy configuration).

You can see there are a few services already in this `docker-compose.yaml` file:
- `traefik`: you don't need to change anything here 
- `starwars`: that's the service you should update

You should:
- Update this service so it runs your application from your Dockerfile 
- No need to expose any ports. Please configure labels so that you will be able to access your application via `traefik` service, calling http://localhost/movies

# Task 4 
For the last task, we want to check your Terraform skills.

Please navigate to `terraform` folder.
A data resource already makes calls to the same Star Wars API endpoint as in the first task. 
Please update the logic as follows: 
- All the Star Wars movie titles are returned as output (list of strings) 
- All the Star Wars movie titles are saved into the local fine, each in the new line.

You can see the expected result in `expected_result-all-titles.txt`