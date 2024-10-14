import classes from "../components/view/View.module.css";
import { Routes } from ".";

export const ViewContainer = () => {
  return (
    <div className={classes.viewContainer}>
      <Routes />
    </div>
  );
};
