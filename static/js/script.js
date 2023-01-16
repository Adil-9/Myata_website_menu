document.getElementById("myLink").click();

function openMenu(event, menuName) {
    const tablinks = document.getElementsByClassName("tablink");
    for (var i = 0; i < tablinks.length; i++) {
        tablinks[i].classList.remove("myata_bg");
    }
    event.target.classList.add("myata_bg")
    const menus = document.getElementsByClassName("menu")
    for (var i = 0; i < menus.length; i++) {
        menus[i].style.display = "none"
    }
    document.getElementById(menuName).style.display = "block"
}

class item {
    constructor(name, description, price) {
        this.name = name
        this.description = description
        this.price = price
    }
}

var cart = []

var target = document.getElementById("cart-content")

var template = "<div><div style=\"margin: 0px 5px;\" class=\"d-flex justify-content-between\"><h1>Sample name</h1><h1>3000 тг</h1></div><div style=\"margin: 0px 5px;\" class=\"d-flex justify-content-between\"><p>Lorem ipsum dolor sit amet consectetur adipisicing elit.</p><div class=\"d-flex justify-content-center\"><i style=\"padding-top: 15px; color: red;\" class=\"fa-solid fa-square-minus fa-xl\"></i><text style=\"margin: 0 10px;\"> 0 </text><i style=\"padding-top: 15px; color: #61ede0;\" class=\"fa-solid fa-square-plus fa-xl\"></i></div></div></div>";

function cartData(name, description, price, btn) {
    if(btn.className === "fa-solid fa-square-plus fa-2x") {
        cart.push(new item(name, description, price))
        btn.className = "fa-solid fa-square-minus fa-2x"
        btn.style.color = "red"
        document.getElementById("cart-len").innerHTML = cart.length
        target.insertAdjacentHTML("beforeend", template)
    } else {
        cart.pop()
        btn.className = "fa-solid fa-square-plus fa-2x"
        btn.style.color = "#61ede0"
        document.getElementById("cart-len").innerHTML = cart.length
        target.removeChild(target.lastChild)
    }
}

function openCart() {
    for (let element of document.getElementsByClassName("to-hide")){
        element.style.display = "none";
    }
    document.getElementById("cart-container").style.display = "block"
    window.scrollTo(0, 0);
}

function closeCart() {
    for (let element of document.getElementsByClassName("to-hide")){
        element.style.display = "block";
    }
    document.getElementById("cart-container").style.display = "none"
}