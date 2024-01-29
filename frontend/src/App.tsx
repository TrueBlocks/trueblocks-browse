import React, { useState } from "react";
import { Greet } from "../wailsjs/go/app/App";
import "./App.css";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇"
  );
  const [name] = useState("Thomas");
  const updateResultText = (result: string) => setResultText(result);
  function greet() {
    Greet(name).then(updateResultText);
  }

  return (
    <div id="App">
      <div id="result" className="result">
        {resultText}
      </div>
      <button className="btn" onClick={greet}>
        Greet
      </button>
    </div>
  );
}

export default App;
