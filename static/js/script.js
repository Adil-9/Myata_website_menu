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
document.getElementById("myLink").click();

// window.onscroll = function() {scrollFunction()};

// function scrollFunction() {
//     var menu = document.getElementById("menu")
//     var empty = document.getElementById("empty")
//     if (document.body.scrollTop > screen.height*1.235 || document.documentElement.scrollTop > screen.height*1.235) {
//         menu.classList.add("top")
//         empty.style.display = "block"
//     } else {
//         menu.classList.remove("top")
//         empty.style.display = "none"
//     }
// }   