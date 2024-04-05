const nameUser = prompt("Por favor, introduce tu nombre:");
let messageInput = document.getElementById("messageInput");
let messageList = document.getElementById("messageList");
let buttonSend = document.getElementById("sendMessage");
let socket = new WebSocket("ws://192.168.0.215:8000/room");

if (!nameUser) {
  alert("Por favor, introduce tu nombre");
  location.reload();
}

buttonSend?.addEventListener("click", () => {
  let message = messageInput.value;
  if (message) {
    let msg = {
      username: nameUser,
      text: message
    };
    socket.send(JSON.stringify(msg));
    messageInput.value = "";
  }
});

socket.onmessage = function(event) {
  var msg = JSON.parse(event.data);
  console.log("Received message:", msg);
  messageList.innerHTML += `<li>${msg.username}: ${msg.text}</li>`;
};

socket.onclose = function() {
  alert("Connection has been closed.");
};
