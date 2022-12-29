# Virtual Internship Program - BTPN Syariah Fullstack Developer
This is Project for Virtual Internship Program - BTPN Syariah Fullstack Developer by Rakamin.
This project is created using Golang Programming Language with Echo Framework, GORM for object relational mapping for PostgreSQL database.

## Usage
* Change the database configuration in `app.config.yaml` file.
* Run with `go run main.go`
* You're ready to go!

## Endpoints
|  Method | URL | Description |
| ------------ | ------------ | ------------ |
| POST | /users/register | Register a new `user` |
| POST | /users/login | Login `user` |
| PUT | /users/:id | Update an `user` |
| DELETE | /users/:id | Delete an `user` |
| POST | /photos | Create a new `photo` |
| GET | /photos | Get all `photo` |
| PUT | /photos/:id | Update a `photo` |
| DELETE | /photos/:id | Delete a `photo` |
