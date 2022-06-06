const aviso = document.querySelector(".errMessage")

const myTimeout = setTimeout(hide, 4000);

function hide() {
  aviso.style.display = "none";
}