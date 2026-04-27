import React from "react";

const CreateRoom = (props) => {
  const handleCreateRoom = async (e: React.FormEvent) => {
    // Logic to create a room
    e.preventDefault();

    const resp = await fetch("http://localhost:8000/create");
    const { room_id } = await resp.json();

    props.history.push(`/room/${room_id}`); // Navigate to the newly created room
  };

  return (
    <div>
      <button onClick={handleCreateRoom}> Create Room</button>
    </div>
  );
};

export default CreateRoom;
