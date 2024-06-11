import React, { useState, useEffect } from "react";
import { GetNames } from "@gocode/app/App";
// import { app } from "@gocode/models";
import { useHotkeys } from "react-hotkeys-hook";
import classes from "../Blocks/BlocksView.module.css";
import View from "@/components/view/View";

function NamesView() {
  const [names, setName] = useState<string[]>();
  const [curName, setCurName] = useState<number>(0);

  useHotkeys("home", (event) => {
    event.preventDefault();
    setCurName(0);
    console.log("home", curName);
  });
  useHotkeys("left", (event) => {
    event.preventDefault();
    setCurName(curName - 1 < 0 ? 0 : curName - 1);
    console.log("left");
    console.log("left", curName);
  });
  useHotkeys("up", (event) => {
    event.preventDefault();
    setCurName(curName - 20 < 0 ? 0 : curName - 20);
    console.log("up", curName);
  });
  useHotkeys("right", (event) => {
    event.preventDefault();
    setCurName(curName + 1 > 19100000 ? 19100000 : curName + 1);
    console.log("right", curName);
  });
  useHotkeys("down", (event) => {
    event.preventDefault();
    setCurName(curName + 20 > 19100000 ? 19100000 : curName + 20);
    console.log("down", curName);
  });
  useHotkeys("end", (event) => {
    event.preventDefault();
    setCurName(12000);
    console.log("end", curName);
  });

  useEffect(() => {
    console.log("useEffect", curName);
    GetNames(curName).then((names: string[]) => setName(names));
  }, [curName]);

  return (
    <View title="Names View">
      <section>
        <button
          className={classes.btn}
          onClick={() => {
            setCurName(curName - 1);
            console.log("prev", curName);
          }}
        >
          {"< prev"}
        </button>
        <button
          className={classes.btn}
          onClick={() => {
            setCurName(curName + 1);
            console.log("next", curName);
          }}
        >
          {"next >"}
        </button>
        <div id="result" className={classes.result}>
          <pre>{JSON.stringify(names, null, 4)}</pre>
        </div>
      </section>
    </View>
  );
}

export default NamesView;