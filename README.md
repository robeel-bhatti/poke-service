# Poke Service

A backend microservice written in Golang.

## About

I am doing this project because I am learning Golang, so I am building a 
microservice application to apply my Go learnings.

This project invokes the Pokemon API to return different types of Pokemon data.

This Pokemon data is then returned to the Pokemon UI I am creating in React.

## Running Locally

Feel free to clone if you want to play around with it.

See the API specification for the endpoints in this project so that you can invoke them from Swagger
or your local API tool (ex. Postman, Hoppscotch).

Run `go run main.go` and the application will start and be ready to process HTTP requests.

## Committing

Use [Angular Conventional Commit](https://github.com/angular/angular/blob/main/contributing-docs/commit-message-guidelines.md)
when committing to this repository for standardized, clear commit messages.

## Merging

This repo's merging strategy is `squash and merge`. This is to create a clean, linear merge history on the `main` branch. 
Individual commits are retains in the merge commit description. 
A link to the PR that created this merge commit is in the merge commit title for easier access in the event of debugging.
