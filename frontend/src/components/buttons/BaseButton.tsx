import { Button, ButtonProps as MantineButtonProps, ActionIcon } from "@mantine/core";
import { base } from "@gocode/models";
import classes from "./BaseButton.module.css";

export interface ButtonProps extends MantineButtonProps {
  tip?: string;
  value?: string | base.Address;
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
  onClose?: (e: React.MouseEvent<HTMLButtonElement>) => void;
}

export const BaseButton = ({
  loading = false,
  tip = "",
  onClick = () => {},
  onClose = () => {},
  children,
  ...props
}: ButtonProps) => {
  const { leftSection } = props;
  const baseProps: MantineButtonProps = { ...props };
  const hasChilren: boolean = children !== undefined;

  const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
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
        {leftSection}
      </ActionIcon>
    );
  }

  return (
    <Button className={classes.baseButton} onClick={handleClick} loading={loading} size="xs" {...baseProps}>
      {children}
    </Button>
  );
};
