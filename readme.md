# BoWatt Docs

This Readme covers the application

## Build
The project is completely dockerized and should be very easy to start up
`docker-compose up --build` will build the entire project

you will get 3 containers
- Backend Go server using gin, this is a very simply and straight forward server using MVC, for a full scale application I recommend using onion architecture, far more complex but also far more scalable and mantainable for large code bases, as well as far far more extendable.
- Frontend, React app by create-react-app
- Mongo database

All parts of the application are dockerized and can be orchestrated with docker-compose, `docker-compose up --build` gets you started
Files are uploaded one at a time or as a batch there are two endpoints with the multiple upload endpoint being able to upload one or more files, contents are extracted and saved in a mongodb, the entire file is hashed and this hash is what is used to detect duplicate files, preventing users from submitting the same file over and over again, but allowing them to upload files with subtle changes, different users may upload the same file.

- user management used a basic token (simple UUID) I could have gone with something more complex such as jwt but the basic principals are the same and I wanted to keep things simple but functional as this is just a one time use demo application

- we have a basic in memory cache which again is just a map, but it is enough to demonstrate the proof of concept.

## User Management
Basic user management is included a user can
- signup
- login

# Auth
- The backend expects a Authorization header with the token which is provided on sign up or on login
- Frontend will automatically send this token back to the backend with each request, which is how the backend can differentiate betwean the different users, which we do via our handy middleware which will search the db for a user which this token and add this to the context allowing us to reference this user throughout the request lifecycle

# Frontend
- Uses react router to do routing
- Stores tokens in local storage
- uses reacts useState hooks for state management, there was no need for anything more complex but we could have gone with tencent query or zustand
- requests are done with fetch
-  A user can upload as many files as they wish one after the other and they can select one at a time and search them
-  we could upload the original file to a cms or aws bucket or something similar after upload
- allow multi file uploads

# Backend routes


### Unprotected routes
| Method | Endpoint  | Description       |
| ------ | --------- | ----------------- |
| GET    | `/ping`   | Health check      |
| POST   | `/signup` | Register new user |
| POST   | `/login`  | User login        |


### Protected routes
these routes require a header
Authorization: <token>
| Method | Endpoint   | Description            |
| ------ | ---------- | ---------------------- |
| GET    | `/profile` | Get user profile       |
| GET    | `/files`   | List uploaded files    |
| POST   | `/upload`  | Upload a file          |
| POST   | `/uploads` | Upload multiple files  |
| POST   | `/query`   | Search content in file |

# Architecture
To keep it simple I have gone with a basic implementation of the MVC architecture, models, views and controller, for a more complex programs I would recommend going with onion/ports and adapters as its much more adaptable, 
