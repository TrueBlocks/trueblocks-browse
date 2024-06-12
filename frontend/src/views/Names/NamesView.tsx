import React, { useState, useEffect } from "react";
import { GetNames } from "@gocode/app/App";
import classes from "../View.module.css";
import View from "@/components/view/View";

function NamesView() {
  const [names, setName] = useState<string[]>();

  useEffect(() => {
    GetNames(0).then((names: string[]) => setName(names));
  }, []);

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
