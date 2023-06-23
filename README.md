# Social Network Application

## Overview

This project aims to develop a social network application with various features such as user profiles, posts, groups, notifications, chat functionality, and a responsive front-end interface. The application will be built using HTML, CSS, JavaScript, and a suitable framework.

## Objectives

The main objectives of this project are as follows:

1. Implement the front-end interface using HTML, CSS, JavaScript, and a framework.
2. Develop the back-end with three major components: server, app, and database.
3. Create a database to store and persist data related to users, posts, groups, and notifications.
4. Implement middleware for handling various functionalities, including authentication, image handling, websockets for real-time connections, and a web server for serving the application.
5. Utilize SQLite as the database management system and create appropriate database structures.
6. Implement migrations to create and manage database tables required for the application to function properly.
7. Utilize Docker to create two images, one for serving the backend and another for serving the front-end.
8. Enable user authentication with email and password, along with additional optional information such as first name, last name, date of birth, avatar/image, nickname, and about me section.
9. Implement a follower system where users can follow or unfollow others, send follow requests, and accept or decline follow requests.
10. Develop user profiles that display user information, recent activity, followers, and following. Allow users to switch between private and public profiles.
11. Implement post functionality with three types: public (visible to all users), private (visible only to followers), and almost private (visible to selected followers).
12. Allow users to create groups with titles, descriptions, and invite others to join. Users can accept or decline invitations or send requests to join groups.
13. Enable search functionality for finding all groups.
14. Within groups, users can create posts and comment on them, visible only to group members. Users can also create events within groups, specifying title, description, day/time, and attendance options.
15. Implement chat functionality for private and group messages, including support for emojis.
16. Notify users of new notifications, with different notifications for followers, group invitations, join requests, and event creation within groups.
17. Provide an npm command to run the front-end development environment (`npm run dev`) within the frontend folder.
18. Run `go run main.go` or `go run .` to run the server with the backend folder.

## Technologies Used

The project will utilize the following technologies:

- HTML, CSS, and JavaScript for front-end development.
- A suitable framework for front-end development (e.g., React, Angular, Vue.js).
- Back-end development using a server, app, and database structure.
- SQLite as the database management system.
- Migration tools (e.g., Golang-migrate) for managing database structure and data.
- Docker for containerization, creating separate images for the backend and front-end.
- Sessions and cookies for user authentication.
- Websockets for real-time connections.
- Caddy or a custom web server for serving the application.

## Project Structure

The project's file structure should resemble the following example:

backend
├── pkg
│ ├── db
│ │ ├── migrations
│ │ │ └── sqlite
│ │ │ ├── 000001_create_users_table.down.sql
│ │ │ ├── 000001_create_users_table.up.sql
│ │ │ ├── 000002_create_posts_table.down.sql
│ │ │ └── 000002_create_posts_table.up.sql
│ │ └── sqlite
│ │ └── sqlite.go
│ └── ...other_pkgs.go
└── server.go

## Getting Started

To run the application locally, follow these steps:

1. Clone the repository.
2. Set up the front-end environment by navigating to the frontend folder and running the command `npm install`.
3. Start the front-end development environment by running `npm run dev` in the frontend folder.
4. Set up the back-end environment by configuring the server, app, and database components as specified in the documentation.
5. Run the back-end server using the appropriate commands or tools.
6. Access the application by opening a web browser and navigating to the specified URL.

## Contributors

- PENDING..