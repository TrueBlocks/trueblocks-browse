import { Fieldset } from "@mantine/core";
import classes from "./FieldsetWrapper.module.css";

interface FieldsetWrapperProps {
  legend: string;
  children: React.ReactNode;
}

export const FieldsetWrapper = ({ legend, children }: FieldsetWrapperProps) => (
  <Fieldset bg="white" className={classes.fieldSetWrapper} legend={legend}>
    {children}
  </Fieldset>
);
