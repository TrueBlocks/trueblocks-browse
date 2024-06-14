import React, { useState, useEffect } from "react";
import { GetNames, MaxNames } from "@gocode/app/App";
import { useHotkeys } from "react-hotkeys-hook";
import classes from "../View.module.css";
import View from "@/components/view/View";

function NamesView() {
  const [names, setName] = useState<string[]>();
  const [curName, setCurName] = useState<number>(0);
  const [maxNames, setMaxNames] = useState<number>(0);

  useHotkeys("left", (event) => {
    event.preventDefault();
    setCurName(curName - 1 < 0 ? 0 : curName - 1);
  });
  useHotkeys("up", (event) => {
    event.preventDefault();
    setCurName(curName - 20 < 0 ? 0 : curName - 20);
  });
  useHotkeys("right", (event) => {
    event.preventDefault();
    setCurName(curName + 1 > maxNames ? maxNames : curName + 1);
  });
  useHotkeys("down", (event) => {
    event.preventDefault();
    setCurName(curName + 20 > maxNames ? maxNames - 20 : curName + 20);
  });
  useHotkeys("home", (event) => {
    event.preventDefault();
    setCurName(0);
  });
  useHotkeys("end", (event) => {
    event.preventDefault();
    setCurName(maxNames - 20);
  });

  useEffect(() => {
    console.log("useEffect", curName);
    GetNames(curName, 20).then((names: string[]) => setName(names));
    MaxNames().then((maxNames: number) => setMaxNames(maxNames));
  }, [curName]);

  return (
    <View title="Names View">
      <section>
        <div id="result" className={classes.result}>
          <pre>Number of records: {maxNames}</pre>
          <pre>{JSON.stringify(names, null, 4)}</pre>
        </div>
      </section>
    </View>
  );
}

export default NamesView;
