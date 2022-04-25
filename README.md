# dev-jokes
An application to test out goth and chuck norris api. This is also to explose github workflow for CI/CD and its integrity with AWS. 

# Prerequisites 
* Golang 1.17 or higher
* Docker 
* Docker compose (for local installation)
* A google account
* A github account

### Run locally 


Here we build the source code locally and run it in a container. To run this applicaiton we need some client secrets from our Auth service providers in our case those are google and github. We would require following informations for those provides ClientId, Client Scecret and Callback url. For google here you can create Auth 2.0 clients for your applications
https://console.cloud.google.com/apis/credentials (You must login)
For github once you are logged in you can go to `Settings` from top righ menu then go to  `Applications` under Integrations, you should be able to create auth clients under the `Authorised OAuth Apps`
Great now we are done with all the boring stuff. Lets create an `.env` file with follwing keys and values for these would be what we get from google and github after creating clients. Here callback url is the catch, this is where google will redirect back after authentication which should be our application url. In Oauth2 this url need to be an exact match so make sure you got it correctly defined in .env file. 

```
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
GOOGLE_CALLBACK_URL=
### github configuration
GITHUB_CLIENT_ID=
GITHUB_CLIENT_SECRET=
GITHUB_CALLBACK_URL=
```
Once we have created the .env file with correct values we should be now ready to run the application and to do so we just need to do the following.

```
cd dev-jokes
make 
make deploy

```

This should start the application on http://localhost:3000.


### CI/CD

We have tried explore github CI/CD here, what I have liked the most about it is the what ever job/stage we are writing can be tested locally. For that you need to install one package called `act`. Then you can just run `act` in your terminal at the root of the project. You should be able to run a perticular job by `act -j <job_name>`

```
[checks/test                ] 🚀  Start image=catthehacker/ubuntu:act-latest
[checks/lint                ] 🚀  Start image=catthehacker/ubuntu:act-latest
[Deploy to Amazon ECS/Deploy]   🐳  docker pull image=catthehacker/ubuntu:act-latest platform= username= forcePull=false
````

Currently, we have considered following actions 
* Test go code
* Lint & format go code
* Build a docker image
* Deploy it to ECR
