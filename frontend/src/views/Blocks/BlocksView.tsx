import React, { useState, useEffect } from "react";
import { GetBlock } from "@gocode/app/App";
import { app } from "@gocode/models";
import classes from "../View.module.css";
import View from "@/components/view/View";

function BlocksView() {
  const [block, setBlock] = useState<app.Block>();
  useEffect(() => {
    GetBlock(12).then((val: app.Block) => setBlock(val));
  }, []);

  return (
    <View title="BlocksView">
      <section>
        <div id="result" className={classes.result}>
          <pre>{JSON.stringify(block, null, 4)}</pre>
        </div>
      </section>
    </View>
  );
}

export default BlocksView;
