# MyGram

MyGram is a final project from DTS Kominfo FGA, to build instagram clone API that allows users to store and comment on photos of other users. It is built using the Gin Gonic framework and Gorm ORM in Go.

## Table of Contents

- [Requirements](#requirements)
- [ERD (Entity Relationship Diagram)](#erd-entity-relationship-diagram)
- [Endpoints](#endpoints)
  - [User](#user)
  - [Social Media](#social-media)
  - [Photo](#photo)
  - [Comment](#comment)

## Requirements

- Go version 1.15 or higher
- PostgreSQL (13 or higher)

## ERD (Entity Relationship Diagram)

![ERD](./docs/erd.png)

## Endpoints

### User

- **Register [POST]**

  Register a new user.

- **Login [POST]**

  Login with a registered user.

### Social Media

- **GetAll [GET]**

  Get all social media posts.

- **GetOne [GET]**

  Get a specific social media post.

- **CreateSocialMedia [POST]**

  Create a new social media post.

- **UpdateSocialMedia [PUT]**

  Update an existing social media post.

- **DeleteSocialMedia [DELETE]**

  Delete a social media post.

### Photo

- **GetAll [GET]**

  Get all photos.

- **GetOne [GET]**

  Get a specific photo.

- **CreatePhoto [POST]**

  Upload a new photo.

- **UpdatePhoto [PUT]**

  Update an existing photo.

- **DeletePhoto [DELETE]**

  Delete a photo.

### Comment

- **GetAll [GET]**

  Get all comments.

- **GetOne [GET]**

  Get a specific comment.

- **CreateComment [POST]**

  Create a new comment.

- **UpdateComment [PUT]**

  Update an existing comment.

- **DeleteComment [DELETE]**

  Delete a comment.


## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/salmaqnsGH/my-gram.git
   ```

2. Change to the project directory:
    ```bash
    cd my-gram
    ```

3. Install the dependencies:
    ```bash
    go get
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

5. Access the application at http://localhost:3000 in your web browser.

##Acknowledgements
This project uses the following third-party libraries:

* Gin Gonic framework (github.com/gin-gonic/gin)
* Gorm ORM (github.com/go-gorm/gorm)
* jwt-go (github.com/dgrijalva/jwt-go)
* crypto (golang.org/x/crypto)

