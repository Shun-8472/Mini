# 🚀 gRPC Microservice with LLM Integration

This is a simple **microservice** implemented with **gRPC** that interacts with an **LLM Model** to generate responses.

## 📌 Features
- **gRPC Communication**: Efficient inter-service communication using **Protocol Buffers (protobuf)**.
- **LLM Integration**: Uses **Ollama** or other models to generate responses.
- **Dependency Injection**: Managed with **Google Wire**.
- **Database Support**: Uses **Redis** and **MySQL**.

## 🔹 Installation
Before running the service, install the following dependencies:
### **1️⃣ Download Docker Images**
- [Docker](https://www.docker.com/)
    ```shell
    docker pull redis:latest
    docker pull mysql:latest
    docker pull ollama/ollama:latest
    ```
### **2️⃣ Start Redis (Docker)**
- [Redis](https://formulae.brew.sh/formula/redis)
    ```shell
    docker run -d --name redis-server -p 6379:6379 redis:latest
    ```
### **3️⃣ Start MySQL (Docker)**
- [MySQL](https://formulae.brew.sh/formula/mysql)
    ```shell
    docker run -d --name mysql-server -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=mydb mysql:latest
    ```
### **4️⃣ Start Ollama (Docker)**
- [Ollama](https://ollama.com/)
    ```shell
    docker run -d --name ollama-server -p 11434:11434 ollama/ollama:latest
    ```
### **5️⃣ Install buf (for gRPC Protobuf management)**
- [buf](https://formulae.brew.sh/formula/buf)
    ```shell
    $ brew install buf
    ```
### **6️⃣ Install wire (for Dependency Injection)**
- [wire](https://github.com/google/wire)
    ```shell
    $ go install github.com/google/wire/cmd/wire@latest
    ```
  
## 🚀 Running the Service
You can use the Makefile to build and run the service efficiently.
### **1️⃣ Tidy up Go modules**
```shell
  make tidy
```
### **2️⃣ Generate Dependency Injection Code**
```shell
  make inject
```
### **3️⃣ Deploy the service using Docker Compose**
```shell
  make deploy
```