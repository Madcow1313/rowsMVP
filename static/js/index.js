const { app, ipcMain } = require('electron')
const open_btn = document.getElementById("open_btn")
const selector_value = document.getElementById("projectsSelector")

document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        console.log("message received", message)
    })
    open_btn.addEventListener('click', ()=> {
        let str = selector_value.value
        astilectron.sendMessage(str)
    })
})
