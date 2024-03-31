let formulario = document.getElementById("chatbox");
let messages = document.getElementById("messages");
let username = document.getElementById("username");
let text = document.getElementById("message");
let buttonConect = document.getElementById("conect");
let buttonSend = document.getElementById("send");
let socket = new WebSocket("ws://localhost:8000/room");

formulario?.addEventListener("submit", (e) => {
  e.preventDefault();
  if (!socket) {
    alert("Error: There is no socket connection.");
    return false;
  }
  let msg = {
    username: username.value,
    text: text.value
  };
  socket.send(JSON.stringify(msg));
  console.log(msg);
  console.log("Mensaje enviado");
  text.value = "";
  return false;
});

socket.onmessage = function(event) {
  var msg = JSON.parse(event.data);
  console.log("Received message:", msg);
  messages.innerHTML += `<li>${msg.username}: ${msg.text}</li>`;
};

socket.onclose = function() {
  alert("Connection has been closed.");
};
