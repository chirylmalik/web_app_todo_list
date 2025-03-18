# **To-Do List Web App**

A simple **To-Do List Web Application** built with **Golang (Backend), MySQL (Database), and jQuery (Frontend)**.

## **📌 Features**
✅ **Add Task** – Users can add new tasks to the list.  
✅ **Mark Task as Completed** – Toggle task status (completed or not).  
✅ **Edit Task** – Update existing task with confirmation.  
✅ **Delete Task** – Remove a task from the list.  
✅ **Dynamic Notifications** – User-friendly alerts for actions.  
✅ **AJAX Integration** – Seamless UI updates without page reloads.  

## **🛠️ Tech Stack**
### **Backend:**
- **Golang** – Main backend language
- **Gorilla Mux** – Router for handling REST API
- **MySQL (MariaDB 10.4.24)** – Database for storing tasks

### **Frontend:**
- **HTML, CSS** – Basic UI structure and styling
- **jQuery + AJAX** – Dynamic interaction with API

## **⚙️ Installation & Setup**
### **1. Clone the Repository**
```bash
git clone https://github.com/yourusername/todo-list-app.git
cd todo-list-app
```

### **2. Setup Database (MySQL / MariaDB)**
Create a database named **`todo_list`**, then create a `tasks` table:
```sql
CREATE DATABASE todo_list;
USE todo_list;

CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT FALSE
);
```

### **3. Configure and Run Backend**
Make sure you have **Golang** installed. Then, install dependencies and run the server:
```bash
go mod tidy
go run main.go
```
By default, the backend will run on **http://localhost:8080**.

### **4. Open Frontend**
Simply open `index.html` in a browser.

## **🚀 API Endpoints**
| Method | Endpoint            | Description              |
|--------|---------------------|--------------------------|
| GET    | `/tasks`            | Get all tasks           |
| POST   | `/tasks`            | Add a new task          |
| PUT    | `/tasks/{id}`       | Update an existing task |
| DELETE | `/tasks/{id}`       | Delete a task           |
