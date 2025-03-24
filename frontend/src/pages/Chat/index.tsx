const Chat = () => {
  const roomCode = sessionStorage.getItem("roomCode");
  return (
    <div className="p-8 text-center text-xl font-bold space-y-4">
      <div>
        Welcome to <span className="text-blue-500">Chat</span>
      </div>
      <div>
        Room code: <span className="text-blue-500"> {roomCode}</span>
      </div>
    </div>
  );
};

export default Chat;
