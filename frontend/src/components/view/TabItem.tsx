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

  const renderGroups = (groups: FieldGroup<T>[], data: Partial<T>, withLegend: boolean = false) => {
    return groups.map((group, gIndex) => (
      <Grid.Col key={group.label + gIndex} span={group.colSpan ?? 12}>
        <Fieldset
          legend={withLegend ? group.label : undefined} // Add legend only if specified
          bg="white"
          className={classes.fieldSet}
        >
          {group.fields?.map((fld, fIndex) => (
            <FieldRenderer key={String(fld.accessor) + fIndex} field={fld} data={data} />
          ))}
          {group.components?.map((cmp, cmpIndex) => <div key={cmpIndex}>{cmp}</div>)}
        </Fieldset>
      </Grid.Col>
    ));
  };

  return (
    <Container styles={{ root: { minWidth: "100%" } }}>
      <Accordion
        classNames={{
          chevron: classes.chevron,
          root: classes.rootStyle,
          control: classes.controlStyle,
        }}
        data-rotate={headerOn ? "true" : "false"}
        value={headerOn ? "header" : null}
        onChange={(newState) => headerOnChanged(newState === "header")}
        chevron={null}
      >
        <Accordion.Item value="header">
          <CustomAccordionControl isOpen={headerOn} onToggle={() => headerOnChanged(!headerOn)}>
            <ButtonTray buttonGroup={buttonGroup} />
          </CustomAccordionControl>
          <Accordion.Panel>
            <Grid>{renderGroups(collapsableGroups, data)}</Grid>
          </Accordion.Panel>
        </Accordion.Item>
      </Accordion>
      <Grid className={classes.groupStyles}>{renderGroups(nonCollapsableGroups, data, true)}</Grid>
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
