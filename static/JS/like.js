function changeLike(id) {

    let like = document.getElementById(id)

    like.textContent = "Liked post ❤️" + 1

    if(like.textContent === "Liked post ❤️" + 1){
        like.onclick = deleteLike;
    }
}

function deleteLike(id){

    let like = document.getElementById(id)

    like.textContent = "Disliked post 💔" + 0

    if(like.textContent === "Disliked post 💔" + 0){
        like.onclick = changeLike;
    }
}