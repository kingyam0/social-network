let connect = cb => {
    console.log("connecting");
  
    socket.onopen = () => {
      console.log("Successfully Connected");
    };
  
    socket.onmessage = msg => {
      console.log(msg);
      cb(msg);
    };
  
    socket.onclose = event => {
      console.log("Socket Closed Connection: ", event);
    };
  
    socket.onerror = error => {
      console.log("Socket Error: ", error);
    };
  };
  