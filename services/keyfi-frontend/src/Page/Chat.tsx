import React, { useState, useEffect } from 'react';
import './Chat.css'
import { QueryServiceClient } from '../../protos/query/query.client';
import { GetValuesRequest } from '../../protos/query/query';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';

const Chat: React.FC = () => {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  const [inputValue, setInputValue] = useState<string>('');

  const makeCallToBackend = async () => {
    let transport = new GrpcWebFetchTransport({
      baseUrl: "http://ec2-34-236-81-43.compute-1.amazonaws.com:8080"
    });
    const client = new QueryServiceClient(transport);
    const request = GetValuesRequest.create({
      keys: ["yaobin", "foo"]
    })
    const call = client.getValues(request);
    let response = await call.response;
    let status = await call.status;
    console.log("status: " + status)
    console.log(response.keyValuePairs);
  }

  useEffect(() => {
    
    makeCallToBackend();

    // Create a new WebSocket connection when the component mounts
    const newWs = new WebSocket('ws://ec2-34-236-81-43.compute-1.amazonaws.com:8081');

    newWs.onopen = () => {
      console.log('WebSocket connected');
    };

    newWs.onmessage = (event) => {
      // Add received message to the messages state
      setMessages((prevMessages) => [...prevMessages, event.data]);
    };

    newWs.onclose = () => {
      console.log('WebSocket disconnected');
    };

    // Update ws state with the new WebSocket connection
    setWs(newWs);

    // Close the WebSocket connection when the component unmounts
    return () => {
      newWs.close();
    };
  }, []);

  const sendMessage = () => {
    if (ws && inputValue.trim() !== '') {
      // Send the message through the WebSocket connection
      ws.send(inputValue);
      setMessages((prevMessages) => [...prevMessages, inputValue]);
      setInputValue('');
    }
  };

  return (
    <div>
      <h2>Chat</h2>
      <div>
        {messages.map((message, index) => (
          <div key={index}>{message}</div>
        ))}
      </div>
      <input
        type="text"
        value={inputValue}
        onChange={(e) => setInputValue(e.target.value)}
        placeholder="Type a message..."
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
};

export default Chat;