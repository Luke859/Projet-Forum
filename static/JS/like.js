let like = document.getElementById('addlike')

function changeLike() {

    like.textContent = "Liked post ❤️" + 1

    if(like.textContent === "Liked post ❤️" + 1){
        like.onclick() = deleteLike;
    }
    
    like.submit()
}

function deleteLike(){

    like.textContent = "Disliked post 💔" + 0

    if(like.textContent === "Disliked post 💔" + 0){
        like.onclick() = changelike();
    }
    
    like.submit()
}
