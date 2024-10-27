# Project Name

## Overview

This project is an API server that provides functionality for managing users, their League of Legends accounts, and game accounts. It is built using Go and utilizes Gorilla Mux for routing and SQLite for data storage.

## Features

- Create, update, and delete user accounts
- Manage user League of Legends profiles
- Handle game accounts associated with users

## Endpoints

- **User Management**
  - `POST /user`: Create a new user
  - `GET /user`: Retrieve all users
  - `GET /user/{id}`: Retrieve a user by ID
  - `PUT /user/{id}`: Update a user's information
  - `DELETE /user/{id}`: Delete a user

- **League of Legends Management**
  - `POST /user/{id}/league_of_legends`: Create or update a League of Legends profile for a user
  - `DELETE /user/{id}/league_of_legends`: Delete a user's League of Legends profile

- **Game Account Management**
  - `POST /user/{id}/game_account`: Create or update a game account for a user
  - `DELETE /user/{id}/game_account`: Delete a game account from a user

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/yourrepository.git
   ```
2. Navigate to the project directory:
   ```bash
   cd yourrepository
   ```
3. Build the project:
   ```bash
   go build
   ```
4. Run the server:
   ```bash
   ./yourrepository
   ```

## Database

The project uses SQLite for database management. The schema is defined in `schema.sql` and is executed at the start of the server to ensure all tables are created.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
