import React, { useState, useEffect } from "react";
import { GetBlock } from "../wailsjs/go/app/App";
import { app } from "../wailsjs/go/models";
import "./App.css";

function App() {
  const [block, setBlock] = useState<app.Block>();
  const [curBlock, setCurBlock] = useState<number>(1000);

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
