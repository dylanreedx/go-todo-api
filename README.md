# Go Todo API

## Learning

In the process of learning, I thought it would be best to learn in a self-guided fashion and build a simple app from scratch.

## TODO

- [ ] auth
  - [x] base auth routes
  - [ ] create user (signup)
    - [x] base route
  - [x] login user
    - [x] match if credentials are valid
  - [ ] encrypt passwords
    - [x] bcrypt
    - [ ] validation - password must be at least 8 characters
  - [x] set up jwt
    - ~~[ ] go-jwt~~
    - [x] fiber middleware
