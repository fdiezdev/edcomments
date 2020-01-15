let formCommentAdd = document.getElementById("form-comment");
let commentContent = document.getElementById("comment-content");
let pwd = document.getElementById("btn-comment");

formCommentAdd.addEventListener("submit", e => {
    e.preventDefault();
    
    let comment = {
        content: commentContent.nodeValue,
    }
    
    ajaxRequest("POST", "https://localhost:8080/api/coments", JSON.stringify(obj))
        .then(r => {
            if (r.status === 200) {
                console.log("Comentario creado")
            } else {
                console.log(r.response)
            }
        })
        .catch(error => {
            console.log(error)
        })
});