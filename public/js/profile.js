
    function loadPostsProfile() {
        $.ajax({
          url: "/profile",
          method: "GET",
          success: function(response) {
            
              // Exibir os detalhes do perfil do usuário
              var userDetailsHTML = '<div class="user">' +
              '<header>' +
              
              '<img src="public/images/user-icon.jpg" class="user-icon">' +
              '<div class="user-title">' +
              '<p>@' + response.profile.username + '</p>' +
              '</div>' +
              '</header>' +
              '</div>';

          var $userDetails = $(userDetailsHTML);
          $("#user-profile-container").append($userDetails);
          $("#posts-container").empty();
      
            // Iterar sobre os posts retornados e adicionar na página
            response.posts.forEach(function(post) {
              var postHTML = '<div class="post">' +
                '<header>' +
                '<div class="post-title">' +'<div class="post-title">' +
                '<img src="public/images/user-icon.jpg" class="profile-icon">' +
                
                '<p id="username-' + post.postid + '" class="post-username">@' + post.createdby + '</p>' +
                '</div>' +
                '</header>' +
                '<main>' +
                '<div class="post-content">' +
                '<p>' + post.content + '</p>' +
                '</div>' +
                '<div class="post-links">' +
                '<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAAAXNSR0IArs4c6QAAAwtJREFUWEfl11vopWMUx/HPOAwhOYTCZJJxJW6c0pSc4sYxUhrkVCSlSCTG4YYiNy4YIzFyiBJ3zi4c5mZKuJCGxEgYDQk5jn5j7X9vr/3u/e7Ze9c0frVrv/t5nvV+n7WeZ621F9nGtGgb47FdAO2P03EKjsB+9Ymzv8O3+Aiv4ZX6rXcg+npoF1yAK3Fib+tsxht4BC/g93Fr+wCdisdx4DhjY8Y/xwq8M2reKKAd8CCuaRn4s8LxHtbiC2yoOQfjEByPEyqsO7bW34eb/Ou9/6gLaGc8h7MbK37ESjyBTT29tU+F+Xbs3lizBpfhr7adLqDAnN+Y/Fa5+6ueIO1pS/Esjm0MxPvX9QG6pM7MYO5juGrYbiaEW4yncV5j3UnIZhfU9tBuiBf2qhkxcNGELx53Zl/GaTUp5y/eWzhPbaDL8WhN/hqH4ZcZAsXUAfgM2Xx0BgK5RW2g13Fyjd2JO2YMMzCXvJScFj2Ji7uAPsWhNXgUPpgTUEKWLB6tw9FdQD9hjxrcE3meh1J+vinDP2DvLqCfG7ENWJ7noeSn7/sAfYJlNfFIfDgPGhxXWX5syJ7BhQVxM+6dE1DKxw1l+2Fc3RWy5rXfWN5KjGep5Lj12LeMpualLm7RsMT4JRLj6KEhxXVauJdwZhnJkcjRWNCwWpackAI60I24f1oK7ITVuLRsJTsfU9d+JFAGs/CKBkR2dW2jzZiULy3JU9WSDNbegnvahrqqfXqY53FOY8FveKAOet9zNWg/bkXy2kBpR+4etqtRDVpcnDPU9FRs/FFtQ25HU3nhQVUs02YsR6p5s0H7FddjVZeL+7SwZ1XnuKRhpFnn8tJ0BSmao/Ru1awU1k71AcriXavBClwUoLuQUOT7KDvvV5F+sc/B6wsUW6n8aWGjhPLwRmeQ33JrkrvyNygN/Zt4ddICvbVA7c2+XS3voGD2ccbQOdMC/V23JeHL96k1DVDCcy7inZlpa4FmFqL2TiYFuq1uV5LaTEI0DVBahI/bf1tmFqsyNImHZv3uqW/Z/xPoH95OjSUy92gTAAAAAElFTkSuQmCC alt="Comentário" class="comment-button" data-post-id="' + post.postID + '">' +
                '<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAAAXNSR0IArs4c6QAAAX5JREFUWEft1LFKHFEUxvHfFmIvWOgrqE0KsTCJD2GVRpskBNKkSNokZcA2gtpoIT5DKo2GkKRJEfIGEgMBX0AlcmAWlmFn79xZRlaY2wzLuee7//t9d0/PhK3ehPHogFKJdA51DqUcSNXv5Rt6Wtxqb8jt/qduXKp/xzrOq/pSDq3grGh+iG8loVygaH+NrSZAs/iN+Mb6h4Xi29drAvQe73KBpvAFy6XGH1jFVWZUAfC26GkEtINnFYfu4vldAm1gP3HgJg4yoBo7FBFFVBHZqBWRRXQRYZ3VGChimh84YQ2Pi9+fcTJQ+4OIr85qDFQWry2UoKqtk5pDtYUmHahx9G05FBP+tMmfoy2gSPAFtnPHR5tAwXKIJxVQH/GyXGsbaBpf8aB0cMy6GCk34wDVmTf9PYMXjbn2CzNF8QKLuBwmmHLoDT7kkBR7y7qPcIxrxIP/WaWZAlrCJ8xlQg3TfYW/OBqllQLK5Bh/eweU8rBzqHMo5UCq3r2hlEO36T9QJaLRJvIAAAAASUVORK5CYII=" alt="Comment">' +
                '<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAAAXNSR0IArs4c6QAAAsVJREFUWEft1zvoFUcUx/HPXxE02gmJiiii4IMUQQiWFimioBYq+EDFJGohiCZaaAxYGDVBfIIIim9FEFEhAa0kpRZ2ivGBSkJAAlYJCYLRcP7Myrrs3t37+MO/8MDlwsxvZr5zzsyZs30GmfUNMh7vgeoi0tRDofsYszACv6bf7xULTMBUTMO/uIV7eNMt0Ic4hAX4oGSy21iKZ6lvMi7i0xLtX/gZG/CiCqyVh+bjFEbX7OpvfIXhOIJRNfo/sQY/lemqgL7DztyA17iEB6ltCpZhSMXiRX2EbwnvXKKt+LE4vgxoDq7nhNewEb8VBsc5CY/MK7RX6Scl/dyc/jPczI8vAg3DI0xMotP4okUIYnycsTgXYQfxdY3+LFYkTaw1Hf9lY4pAX+JE6nyaxC/rbgbWJc2xBtq4pQ8xPmlX4nwV0A18njoDLg71QFh49HCa+CoWVgFFXsnIx+L5QNCkHBW5LOw+ZlQB/ZMSX/QPRdyWgbCPcpuNTcfm+614hv7AuNQXB7t4s3oFNxN30mSRweMVKAX6BbNT3yqc6xVBYZ4t2JvarmBRFdB2fJ867+KT/JXsEdxIPMaYNF/c0ONVQBHLJ+kZCE0IsyvdI55+r2d5KN60SJjxzpWGLBq/xa7c6nXJrh3QA9iUG7AaZ/ITVL1lxYH7sbmdlUu0+/BNrn034oi8Y61e+8i6a3vkqeIGj2J92QbrCrReQDWGCcA6oNB0A1WEOZlqp8roNwEqg/oB22rO1B5EzZNZwERh1rKMbQrULlRHME1DlndEMXxlnirCXECUGLUFfidAdZ4qg4knqPEj3U7IWnlqR6qv4z+z8ExbMJ16KFuwGL488OVU1Df2TDa4Uw+1ggqY+FZ7WyfX3MbGmbrpPPHWxVdJWHwqLe8UptuQ5YEXIz6Lok5+1XQnZbpuQ9bN2qVj3wPVuXTQeeh/xW2FJWAig/cAAAAASUVORK5CYII=" alt="Curtir" class="like-button" data-post-id="' + post.postID + '">' +
                '</div>' +
                '</main>' +
                '<footer>' +
                '</footer>' +
                '</div>';
      
              var $post = $(postHTML);
              $post.find(".like-button").on("click", function() {
                var postID = $(this).data("post-id");
                // Alterar a imagem para a nova imagem de curtida
                $(this).attr("src", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAAAXNSR0IArs4c6QAAAnBJREFUWEft1t+rDGEYB/Dvd3Zn2yNhZ6ToJEWhzp2kKMyk5IILbtwoodw4F4oSLpQokRJulKSUC8UFyc3Z2Vw5f8CpLTdbpyS1s3I4aubsPNpTK2fNj3f2naOjzl7O+7zP83mf931nh1hiPy4xD5ZBWTui1CEB2KnsHRMp7zKAEQqb3bmwaeP9dFwBH3s2GuXqVqFsi4CfZPShFjSmCIgWaAa714Vm9Z4AR0isGEwmIpNGKMdraLR6Y1/hbu6aeE5iZ0zsDME35ZDjqzDRToIldsgvuYdh4AkIO21VIvIdxGlDUI2AhyRXpsfjCyOcsbr113FxsaBO2b0qBq5ntVdrPIouWXONW4M5/gL5JfcQSnirVUxxshHJgTVz3sSf4QtAgh2mX1ndIrBBMadmmLRqwdotxItuP9ECUMd0Tgn5WLNKzunRCStoPIsFtSvOO4IHc2bUChfglR3UjyaBpgmOalXIOVkgTTvwtseDTGeW5EjOnFrhAvlsB976pA59Ivh7UKuS6mSRKSv0xmJBvuk0QO5TzVVEnEBe2oF3LP6W/YsX4sAqROSsHXqP4rcM+0dh8iPJahGrz8wh0pFwdpONyW+xoN5Dv+xchsEbmckKCKBEJ2th42nim7o/0DaduyTPF1AzJYXctALvymBA8r+96d4HcW5RUIIHVlgfj8ud+oHmLwYqBdMDZn4xForKwCiB5g96EZ1SwCiDtFGKmFwgDdQdK6hfVL0cmWdoMJFfcW8DuKBYIBcmd4f6CEVUbszQoPntS+/UUBgtUApqaIw2KAalhSkE1EcJ8MMO6tcUD3tiWO5bplswa/4y6L/r0C89QuAlNQSzNwAAAABJRU5ErkJggg==");
      
              });
      
              $("#posts-container").prepend($post);
            });
          },
          error: function(xhr, status, error) {
            console.error(error);
          }
        });
      }

      function handleHome(event) {
        event.preventDefault();
      
            window.location.replace("/home");
      }
     
          $(document).ready(function() {
            
            
            loadPostsProfile();

            $("#home-btn").click(handleHome);
      
          });