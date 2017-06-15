
$(".menu-bar").click(function() {
    $(".pagina").toggleClass("abrir");
    $(".pagina").toggleClass("cerrar");


})

$("li a").click(function() {
  $("a").removeClass("activo")
  $(this).toggleClass("activo")
})

$(".f").click(function() {
  $(this).toggleClass("abajo");
  $(this).toggleClass("arriba");
})

$(".submenu").click(function() {
  $(this).children("ul").slideToggle();
})


$("ul").click(function(p) {
  p.stopPropagation();
})
