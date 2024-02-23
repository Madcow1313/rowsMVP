let addBtn = document.getElementById("add_btn")

addBtn.addEventListener('click', ()=> {
    astilectron.sendMessage("add")
})