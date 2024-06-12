import React, { useState, useEffect } from "react";
import { GetNames } from "@gocode/app/App";
import { useHotkeys } from "react-hotkeys-hook";
import classes from "../View.module.css";
import View from "@/components/view/View";

function NamesView() {
  const [names, setName] = useState<string[]>();
  const [curName, setCurName] = useState<number>(0);

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

  useEffect(() => {
    console.log("useEffect", curName);
    GetNames(curName).then((names: string[]) => setName(names));
  }, [curName]);

  return (
    <View title="Names View">
      <section>
        <div id="result" className={classes.result}>
          <pre>{JSON.stringify(names, null, 4)}</pre>
        </div>
      </section>
    </View>
  );
}

export default NamesView;
