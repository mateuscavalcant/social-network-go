document.getElementById("login-form").addEventListener("submit", function(event) {
    event.preventDefault();
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
        
    var formData = new FormData();
    formData.append("email", email);
    formData.append("password", password);
        
    fetch("/login", {
        method: "POST",
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            // Exibe as mensagens de erro correspondentes nos campos do formulário
            document.getElementById("error-email").textContent = data.error.email;
            document.getElementById("error-password").textContent = data.error.password;
        } else {
            console.log(data.message);
            
            // Redirecionamento manual para /create-post
            window.location.href = "http://localhost:8080/home";
        }
    })
    .catch(error => {
        console.error(error);
    });
});