SyncUp is a modern dating application built on a robust, scalable microservices architecture. Designed for seamless connections and engaging interactions, our backend provides the foundation for a vibrant dating community.

🚀 Key Features
Modular Microservices: Independent services for Authentication, User Profiles, Matchmaking, Real-time Chat, and Media Management, ensuring scalability and resilience.

Event-Driven Communication: Leverages RabbitMQ for asynchronous processing of events like new matches and notifications.

Scalable Data Stores: Utilizes PostgreSQL for reliable data persistence and Redis for high-speed caching and real-time operations.

Unified API Gateway: A single entry point for all client interactions, handling routing, authentication, and rate limiting.

Containerized Development: Fully Dockerized for easy local setup and consistent deployment.

🏗️ Architecture Highlights
GoLang Backend: High-performance services written in Go.

gRPC & HTTP/REST: Efficient internal communication via gRPC, external APIs via HTTP/REST.

PostgreSQL & Redis: Primary and caching data stores.

RabbitMQ: Message broker for inter-service communication.

Docker & Kubernetes Ready: Designed for containerization and orchestration.

🏃‍♂️ Getting Started (Local Development)
Clone: git clone <your-repo-url>/dating-app-backend.git && cd dating-app-backend

Setup: Ensure Docker and Docker Compose are installed. Create a .env file from .env.example (if provided).

Run All Services: make dev-up

Access: The API Gateway is available at http://localhost:8080
