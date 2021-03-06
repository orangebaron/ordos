const serverUrl = "http://192.168.1.14"
var myName = "uwe" // change later

function addToChat(speaker,msg) {
  document.getElementById("chatbar").innerHTML += '<p class="chat"><b class="' +
    (speaker == myName ? "" : "not") + 'me">'+speaker+'</b>: '+msg+'</p>'
}

function eventRecieved(data,status) {
  $.get(serverUrl+":8081/event",eventRecieved) //server waits until there's new info to give then gives it

  if (data.substring(0,4) == "CHAT") {
    data = data.substring(4)
    colon = data.search(":")
    if (colon == -1) {
      return
    }
    addToChat(data.substring(0,colon),data.substring(colon+2))
  }
}

$(document).ready(function() {
  document.getElementById("chatbox").onkeyup = function(key) {
    if (key.key=="Enter") {
      var str = document.getElementById("chatbox").value
      document.getElementById("chatbox").value = ""
      str = str.substring(0,str.length-1)
      if (str!="") {
        $.post(serverUrl+":8081/chat",str)
      }
    }
  }

  $.get(serverUrl+":8081/event",eventRecieved)
})
