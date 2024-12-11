import { ReactNode, useMemo } from "react";
import { Container, Fieldset, Grid, Accordion } from "@mantine/core";
import { IconChevronsUp } from "@tabler/icons-react";
import { FieldRenderer, FieldGroup, isCollapsable, isButton, ButtonTray } from "@components";
import { useAppState } from "@state";
import classes from "./TabItem.module.css";

type TabItemProps<T> = {
  data: Partial<T>;
  groups: FieldGroup<T>[];
};

export const TabItem = <T,>({ data, groups }: TabItemProps<T>) => {
  const { headerOn, headerOnChanged } = useAppState();

  const collapsableGroups = useMemo(() => groups.filter((group) => isCollapsable(group) && !isButton(group)), [groups]);
  const nonCollapsableGroups = useMemo(() => groups.filter((group) => !isCollapsable(group)), [groups]);
  const buttonGroup = useMemo(() => groups.find((group) => isButton(group)) || null, [groups]);

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
        data-rotate={headerOn ? "true" : "false"}
        styles={style1}
        value={headerOn ? "header" : null}
        onChange={(newState) => headerOnChanged(newState === "header")}
        chevron={null}
      >
        <Accordion.Item value="header">
          <CustomAccordionControl isOpen={headerOn} onToggle={() => headerOnChanged(!headerOn)}>
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
