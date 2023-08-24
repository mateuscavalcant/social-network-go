$(document).ready(function() {
    // Function to handle the Follow button click event
    $("#follow-btn").click(function() {
      // Disable the button while the request is being processed
      $("#follow-btn").prop("disabled", true);
      var pathParts = window.location.pathname.split("/"); // Divide a URL em partes
      var user_follow_to = pathParts[pathParts.length - 1]; // O último elemento deve ser o nome de usuário
    
      // Perform AJAX request to follow the user
      $.ajax({
        type: "POST",
        url: "/follow", // Replace with the actual endpoint to perform the follow action
        data: {
          username: user_follow_to, // Replace with the user ID you want to follow
        },
        success: function(response) {
          // Change the button text and re-enable the button
          $("#follow-btn").text("Following").prop("disabled", false);
          console.log("Followed successfully:", response);
        },
        error: function() {
          // Re-enable the button in case of an error
          $("#follow-btn").prop("disabled", false);
          console.log("Error following user");
        }
      });
    });
  });

  