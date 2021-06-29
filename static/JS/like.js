let like = document.getElementById('addlike')

function changeLike() {

    like.textContent = 1

    if(like.textContent === 1){
        like.onclick() = deleteLike;
    }
    
    like.submit()
}

function deleteLike(){

    like.textContent = 0

    if(like.textContent === 0){
        like.onclick() = changelike();
    }
    
    like.submit()
}
