"use client";
import React from "react";
import { useRouter } from "next/navigation";

type CreateRoomResponse = {
  room_id: string;
};

const CreateRoom: React.FC = () => {
  const router = useRouter();
  const handleCreateRoom = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    const resp = await fetch("http://localhost:8000/create");
    const data: CreateRoomResponse = await resp.json();

    router.push(`/room/${data.room_id}`);
  };

  return (
    <div>
      <button onClick={handleCreateRoom}>Create Room</button>
    </div>
  );
};

export default CreateRoom;
