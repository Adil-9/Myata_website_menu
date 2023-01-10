function openMenu(event) {
    const tablinks = document.getElementsByClassName("tablink");
    for (var i = 0; i < tablinks.length; i++) {
        tablinks[i].classList.remove("myata_bg");
    }
    event.target.classList.add("myata_bg")
}
document.getElementById("myLink").click();

window.onscroll = function() {scrollFunction()};

function scrollFunction() {
    var menu = document.getElementById("menu")
    if (document.body.scrollTop > 750 || document.documentElement.scrollTop > 750) {
        menu.classList.add("top")
    } else {
        menu.classList.remove("top")
    }
}   