import React, { useState, useEffect } from "react";
import { useHotkeys } from "react-hotkeys-hook";
import { GetBlock } from "../wailsjs/go/app/App";
import { app } from "../wailsjs/go/models";
import "./App.css";

function App() {
  const [block, setBlock] = useState<app.Block>();
  const [curBlock, setCurBlock] = useState<number>(1000);

  useHotkeys("home", (event) => {
    event.preventDefault();
    setCurBlock(0);
  });
  useHotkeys("left", (event) => {
    event.preventDefault();
    setCurBlock(curBlock - 1 < 0 ? 0 : curBlock - 1);
  });
  useHotkeys("up", (event) => {
    event.preventDefault();
    setCurBlock(curBlock - 1000 < 0 ? 0 : curBlock - 1000);
  });
  useHotkeys("right", (event) => {
    event.preventDefault();
    setCurBlock(curBlock + 1 > 19100000 ? 0 : curBlock + 1);
  });
  useHotkeys("down", (event) => {
    event.preventDefault();
    setCurBlock(curBlock + 1000 > 19100000 ? 0 : curBlock + 1000);
  });
  useHotkeys("end", (event) => {
    event.preventDefault();
    setCurBlock(19000000);
  });

  useEffect(() => {
    GetBlock(curBlock).then((val: app.Block) => setBlock(val));
  }, [curBlock]);

  return (
    <div id="App">
      <button
        className="btn"
        onClick={() => {
          setCurBlock(curBlock - 1);
        }}
      >
        {"< prev"}
      </button>
      <button
        className="btn"
        onClick={() => {
          setCurBlock(curBlock + 1);
        }}
      >
        {"next >"}
      </button>
      <div id="result" className="result">
        <pre>{JSON.stringify(block, null, 4)}</pre>
      </div>
    </div>
  );
}

export default App;
