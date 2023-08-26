function handleLogout(event) {
  event.preventDefault();
  $.ajax({
    url: "/loggout",
    method: "GET",
    success: function (response) {
      window.location.replace("/login");
    },
    error: function (xhr, status, error) {
      console.error(error);
    }
  });
}

$(document).ready(function () {

  $("#logout-btn").click(handleLogout);

});

