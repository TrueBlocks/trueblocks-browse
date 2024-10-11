import { Routes } from "../global";
import classes from "./View.module.css";

export const ViewContainer = () => {
  return (
    <div className={classes.viewContainer}>
      <Routes />
    </div>
  );
};
