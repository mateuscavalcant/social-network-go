document.addEventListener("DOMContentLoaded", function() {
    function validateEmail() {
        var email = document.getElementById("email").value;
        var formData = new FormData();

        formData.append("email", email);
        
        fetch("/validate-email", {
            method: "POST",
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                // Exibe as mensagens de erro correspondentes nos campos do formulário
                document.getElementById("error-email").textContent = data.error.email;
            } else {
                document.getElementById("error-email").textContent = "";
            }
        })
        .catch(function(error) {
            console.error(error);
        });
    }
    document.getElementById("email").addEventListener("input", function() {
        validateEmail();
    });
});

        // Ouvinte de evento para enviar o formulário de cadastro
        document.getElementById("signup-form").addEventListener("submit", function(event) {
        event.preventDefault();
        var username= document.getElementById("username").value;
        var email = document.getElementById("email").value;
        var password = document.getElementById("password").value;
        var confirmPassword = document.getElementById("confirm_password").value;
            
        var formData = new FormData();
        formData.append("username", username);
        formData.append("email", email);
        formData.append("password", password);
        formData.append("confirm_password", confirmPassword);
        
            
        fetch("/signup", {
            method: "POST",
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                // Exibe as mensagens de erro correspondentes nos campos do formulário
                document.getElementById("error-username").textContent = data.error.username;
                document.getElementById("error-email").textContent = data.error.email;
                document.getElementById("error-password").textContent = data.error.password;
                document.getElementById("error-confirm-password").textContent = data.error.confirm_password;
            } else {
                console.log(data.message);
                
                // Redirecionamento manual para /create-post
                window.location.href = "/login";
            }
        })
        .catch(error => {
            console.error(error);
        });
    });     