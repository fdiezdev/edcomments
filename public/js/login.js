let formLogin = document.getElementById("form-login");
let email = document.getElementById("email");
let pwd = document.getElementById("password");
let btnLogin = document.getElementById("btn-login");
let mensaje = document.getElementById("mensaje");

formLogin.addEventListener("submit", e => {
    e.preventDefault();
    
    email.classList.remove("is-valid");
    pwd.classList.remove("is-valid");
    email.classList.remove("is-invalid");
    pwd.classList.remove("is-invalid");
    mensaje.innerHTML = ""
    
    let obj = {
        email: email.value,
        password: pwd.value
    }
    
    ajaxRequest("POST", "http://localhost:8080/api/login/", JSON.stringify(obj), { mode: "no-cors" })
        .then(r => {
            if (r.status === 200) {
                email.classList.add("is-valid");
                pwd.classList.add("is-valid");
                
                sessionStorage.setItem('token', r.response.token)
            } else {
                console.log(r.response)
                email.classList.add("is-invalid");
                pwd.classList.add("is-invalid");
                mensaje.innerHTML = r.response.message;    
            }
        })
        .catch(error => {
            console.log("Error: "+error)
            console.log(error)
        })
});

