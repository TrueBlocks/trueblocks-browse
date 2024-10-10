import { BaseButton, ButtonProps } from "@components";
import classes from "./BaseButton.module.css";

export function CloseButton({ onClose }: ButtonProps) {
  return (
    <BaseButton tip="Close" onClick={onClose} className={classes.closeButton}>
      X
    </BaseButton>
  );
}
