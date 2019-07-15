$(function () {
  var socket = null;
  var $msgBox = $("#chatbox textarea");


  if (!window["WebSocket"]) {
    alert("エラー: WebSocketに対応していないブラウザです。")
  } else {
    socket = new WebSocket("ws://{{.Host}}/room");

    socket.onclose = function () {
      alert("接続が終了しました。");
    }

    socket.onmessage = function (e) {
      $("#messages").append($("<li>").text(e.data));
    }
  }


  $msgBox.on("input", function (e) {
    var value = $msgBox.val();
    socket.send($msgBox.val());
  })



/*
  $("#chatbox").submit(function () {
    if (!$msgBox.val()) return false;
    if (!socket) {
      alert("エラー: WebSocket接続が行われていません。");
      return false;
    }

    socket.send($msgBox.val());
    $msgBox.val("");
    return false;
  });
*/

});
