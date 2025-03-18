const API_URL = "http://localhost:8080/tasks"
let editingId = null
let oldTask = ""

$(document).ready(function () {
    fetchTasks()
})

function fetchTasks() {
    $.get(API_URL, function (tasks) {
        console.log("Response from API:", tasks)
        
        if (typeof tasks === "string") {
            try {
                tasks = JSON.parse(tasks)
            } catch (error) {
                console.error("Failed to parse JSON:", error)
                return
            }
        }

        $("#todoList").empty()
        tasks.forEach(task => {
            let checked = task.completed ? "checked" : ""
            let textDecoration = task.completed ? "line-through" : "none"

            $("#todoList").append(`
                <li>
                    <div class="task-item">
                        <input type="checkbox" class="task-checkbox" ${checked} onclick="toggleComplete(${task.id}, ${!task.completed})">
                        <span class="task-text" id="task-${task.id}" style="text-decoration: ${textDecoration};">${task.task}</span>
                    </div>
                    <div class="task-actions">
                        <button onclick="editTask(${task.id}, '${task.task}')">Edit</button>
                        <button class="del-btn" onclick="confirmDelete(${task.id})">Delete</button>
                    </div>
                </li>
            `)
        })

        resetInput()
    })
}

function addTask() {
    let task = $("#taskInput").val().trim()
    if (!task) {
        showNotification("Task cannot be empty!", "error")
        return
    }

    $.ajax({
        url: API_URL,
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({ task: task, completed: false }),
        success: function () {
            fetchTasks()
            showNotification("Task added successfully!", "success")
        }
    })
}

function editTask(id, task) {
    $("#taskInput").val(task)
    $("#addBtn").hide()
    $("#saveEditBtn").show()
    editingId = id
    oldTask = task
}

function saveEdit() {
    let newTask = $("#taskInput").val().trim()
    if (!newTask || editingId === null) {
        showNotification("Task cannot be empty!", "error")
        return
    }

    if (oldTask === newTask) {
        let continueEditing = confirm("No changes were made. Do you want to continue editing this task?")
        if (!continueEditing) {
            resetInput()
        }
        return
    }

    let isConfirmed = confirm(`Do you want to save these changes?\n\nFrom: "${oldTask}"\nTo: "${newTask}"`)
    
    if (!isConfirmed) {
        resetInput()
        return
    }

    $.ajax({
        url: `${API_URL}/${editingId}`,
        method: "PUT",
        contentType: "application/json",
        data: JSON.stringify({ task: newTask, completed: false }),
        success: function () {
            fetchTasks()
            showNotification("Task updated successfully!", "success")
        }
    })
}

function confirmDelete(id) {
    if (confirm("Are you sure you want to delete this task?")) {
        deleteTask(id)
    }
}

function deleteTask(id) {
    $.ajax({
        url: `${API_URL}/${id}`,
        method: "DELETE",
        success: function () {
            fetchTasks()
            showNotification("Task deleted successfully!", "error")
        }
    })
}

function resetInput() {
    $("#taskInput").val("")
    $("#addBtn").show()
    $("#saveEditBtn").hide()
    editingId = null
    oldTask = ""
}

function showNotification(message, type) {
    let color = type === "success" ? "green" : "red"
    let notification = $(`<div class="notification">${message}</div>`)
    
    notification.css({
        "background-color": color,
        "color": "white",
        "padding": "10px",
        "position": "fixed",
        "top": "20px",
        "right": "20px",
        "border-radius": "5px",
        "z-index": "1000",
        "display": "none"
    })

    $("body").append(notification)
    notification.fadeIn(300).delay(2000).fadeOut(300, function () {
        $(this).remove()
    })
}

function toggleComplete(id, completed) {
    let taskText = $(`#task-${id}`).text()
    $.ajax({
        url: `${API_URL}/${id}`,
        method: "PUT",
        contentType: "application/json",
        data: JSON.stringify({ task: taskText, completed: completed }),
        success: function () {
            fetchTasks()
        }
    })
}
