import React, { useState, useEffect } from "react";
import { GetNames } from "@gocode/app/App";
// import { app } from "@gocode/models";
import { useHotkeys } from "react-hotkeys-hook";
import classes from "../Blocks/BlocksView.module.css";
import View from "@/components/view/View";

function NamesView() {
  const [block, setBlock] = useState<string[]>();
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
    setCurBlock(curBlock + 1 > 19100000 ? 19100000 : curBlock + 1);
  });
  useHotkeys("down", (event) => {
    event.preventDefault();
    setCurBlock(curBlock + 1000 > 19100000 ? 19100000 : curBlock + 1000);
  });
  useHotkeys("end", (event) => {
    event.preventDefault();
    setCurBlock(19000000);
  });

  const page = 0;
  useEffect(() => {
    GetNames(page).then((val: string[]) => setBlock(val));
  }, [curBlock]);

  return (
    <View title="BlocksView">
      <section>
        <button
          className={classes.btn}
          onClick={() => {
            setCurBlock(curBlock - 1);
          }}
        >
          {"< prev"}
        </button>
        <button
          className={classes.btn}
          onClick={() => {
            setCurBlock(curBlock + 1);
          }}
        >
          {"next >"}
        </button>
        <div id="result" className={classes.result}>
          <pre>{JSON.stringify(block, null, 4)}</pre>
        </div>
      </section>
    </View>
  );
}

export default NamesView;
