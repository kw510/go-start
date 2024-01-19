# Go Start

A starting point for building a full stack application. The stack is:
- Postgres
- Go + Gin Router
- Vite + React + Typescript

## Starting the application

### Development
#### Start the application
```
docker-compose up --build
```
Which when complete, will expose these two ports:

- Backend: `http://localhost:8000/`
- Frontend: `http://localhost:3000/`

Both the backend and frontend are mostly hot reloadable.

If you install a new npm package, you will have to run the build command again. (or `exec` into frontend container and install it there)


#### Cleaning the database
If you need to reinitialise the database with `schema.sql`, then remove the volume for the database:
```
docker-compose down --volumes
```

### Production
The `Containerfile` will build a production ready image to use. You should `cd` into the respective service you want to build a image and use the following command 
```
docker build -t service-name . -f ./Containerfile
```
