# Project Name

This project uses Docker Compose to manage several services, including a Go server, PostgreSQL, Redis, pgAdmin, and Jenkins. Below are the instructions to set up and run the project using Docker Compose and expose the necessary services.

## Prerequisites

Before running the services, ensure you have the following installed:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/onlymachiavelli/Tendanz-challenge-api
   cd Tendanz-challenge-api
   ```

2. Build the Docker images (if not already done):

   ```bash
   docker-compose up -d
   ```

## Running Docker Compose

To start all services defined in the `docker-compose.yml` file:

1. Run the following command:

   ```bash
   docker-compose up
   ```

   This will start all the services in detached mode. If you want to see the logs of the services, use:

   ```bash
   docker-compose up -d
   ```

2. To stop the services, run:

   ```bash
   docker-compose down
   ```

3. To restart the services:

   ```bash
   docker-compose restart
   ```

## Exposing Services

### Services and Exposed Ports:

1. **Go Server (`go-server`)**

   - **Exposed Port**: `8080` (Local) → `80` (Container)
   - Accessible at: `http://localhost:8080`

2. **Redis (`redis`)**

   - **Exposed Port**: `6379` (Local) → `6379` (Container)
   - Accessible at: `redis://localhost:6379`

3. **PostgreSQL (`postgres`)**

   - **Exposed Port**: `5432` (Local) → `5432` (Container)
   - Accessible at: `postgresql://localhost:5432/tendanz`

4. **pgAdmin (`pgadmin`)**

   - **Exposed Port**: `5050` (Local) → `80` (Container)
   - Accessible at: `http://localhost:5050`
   - Default login credentials:
     - **Email**: `admin@admin.com`
     - **Password**: `admin`

5. **Jenkins (`jenkins`)**
   - **Exposed Port**: `8081` (Local) → `8080` (Container), `50000` (Local) → `50000` (Container)
   - Accessible at: `http://localhost:8081` (for Jenkins UI)

### Environment Variables

- **Go Server**:

  - `PORT=80`: The port the Go server listens on.
  - `DATABASE_URL=postgresql://user:root@postgres:2345/tendanz`: Connection URL for PostgreSQL.
  - `REDIS_URL=redis://redis:6379`: Connection URL for Redis.

- **PostgreSQL**:

  - `POSTGRES_USER=user`: Database user.
  - `POSTGRES_PASSWORD=root`: Database password.
  - `POSTGRES_DB=tendanz`: Database name.

- **pgAdmin**:
  - `PGADMIN_DEFAULT_EMAIL=admin@admin.com`: Default admin email.
  - `PGADMIN_DEFAULT_PASSWORD=admin`: Default admin password.

## Volumes

- **postgres_data**: Stores PostgreSQL data persistently.
- **jenkins_home**: Stores Jenkins data persistently.

## Accessing Logs

You can view the logs of a specific service using:

```bash
docker-compose logs <service-name>
```
