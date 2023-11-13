const addForm = document.getElementById("addForm");
const titleInput = document.getElementById("title");
const goalsList = document.getElementById("goalsList");

addForm.addEventListener("submit", function (e) {
  e.preventDefault();

  const title = titleInput.value;

  fetch("/add", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title }),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log("Success:", data);
      titleInput.value = "";
      fetchGoals();
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});

function fetchGoals() {
  fetch("/goals")
    .then((response) => response.json())
    .then((data) => {
      goalsList.innerHTML = "";
      data.forEach((goal) => {
        const li = document.createElement("li");
        li.innerHTML = `${goal.title} <button onclick="updateGoal(${goal.id})">Update</button> <button onclick="deleteGoal(${goal.id})">Delete</button>`;
        goalsList.appendChild(li);
      });
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}

function updateGoal(id) {
  const updatedTitle = prompt("Enter updated title:");
  if (updatedTitle !== null) {
    fetch("/update", {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id, title: updatedTitle }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Success:", data);
        fetchGoals();
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }
}

function deleteGoal(id) {
  fetch("/delete", {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(id),
  })
    .then((response) => {
      if (response.ok) {
        console.log("Success: Goal deleted");
        fetchGoals();
      } else {
        console.error("Error:", response.statusText);
      }
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}

fetchGoals();
