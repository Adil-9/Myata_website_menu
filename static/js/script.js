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
    constructor(id, name, description, price, count) {
        this.id = id
        this.name = name
        this.description = description
        this.price = price
    }
}

var id = 1

var cart = []

var target = document.getElementById("cart-content")

var template = "<div id=\"SampleId\"><div style=\"margin: 0px 5px;\" class=\"d-flex justify-content-between\"><h1> SampleName </h1><h1> SamplePrice </h1></div><div style=\"margin: 0px 5px;\" class=\"d-flex justify-content-between\"><p> SampleDescription </p><div class=\"d-flex justify-content-center\"><i onclick=\"decrementCount(\'CountId\', \'SampleId\')\" style=\"padding-top: 15px; color: red;\" class=\"fa-solid fa-square-minus fa-xl\"></i><text id = \"CountId\" style=\"margin: 0 10px;\"> 1 </text><i onclick=\"incrementCount(\'CountId\', \'SampleId\')\" style=\"padding-top: 15px; color: #61ede0;\" class=\"fa-solid fa-square-plus fa-xl\"></i></div></div></div>";

var cid = "a"

var costTotal = 0

function cartData(id, name, description, price, btn) {
    if(btn.className === "fa-solid fa-square-plus fa-2x") {
        cart.push(new item(id, name, description, price, 1))
        btn.className = "fa-solid fa-square-minus fa-2x"
        btn.style.color = "red"
        document.getElementById("cart-len").innerHTML = cart.length
        var t = template.replace("SampleName", name).replace("SamplePrice", price + " тг").replaceAll("SampleId", id).replaceAll("CountId", cid)
        target.insertAdjacentHTML("beforeend", t.replace("SampleDescription", description))
        cid += "a"
        for (let i = 0; i < cart.length; i++) {
            if (cart[i].id == id) {
                costTotal += Number(cart[i].price)
                break
            }
        }
    } else {
        for (let i = 0; i < cart.length; i++) {
            if (cart[i].id == id) {
                costTotal -= Number(cart[i].price)
                cart.splice(i, 1)
                break
            }
        }
        btn.className = "fa-solid fa-square-plus fa-2x"
        btn.style.color = "#61ede0"
        document.getElementById("cart-len").innerHTML = cart.length
        document.getElementById(id).remove()
    }
    document.getElementById("cost-total").innerHTML = "Total: " + costTotal + " тг"
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

function incrementCount(cid, id) {
    var num = Number(document.getElementById(cid).innerHTML) + 1
    document.getElementById(cid).innerHTML = num
    for (let i = 0; i < cart.length; i++) {
        if (cart[i].id == id) {
            costTotal += Number(cart[i].price)
            break
        }
    }
    document.getElementById("cost-total").innerHTML = "Total: " + costTotal + " тг"
}

function decrementCount(cid, id) {
    var num = Number(document.getElementById(cid).innerHTML) - 1
    if (num == 0) {
        document.getElementById("cart-len").innerHTML = cart.length
        document.getElementById(id).remove()
        for (let i = 0; i < cart.length; i++) {
            if (cart[i].id == id) {
                costTotal -= Number(cart[i].price)
                cart.splice(i, 1)
                break
            }
        }
        document.getElementById("cart-len").innerHTML = cart.length
        document.getElementById("cost-total").innerHTML = "Total: " + costTotal + " тг"
        return
    }
    document.getElementById(cid).innerHTML = num
    for (let i = 0; i < cart.length; i++) {
        if (cart[i].id == id) {
            costTotal -= Number(cart[i].price)
            break
        }
    }
    document.getElementById("cost-total").innerHTML = "Total: " + costTotal + " тг"
}