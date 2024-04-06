# Assets Manager 

This is a simple Go application that manages assets and organize it by groups and set which user are using it.

This is beeing made using Hexagonal Architecture/Ports and Adapters concept.


## Entities

### User
- ID (int)
- Name (string)
- Email (string)

### Group
- ID (int)
- Name (string)

### Asset
- ID (int)
- Name (string)
- Status ([ACTIVE, BUSY, REPAIR, RESERVED, INACTIVE])
- Group (Group)
- CurrentUser (User)

## Routes

### User
- POST /users - Create user
- GET /users - List all users based on query
- GET /users/:id - Search user by id
- PUT /users/:id - Update user data
- DELETE /users/:id - Remove user

### Group
- POST /groups - Create group
- GET /groups - List all groups based on query
- GET /groups/:id - Search group by id
- PUT /groups/:id - Update group data
- DELETE /groups/:id - Remove group

### Asset
- POST /assets - Create asset
- GET /assets - List all assets based on query
- GET /assets/:id - Search asset by id
- PUT /assets/:id - Update some asset data
- DELETE /assets/:id - Remove asset
- PUT /assets/:id/status - Update asset status
- PUT /assets/:id/user - Update asset current user