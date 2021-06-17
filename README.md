# GoLang Rest Api

## Features Implemented

1. Able to create user account
2. Able to Login to the account
3. Create and set jwt token as cookie
4. Able to logout

## What did I learn?

1. Learn to use Gorm to connect with the database
2. Use fiber handle routes
3. Creating the jwt token
4. Hashing the password with bcrypt
5. Handling errors
6. Creating docker compose file
7. Using environment variables

## How to run?

1. Install go and Clone the repo
2. Install the Docker and run `docker compose up` - this will run the database
3. Run main.go with the `DB_NAME`, `DB_USERNAME` and `DB_PASSWORD` environment variables
4. Test the application to create and login to the account, get user and logout

## TODO:

1. Write the tests