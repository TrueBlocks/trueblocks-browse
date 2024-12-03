import { ReactNode } from "react";
import { Container, Fieldset, Grid, Accordion } from "@mantine/core";
import { IconChevronsUp } from "@tabler/icons-react";
import { FieldRenderer, FieldGroup, isCollapsable, isButton, ButtonTray } from "@components";
import { useViewState } from "@state";
import classes from "./FormTable.module.css";

type FormTableProps<T> = {
  data: Partial<T>;
  groups: FieldGroup<T>[];
};

export const FormTable = <T,>({ data, groups }: FormTableProps<T>) => {
  const { headerShows, handleCollapse, route } = useViewState();

  const collapsableGroups = groups.filter((group) => isCollapsable(group) && !isButton(group));
  const nonCollapsableGroups = groups.filter((group) => !isCollapsable(group));
  const buttonGroup = groups.find((group) => isButton(group)) || null;

  if (headerShows == null) {
    // avoids flashing
    return <></>;
  }

  // TODO: This is pretty dumb
  const style1 = {
    root: {
      paddingRight: "0px",
      paddingLeft: "0px",
      marginBottom: "-50px",
      marginTop: "-50px",
    },
  };
  const style2 = {
    root: {
      paddingRight: "12px",
      paddingLeft: "12px",
      marginTop: "40px",
      backgroundColor: "white",
    },
  };
  return (
    <Container styles={{ root: { minWidth: "100%" } }}>
      <Accordion
        classNames={{ chevron: classes.chevron }}
        data-rotate={headerShows ? "true" : "false"}
        styles={style1}
        value={headerShows ? "header" : null}
        onChange={(newState) => handleCollapse(route, newState)}
        chevron={null}
      >
        <Accordion.Item value="header">
          <CustomAccordionControl
            isOpen={headerShows[route]}
            onToggle={() => handleCollapse(route, headerShows[route] ? null : "header")}
          >
            <ButtonTray buttonGroup={buttonGroup} />
          </CustomAccordionControl>
          <Accordion.Panel>
            <Grid>
              {collapsableGroups.map((group, gIndex) => {
                return (
                  <Grid.Col key={group.label + gIndex} span={group.colSpan ?? 12}>
                    <Fieldset bg="white" className={classes.fieldSet}>
                      {group.fields?.map((fld, fIndex) => (
                        <FieldRenderer key={String(fld.accessor) + fIndex} field={fld} data={data} />
                      ))}
                      {group.components?.map((cmp, gIndex) => <div key={gIndex}>{cmp}</div>)}
                    </Fieldset>
                  </Grid.Col>
                );
              })}
            </Grid>
          </Accordion.Panel>
        </Accordion.Item>
      </Accordion>
      <Grid styles={style2}>
        {nonCollapsableGroups.map((group, gIndex) => (
          <Grid.Col key={group.label + gIndex} span={group.colSpan ?? 12}>
            <Fieldset legend={group.label} bg="white" className={classes.fieldSet}>
              {group.fields?.map((fld, fIndex) => (
                <FieldRenderer key={String(fld.accessor) + fIndex} field={fld} data={data} />
              ))}
              {group.components?.map((cmp, gIndex) => <div key={gIndex}>{cmp}</div>)}
            </Fieldset>
          </Grid.Col>
        ))}
      </Grid>
    </Container>
  );
};

type CustomAccordionControlProps = {
  isOpen: boolean;
  onToggle: () => void;
  children: ReactNode;
};

export const CustomAccordionControl = ({ isOpen, onToggle, children }: CustomAccordionControlProps) => {
  return (
    <div
      onClick={onToggle}
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "flex-end",
        gap: "8px",
        cursor: "pointer",
        padding: "10px",
      }}
      role="button"
      tabIndex={0}
    >
      {children}
      <IconChevronsUp
        className={`${classes.icon} ${classes.chevron} ${classes.buttonIcon}`}
        data-rotate={isOpen ? "true" : "false"}
        style={{
          paddingBottom: "2px",
        }}
      />
    </div>
  );
};
