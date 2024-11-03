import { Button, ButtonProps as MantineButtonProps, ActionIcon } from "@mantine/core";
import { base } from "@gocode/models";
import classes from "./BaseButton.module.css";

export type ButtonMouseEvent = React.MouseEvent<HTMLButtonElement>;

export interface ButtonProps extends MantineButtonProps {
  tip?: string;
  value?: string | base.Address;
  icon?: React.ReactNode;
  onClick?: (e: ButtonMouseEvent) => void;
  onClose?: (e: ButtonMouseEvent) => void;
}

// BaseButton is a generic button that can have a loading spinner, a tip, and a
// left section. It can also be used as an action icon if it has no children.
export const BaseButton = ({
  loading = false,
  tip = "",
  onClick = () => {},
  onClose = () => {},
  children,
  ...props
}: ButtonProps) => {
  const { icon } = props;
  const baseProps: MantineButtonProps = { ...props };
  const hasChilren: boolean = children !== undefined;

  const handleClick = (e: ButtonMouseEvent) => {
    if (onClick) {
      onClick(e);
    }
    if (onClose) {
      onClose(e);
    }
  };

  if (!hasChilren) {
    return (
      <ActionIcon className={classes.actionButton} onClick={handleClick} title={tip}>
        {icon}
      </ActionIcon>
    );
  }

  return (
    <Button className={classes.baseButton} onClick={handleClick} loading={loading} size="xs" {...baseProps}>
      {children}
    </Button>
  );
};
