const { app, ipcMain, dialog } = require('electron').remote
const open_btn = document.getElementById("open_btn")
const open_xlsx_btn = document.getElementById("open_xlsx_btn")
const selector_value = document.getElementById("projectsSelector")

document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        console.log("message received", message)
    })
})

open_xlsx_btn.addEventListener('click', ()=> {
    astilectron.sendMessage("File" + dialog.showOpenDialogSync({ properties: ['openFile'] })[0])
})

open_btn.addEventListener('click', ()=> {
    astilectron.sendMessage("Project" + dialog.showOpenDialogSync({ properties: ['openFile'] })[0])
})
