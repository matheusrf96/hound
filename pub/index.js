function getInputValue(inputName) {return document.getElementById(inputName) ? document.getElementById(inputName).value : ''}

function getData() {
    return {
        'name': getInputValue('name'),
        'email': getInputValue('email'),
        'birth_date':getInputValue('birth_date'),
    }
}

function sendMessage() {
    const data = getData()
    const socket = new WebSocket(`ws://localhost:8000/ws`)

    socket.onopen = () => {
        socket.send(JSON.stringify(data))
    }
}
