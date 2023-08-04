import 'whatwg-fetch';

document.getElementById("signup-form").addEventListener("submit", function (event) {
    event.preventDefault();

    var username = document.getElementById("username").value;
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
            }
        })
        .catch(error => {
            console.error(error);
        });
});

function validateEmail() {
    var emailInput = document.getElementById("email");
    var emailError = document.getElementById("error-email");
    // Limpa a mensagem de erro atual
    emailError.textContent = "";
    // Verifica se o campo de e-mail está vazio
    if (emailInput.value.trim() === "") {
        return;
    }
    // Cria um objeto FormData e adiciona o valor do campo de e-mail
    var formData = new FormData();
    formData.append("email", emailInput.value);
    // Envia uma solicitação POST assíncrona para a rota /validate-email
    fetch("/validate-email", {
        method: "POST",
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                // Exibe a mensagem de erro retornada pelo back-end
                emailError.textContent = data.error;
            }
        })
        .catch(error => {
            console.error(error);
        });
}



