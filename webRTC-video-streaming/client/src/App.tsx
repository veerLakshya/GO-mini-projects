import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import CreateRoom from "./components/CreateRoom";
import Room from "./components/Room";

function App() {
  return (
    <div className="App">
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="/" Component={CreateRoom} />
            <Route path="/room" Component={Room} />
          </Routes>
        </BrowserRouter>
      </div>
    </div>
  );
}

export default App;
