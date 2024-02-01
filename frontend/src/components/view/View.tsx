import React, { ReactNode } from "react";
import { Title } from "@mantine/core";
import classes from "./View.module.css";

function View(params: { title: string; children: ReactNode }) {
  return (
    <section>
      <Title order={1}>{params.title}</Title>

      <div className={classes.content}>{params.children}</div>
    </section>
  );
}

export default View;
