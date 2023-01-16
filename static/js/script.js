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

document.getElementById("cart-len").innerHTML = cart.length

function cartData(name, description, price, btn) {
    if(btn.className === "fa-solid fa-square-plus fa-2x") {
        cart.push(new item(name, description, price))
        btn.className = "fa-solid fa-square-minus fa-2x"
        btn.style.color = "red"
        document.getElementById("cart-len").innerHTML = cart.length
    } else {
        cart.pop()
        btn.className = "fa-solid fa-square-plus fa-2x"
        btn.style.color = "#61ede0"
        document.getElementById("cart-len").innerHTML = cart.length
    }
}

function loadCart() {
    
}
