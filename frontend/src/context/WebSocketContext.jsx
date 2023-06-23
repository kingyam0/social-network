import React, { createContext, useEffect, useContext, useState, useRef } from 'react';

const WebSocketContext = createContext();

const WebSocketProvider = ({ children }) => {

    const [isConnected, setConnected] = useState(false);

    const ws = useRef(null);

    const openWebSocketConnection = () => {

        ws.current = new WebSocket('ws://localhost:9090/ws');

        ws.current.onopen = () => {

            setConnected(true);
    };

    ws.current.onclose = () => {

        setConnected(false);
      console.log("WebSocket:", isConnected);
      console.log('WebSocket connection is closed.'); // Log closure here

    };

    // Handle other WebSocket events and logic here
  };

  const sendWebSocketMessage = (message) => {

    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(message);
    }
  };

  const closeWebSocketConnection = () => {

    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.close();
    }
  };

  useEffect(() => {

    openWebSocketConnection();

    return () => {

        closeWebSocketConnection();
    };
  }, []);

  return (
    <WebSocketContext.Provider value={{ isConnected, sendWebSocketMessage, closeWebSocketConnection }}>
      {children}
    </WebSocketContext.Provider>
  );
};

const useWebSocket = () => useContext(WebSocketContext);

export { WebSocketProvider, useWebSocket };