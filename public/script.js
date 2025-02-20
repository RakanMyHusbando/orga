const loginP = document.querySelector("header nav a.login p");

window.onload = () => {
    if (
        document.cookie.includes("session_token") &&
        document.cookie.includes("csrf_token")
    )
        loginP.innerHTML = "logout";
    else loginP.innerHTML = "login";
};
