import { BaseButton, ButtonProps } from "@components";
import classes from "./BaseButton.module.css";

// CloseButton is used currently only for the Help window to close it if it's open.
export const CloseButton = ({ onClose }: ButtonProps) => {
  return (
    <BaseButton tip="Close" onClick={onClose} className={classes.closeButton}>
      X
    </BaseButton>
  );
};
