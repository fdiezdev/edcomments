
/* Peticiones AJAX */
function ajaxRequest(method, url, obj) {
    
    return new Promise(function(resolve, reject) {
        let xhr = new XMLHttpRequest();
        xhr.open(method, url, true);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.addEventListener("load", e => {
            
            let self = e.target;
            let response = {
                status: self.status,
                response: JSON.parse(self.response)
            }
            
            resolve(response);
            
        });
        xhr.addEventListener("error", e => {
            
            let self = e.target;
            reject(self);
            
        });
        xhr.send(obj);       
    });
    
}

function $(elemento) {
    return document.getElementById(elemento);
}