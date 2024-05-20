# Hotel reservation backend @zacksfF

## Business Problem
In the competitive hospitality industry, efficient and reliable management of hotel reservations is crucial for both customer satisfaction and operational success. Traditional reservation systems often suffer from scalability issues, limited flexibility, and difficulties in handling large volumes of concurrent users. Our Hotel Reservation Backend aims to address these challenges by leveraging modern technologies to provide a robust, scalable, and high-performance solution.

## Features
``This project contains the following features:``
- Middleware integration.
- User Authentication and Authorization: Secure login and registration system with JWT-based authentication.
- Room Management: CRUD operations for hotel rooms, including availability status and room details.
- Database support using MongoDB.
- Reservation System: Endpoints for creating, updating, and canceling reservations, with validation to ensure room availability.
- Search and Filter: Search for available rooms based on criteria such as date, room type, and price range.
- Notifications: Email notifications for booking confirmations and reminders.
- Distributed Logging and Tracing: Implementation of distributed logging and tracing to monitor and debug the application effectively.
- Integration with OpenTelemetry: Enterprise-level observability with integration to OpenTelemetry for monitoring and tracing.
- Testing patterns.
- Use of Docker, Docker Compose, and Makefiles.
- Github Actions Integration: Continuous Integration and Continuous Deployment (CI/CD) pipelines using CircleCI for automated testing, building, and deployment.



## Project environment variables
```
HTTP_LISTEN_ADDRESS=:3000
JWT_SECRET=somethingsupersecretthatNOBODYKNOWS
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

## Project outline
- users -> book room from an hotel 
- admins -> going to check reservation/bookings 
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Requirements
### Mongodb driver 
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber 
Documentation
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```

## Getting Started
NB: Make sure to have docker and docker-compose installed and running. For docker installation
```
git clone https://github.com/zacksfF/Hotel-Reservation-Backend.git
cd Hotel-Reservation-Backend
code . #if you are uisng vscode
```
- Run the application:
```
go run main.go
```

### Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes. Ensure your code follows the established style and includes appropriate tests.

### License
This project is licensed under the MIT License.

