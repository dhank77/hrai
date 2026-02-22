# HRAI (HR AI Project)

A modern Human Resources application built with **Golang** (Backend) and **Next.js** (Frontend).

## 🚀 Tech Stack

### Backend
- **Framework**: [Gin](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Live Reload**: [air](https://github.com/cosmtrek/air)

### Frontend
- **Framework**: [Next.js](https://nextjs.org/)
- **Linkage**: Connects to the Backend via REST API

---

## 🛠️ Getting Started

### Prerequisites
- **Go**: v1.25 or higher
- **Node.js**: v18.x or higher
- **PostgreSQL**: Running locally or or via Docker
- **Air**: `go install github.com/cosmtrek/air@latest` (Optional, for development)

### 1. Backend Setup

1.  **Navigate to the Backend directory**:
    ```bash
    cd Backend
    ```

2.  **Configure Environment Variables**:
    Copy the example file and update it with your database credentials:
    ```bash
    cp .env.example .env
    ```
    Edit `.env`:
    ```dotenv
    APP_PORT=9090
    JWT_SECRET=your_super_secret_key
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=hrai_db
    DB_PORT=5432
    ```

3.  **Database Migration**:
    Ensure your PostgreSQL database is running and the database specified in `.env` exists. You can use the provided SQL file for initial schema:
    ```bash
    psql -U postgres -d hrai_db -f database/users.sql
    ```

4.  **Run the Application**:
    Using `air` (recommended for dev):
    ```bash
    air
    ```
    Or standard go run:
    ```bash
    go run main.go
    ```

### 2. Frontend Setup

1.  **Navigate to the Frontend directory**:
    ```bash
    cd ../Frontend
    ```

2.  **Install Dependencies**:
    ```bash
    npm install
    # or
    yarn install
    ```

3.  **Configure Environment Variables**:
    Create a `.env.local` file:
    ```bash
    NEXT_PUBLIC_API_URL=http://localhost:9090
    ```

4.  **Run Development Server**:
    ```bash
    npm run dev
    ```
    Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

---

## 📂 Project Structure

- `Backend/`: Golang source code, configurations, and database logic.
  - `module/`: Core business logic (Dashboard, Employee, Users).
  - `config/`: Router and application settings.
  - `database/`: DB connection and SQL migration files.
- `Frontend/`: Next.js application (Pages, Components, Styles).

---

## 📝 License

This project is licensed under the MIT License.