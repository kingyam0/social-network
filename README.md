Objectives
* Followers
* Profile
* Posts
* Groups
* Notification
* Chat

FrontEnd
* Must consist of HTML, CSS, Javascript and Framework

BackEnd 
Must consist of 3 Major Parts
    1. Server
        * Which is the computer that receives the requests.
    2. App 
        * Listens for requests.
        * Retrieves information from the Database and sends responses.
        * Contains all information on how to respond to various requests based on HTTP or 
          Other types of Protocols.#
        * Some of the Handlers will be Middleware.
            - Middleware is any code that executes between the server receiving requests and 
              sending a response.
    3. Database
        * To organize and persist Data. 
        * Many requests sent to the server might require a Database Query.
            - A client might request information stored in the Database.
            - Or a client might submit data with their requests to be added to the Database.

Application 
    * Will consist of several Middleware.
        - Authentication(You must use sessions and cookies)
    * Images handling.
        - Must handle various types of extensions(eg. JPEG, PNG and GIF). It can be done by 
            storing the file/path in the Database and saving the image in specific file system.
    * Websockets
        - Handling connections in real time between clients.(This will help private chat).
    * Web Server
        - we can use Caddy which is written in GO or we can create one of our own.

SQlite
    * Creating you Database
        -structures that will benefit us.

Migrate
    * We will have to create migrations.
        - Every time the Application runs, it creates specific tables to make the project
            work properly.

        - Here is an Example of This :
                                    backend
                                    ├── pkg
                                    │   ├── db
                                    │   │   ├── migrations
                                    │   │   │   └── sqlite
                                    │   │   │       ├── 000001_create_users_table.down.sql
                                    │   │   │       ├── 000001_create_users_table.up.sql
                                    │   │   │       ├── 000002_create_posts_table.down.sql
                                    │   │   │       └── 000002_create_posts_table.up.sql
                                    │   │   └── sqlite
                                    │   |       └── sqlite.go
                                    |   |
                                    |   └── ...other_pkgs.go
                                    |
                                    └── server.go
        - We must implement a file structure similar to this as it will be tested.
        - We can use Golang-migrate package or other packages that may suit our project.
        - All migrations should be stored on a specific folder.
        - The sqlite.go should present the connection to the Database.(applying of the migrations
                and other functionalities that we may need to implement).
        - This migration will help you manage your time and testing by filling the database.

Docker
    * Create two images.
        1. One to serve the backend.
        2. Two to serve the FrontEnd.
    * Both backend and frontEnd must communicate.
    * The communication can be done on the browser as both will have to publish backend & FrontEnd.

Authentication
    * Email
    * Password
    * First Name
    * Last Name
    * Date of Birth
    * Avatar/Image (Optional)
    * Nickname (Optional)
    * About Me (Optional)

    - Note Avatar/Image, Nickname, AboutMe can be left BLANK.

Followers
    * Follow
    * Unfollow (you have to be following to Unfollow)
    * Follow requests need to be sent
    * Follow requests should be accepted first to follow the user.
    * If the user has private or public profile( needs to taken into account.)

Profile
    * Should display
        - User information
        - Recent Activity (every post made by user)
        - Followers 
        - following
    * There are two types of profiles
        - private
        - public
    * When user is in their profile they can change their profiles either private or public.

Post
    *Three types of Post
        -public(can be seen by all users)
        -private(can be seen only by followers)
        -almost Private (can be seen only by chosen followers)

Group
    * User Can create group
        - Should have title
        - Should have description chosen by the creator
        - Users need to be invited
        - Invited users need to accept before they can be a part of the group
        - Another way to enter group is by sending a request to join.(which can be accepted or decline by the creator of the group)
    * Must have a search section for all groups
    * Within the group they can create posts and comment on posts but it can only be seen by group members
    * Users within the group can create EVENTs:
                                        - Title
                                        - Description
                                        - Day / Time
                                        - 2 Options 
                                            - Going
                                            - Not Going

Chat
    * Private messages
    * Group Messages
    * Emojis need to enabled

Notifications
    * New notifications displayed on every page
    * New private messages should be notified in a different way.
    * Notifications for :
                        - Requests of followers
                        - Request of groups invitation
                        - Request to creator of the group if another users want to join.
                        - When event is created within the group

RUN FrontEnd:   npm run dev
MAKE SURE YOUR IN FRONTEND FOLDER