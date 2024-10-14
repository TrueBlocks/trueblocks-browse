import { Routes } from ".";
import classes from "./ViewContainer.module.css";

export const ViewContainer = () => {
  return (
    <div className={classes.viewContainer}>
      <Routes />
    </div>
  );
};
