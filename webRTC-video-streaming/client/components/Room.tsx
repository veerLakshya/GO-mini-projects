"use client";
import { useParams } from "next/navigation";
import React, { useEffect, useRef } from "react";

const Room = () => {
  const params = useParams();
  const { id } = params;
  console.log(id);

  useEffect(() => {
    if (!id) return;

    const ws = new WebSocket(`ws://localhost:8000/join?roomID=${id}`);
    ws.addEventListener("open", () => {
      ws.send(JSON.stringify({ join: "true" }));
    });
  }, [id]);

  return (
    <div>
      {id}
      <video autoPlay controls></video>
      <video autoPlay controls></video>
    </div>
  );
};

export default Room;
