# Go Todo API

## Learning

In the process of learning, I thought it would be best to learn in a self-guided fashion and build a simple app from scratch.

## TODO

- [ ] auth
  - [x] base auth routes
  - [x] create user (signup)
    - [x] base route
  - [x] login user
    - [x] match if credentials are valid
  - [x] encrypt passwords
    - [x] bcrypt
    - [x] validation - password must be at least 8 characters
  - [x] set up jwt
    - ~~[ ] go-jwt~~
    - [x] fiber middleware
  - [ ] route protection
    ```go
    func restricted(c *fiber.Ctx) error {
      user := c.Locals("user").(*jwt.Token)
      claims := user.Claims.(jwt.MapClaims)
      name := claims["name"].(string)
      return c.SendString("Welcome " + name)
    }
    ```
- [ ] organization / architecture

  - [ ] proper organization of return statuses for routes
  - [ ] find more efficient way to test

- [ ] frontend
- [ ] scaffold
  - [ ] Next.js
  - [ ] typescript
