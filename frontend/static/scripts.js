document.addEventListener('DOMContentLoaded', function () {
    const userForm = document.getElementById('userForm');
    const userTableBody = document.getElementById('userTableBody');

    let users = [];

    function renderUsers() {
        userTableBody.innerHTML = '';
        users.forEach((user, index) => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${user.name}</td>
                <td>${user.username}</td>
                <td>
                    <button class="edit" onclick="editUser(${index})">EDIT</button>
                    <button class="delete" onclick="deleteUser(${index})">DELETE</button>
                </td>
            `;
            userTableBody.appendChild(row);
        });
    }

    userForm.addEventListener('submit', function (e) {
        e.preventDefault();
        const name = document.getElementById('name').value;
        const username = document.getElementById('username').value;
        users.push({ name, username });
        renderUsers();
        userForm.reset();
    });

    window.editUser = function (index) {
        const user = users[index];
        document.getElementById('name').value = user.name;
        document.getElementById('username').value = user.username;
        users.splice(index, 1); // Remove user being edited
        renderUsers();
    };

    window.deleteUser = function (index) {
        users.splice(index, 1);
        renderUsers();
    };

    renderUsers();
});
