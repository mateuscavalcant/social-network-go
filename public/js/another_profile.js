function loadProfile(username) {
  $.ajax({
    url: "/profile/" + username,
    method: "GET",
    success: function (response) {
      var userHeaderDetailsHTML = '<div class="name">' +
        '<header>' +
        '<div class="user-name">' +
        '<p>' + response.profile.name + '</p>' +
        '</div>' +
        '<div class="posts-count">' + 
        '<p>' + response.profile.countposts + ' posts</p>' + 
        '</div>' + 
        '<main>' +
        '</main>' +
        '</div>';
      var $userHeaderDetails = $(userHeaderDetailsHTML);
      $("#profile-header-container").append($userHeaderDetails);
      if (response.profile.followby) {
        var userDetailsHTML = '<div class="user">' +
          '<header>' +
          '<img src="public/images/golang-icon2.jpeg" class="user-icon">' +
          '<div class="user-title">' +
          '<p>@' + response.profile.username + '</p>' +
          '</div>' +
          '<div class="user-bio">' +
          '<p>' + response.profile.bio + '</p>' +
          '</div>' +
          '</header>' +
          '<main>' +
          '<div class="user-followby">' +
          '<p>' + response.profile.followbycount + '</p>' +
          '<p id="followers-name">Followers</p>' +
          '</div>' +
          '<div class="user-followto">' +
          '<p>' + response.profile.followtocount + ' </p>' +
          '<p id="following-name">Following</p>' +
          '</div>' +
          '</main>' +
          '<footer>' +
          '<div class="create-btn">' +
          '<button id="following-btn">Following</button>' +
          '<button id="follow-btn" style="display: none;">Follow</button>' +
          '</div>' +
          '</footer>' +
          '</div>';
        var $userDetails = $(userDetailsHTML);
        $(document).ready(function () {
          $("#following-btn").click(function () {
            $("#following-btn").prop("disabled", true);
            var pathParts = window.location.pathname.split("/");
            var user_follow_to = pathParts[pathParts.length - 1]; 
            $.ajax({
              type: "POST",
              url: "/unfollow",
              data: {
                username: user_follow_to, 
              },
              success: function (response) {
                $("#following-btn").hide();
                $("#follow-btn").show();

                $("#following-btn").text("Follow").prop("disabled", false);
                console.log("Unfollowed successfully:", response);
              },
              error: function () {
                $("#follow-btn").prop("disabled", false);
                console.log("Error following user");
              }
            });
          });
        });
        $("#user-profile-container").append($userDetails);
      } else {
        var userDetailsHTML = '<div class="user">' +
          '<header>' +
          '<img src="public/images/golang-icon2.jpeg" class="user-icon">' +
          '<div class="user-title">' +
          '<p>@' + response.profile.username + '</p>' +
          '</div>' +
          '<div class="user-bio">' +
          '<p>' + response.profile.bio + '</p>' +
          '</div>' +
          '</header>' +
          '<main>' +
          '<div class="user-followby">' +
          '<p>' + response.profile.followbycount + '</p>' +
          '<p id="followers-name">Followers</p>' +
          '</div>' +
          '<div class="user-followto">' +
          '<p>' + response.profile.followtocount + ' </p>' +
          '<p id="following-name">Following</p>' +
          '</div>' +
          '</main>' +
          '<footer>' +
          '<div class="create-btn">' +
          '<button id="follow-btn">Follow</button>' +
          '<button id="following-btn" style="display: none;">Following</button>' +
          '</div>' +
          '</footer>' +
          '</div>';
        var $userDetails = $(userDetailsHTML);
        $(document).ready(function () {
          $("#follow-btn").click(function () {
            $("#follow-btn").prop("disabled", true);
            var pathParts = window.location.pathname.split("/"); 
            var user_follow_to = pathParts[pathParts.length - 1]; 

            $.ajax({
              type: "POST",
              url: "/follow", 
              data: {
                username: user_follow_to, 
              },
              success: function (response) {
                $("#follow-btn").hide();
                $("#following-btn").show();

                $("#follow-btn").text("Following").prop("disabled", false);
                console.log("Followed successfully:", response);
              },
              error: function () {
                $("#follow-btn").prop("disabled", false);
                console.log("Error following user");
              }
            });
          });
        });
        $("#user-profile-container").append($userDetails);
      }
      $("#posts-container").empty();
      response.posts.forEach(function (post) {
        var postHTML = '<div class="post">' +
          '<header>' +
          '<img src="public/images/golang-icon2.jpeg" class="profile-icon">' +
          '<div class="post-title">' +
          '<div class="user-name-post">' +
          '<p class="name-user' + post.postID + '">' + post.createdbyname + '</p>' +
          '</div>' +
          '<div class="user-username">' +
          '<p class="username-user' + post.postID + '">@' + post.createdby + '</p>' +
          '</div>' +
          '</div>' +
          '</header>' +
          '<main>' +
          '<div class="post-content">' +
          '<p>' + post.content + '</p>' +
          '</div>' +
          '<div class="post-links">' +
          '<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAACu0lEQVRYR+2WX27aQBDGbUA80xs4J2jyioC6N2huEE7Q5gRNThBuEG5QcoI6gHhNeoL6BuUZ8affF82iZdm114sVJVJWskD2/vnNN7MzE0dvbMRvjCd6/0BpmnY2m00KZb/gOY/juLPb7RJReonfXJ7HZrM5ybKM77yHt0K9Xi/F4T8JgafjfUIUTTD3YTabjX3WlAIRpNFo3EEFggQPGPO8Xq+Hi8XiuWiTQiDA3GCj7xZFMgA+brdbWh+12226iYPuJDhdSJd+Mw/Hutv5fH7jgnICCQxdpI8MB16XWakWdLvdc8QRDboy9hnBhdc2KCuQBaYSiHmQgNE4XbExoIbm3CMgCd5fmptGsPK26m2xWd/v9+/w/of6Bvd9hfsyfe4R0GAweNICeAIrLl3+DnkPKBqrlFpi/09OIFHnt0xYQpmzOpTRD5Q89iSBz0+XgHq5HBwHCoH+XgtAZ+CFKKOvMVx3EEsmENVJZfEB+akQhkoJbutfeZdDoTOXQpyU8KO4K68TRO0lbvunQkOPI1MhTnopC68IFAFoz+F0me1K1qWW5CUGdsSSMp1OL1wu2+eJshR/CpyReA9Si6kQ8wPzBAev/QWufa1xVOnak+IVEqOerXMxet8zuUqHSo5ktNacEJdZauRRarEWVyNB8uwMlgxD3Qc3Me+wuF5phlgTr7P9sEBxr0rZW0BU+6F3mc6CXdigmdVZs+5IagYrv69WqwRqpvj7GQ8vidnuFnYPpS0sc0ar1brXW1g9JUhBZg1MiuKK+Ybr9EJqm18KpBbpbYMCwjv2NoyNoqY/w/cHqDb26Ry8gYwbMhJFzJ55CSVyAPP5w57bt91VhocCmWqfdAv1zeoAqq3FJdgpQMyuw7IgrZpAQ4Fqc5EJHAJUq4uCgeSKs93cN+RV3eEz31shn83qmPMBVKbif1TPbjSKmlVJAAAAAElFTkSuQmCC" alt="Comentário" class="comment-button" data-post-id="' + post.postID + '">' +
          '<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAABrUlEQVRYR+2W7W2DMBCG+VggI7QbJH8RULJBu0GyQTNB0wmaDdoNmg3qCsTfZINkBBZA9D0JIst8HKYgEcn+e9z5ufde29jWzJY9Mx7LAHETMQoZhTgFuPj9eSgIgld0lcVx/KV2h1jBdazEheu6WyHEtS2vUyHP85Yo8IPkRZ7nqzRNz3KhAUCUvkNzB22gKIoI4oTEhzI5A9wjusuqYkOAiqJ4T5JkrwVUwnwi6VmVHN2tNcdk+b6/t237jfIGAaHzD+SSd5rWEVAvOlD/AipNTEBda9tk8raEwUCyiTkFmkw+OhDU2WDGlYktzP0Jm0TlRgKx32pTxOgqaD0tMtxghdQOdQp1KapTp/Me0ik0ayA6GBjvQhpv79FPohCUjeAxuuG5RZcsndjjDX4sqdU6UIku1W+GqPaMTKJQBQEouu03LVAHKLNTY5MC0WZhGJ7gp6W8McZ5dhxnLb+L2iPjzCDH0fmt0fJdvCBemZwe6VXbL0inQj2fkRqrDERByeQ1E2uNrOEXpJdQKhAl0StAKnG3+/39wvaSZMSPjEKcmEYhoxCnABefnYf+AGy7CDSivNhoAAAAAElFTkSuQmCC" alt="Comment">' +
          '<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAACUElEQVRYR+2X3W3CMBDHiWAAugFsAK8IUNigbAAblAlaJigbNBuUDRoB4hU2KCOwAKL3j86Vc3Fix84DD1jqQ+Pz3e8+baLWg63owXhaTyBbRpwjFMdx93a7xVDYbrfPaZpeqpSTfI/kB/f7/drpdCB/tcFgvxKIId5J7pX+egaFCcGtlDFNfkGyXSEPoK0ubwIsBRqPx3EURV8lIFLXkj98GkCkLMBW+/0+cQYimA+CQWT0BUUpGxw4Gj6zApM8oDYSqhCh0Wg0oLD+KIMEdqY6WNPhrX6Y5b5lBOvIU40Nj8ejgs7U54C4BmAkZuPI+bKqICeTCdK6YPkNya8t8tCPmsS6kKN93dEcECmHIA5gXcmDmfRAhhj/87kuwWxt3cRO/2opz6UuByRqZ0P0KxNA6DdhJyE7qinyKSNP9XDOZd2EgqjzeiZQc7vdbqj2ZMoQyh42Kfx92/DzBeShCVtZaZDjL2VA6K4Ym9RZs8PhkPoarTrHHXqCjC1CGGxvrMw4J5oApJTBBmxh5Wq1sssobcOm02bosqU+tQuDcTqdnihdmKxYWxKeNxEVpUM0zoWd/r94C0Ci4AohDYETQ9RYp8bLlQ7qtdQIlISRtWPsMt17VwUuETPpKrtiKt9DTUDVgcnGgM3DEKi6ME5AEPKB8oFxBqoL5QtTC8gVKgSmNpANKhTGC6gMir7jV8ZCaxLr69HUUNYuK+tCQzR0US8Y7whp95L+nlafvWGCgQzpC4JpBIihsjcUXQeJ7ZFvG8TeNWRT7Lv/BLJF7uEi9AesB1Q04shzcgAAAABJRU5ErkJggg==" alt="Curtir" class="like-button" data-post-id="' + post.postID + '">' +
          '</div>' +
          '</main>' +
          '<footer>' +
          '</footer>' +
          '</div>';
        var $post = $(postHTML);
        $post.find(".like-button").on("click", function () {
          var postID = $(this).data("post-id");
          // Alterar a imagem para a nova imagem de curtida
          $(this).attr("src", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAAAXNSR0IArs4c6QAAAnBJREFUWEft1t+rDGEYB/Dvd3Zn2yNhZ6ToJEWhzp2kKMyk5IILbtwoodw4F4oSLpQokRJulKSUC8UFyc3Z2Vw5f8CpLTdbpyS1s3I4aubsPNpTK2fNj3f2naOjzl7O+7zP83mf931nh1hiPy4xD5ZBWTui1CEB2KnsHRMp7zKAEQqb3bmwaeP9dFwBH3s2GuXqVqFsi4CfZPShFjSmCIgWaAa714Vm9Z4AR0isGEwmIpNGKMdraLR6Y1/hbu6aeE5iZ0zsDME35ZDjqzDRToIldsgvuYdh4AkIO21VIvIdxGlDUI2AhyRXpsfjCyOcsbr113FxsaBO2b0qBq5ntVdrPIouWXONW4M5/gL5JfcQSnirVUxxshHJgTVz3sSf4QtAgh2mX1ndIrBBMadmmLRqwdotxItuP9ECUMd0Tgn5WLNKzunRCStoPIsFtSvOO4IHc2bUChfglR3UjyaBpgmOalXIOVkgTTvwtseDTGeW5EjOnFrhAvlsB976pA59Ivh7UKuS6mSRKSv0xmJBvuk0QO5TzVVEnEBe2oF3LP6W/YsX4sAqROSsHXqP4rcM+0dh8iPJahGrz8wh0pFwdpONyW+xoN5Dv+xchsEbmckKCKBEJ2th42nim7o/0DaduyTPF1AzJYXctALvymBA8r+96d4HcW5RUIIHVlgfj8ud+oHmLwYqBdMDZn4xForKwCiB5g96EZ1SwCiDtFGKmFwgDdQdK6hfVL0cmWdoMJFfcW8DuKBYIBcmd4f6CEVUbszQoPntS+/UUBgtUApqaIw2KAalhSkE1EcJ8MMO6tcUD3tiWO5bplswa/4y6L/r0C89QuAlNQSzNwAAAABJRU5ErkJggg==");
        });
        $("#posts-container").prepend($post);
      });
    },
    error: function (xhr, status, error) {
      console.error(error);
    }
  });
}

function handleHome(event) {
  event.preventDefault();

  window.location.replace("/home");
}

$(document).ready(function () {
  var pathParts = window.location.pathname.split("/"); 
  var username = pathParts[pathParts.length - 1]; 
  loadProfile(username);

  $("#home-btn").click(handleHome);

  $("#follow-btn").click(followUser(username))


});