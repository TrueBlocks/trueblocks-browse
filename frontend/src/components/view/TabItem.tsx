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
  const { route, activeTab, headerOn, headerOnChanged } = useAppState();

  const collapsableGroups = useMemo(() => groups.filter((group) => isCollapsable(group) && !isButton(group)), [groups]);
  const nonCollapsableGroups = useMemo(() => groups.filter((group) => !isCollapsable(group)), [groups]);
  const buttonGroup = useMemo(() => groups.find((group) => isButton(group)) || null, [groups]);

  const renderGroups = (groups: FieldGroup<T>[], data: Partial<T>, withLegend: boolean = false) => {
    return groups.map((group, gIndex) => (
      <Grid.Col key={group.label + gIndex} span={group.colSpan ?? 12}>
        <Fieldset legend={withLegend ? group.label : undefined} bg="white" className={classes.fieldSet}>
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
        key={`${route}-${activeTab}`}
        classNames={{
          root: classes.rootStyle,
        }}
        data-rotate={headerOn ? "true" : "false"}
        value={headerOn ? "header" : null}
        onChange={(newState) => {
          headerOnChanged(newState === "header");
        }}
      >
        <Accordion.Item value="header">
          <CustomAccordionControl>
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
  children: ReactNode;
};

export const CustomAccordionControl = ({ children }: CustomAccordionControlProps) => {
  const { headerOn, headerOnChanged } = useAppState();

  return (
    <div onClick={() => headerOnChanged(!headerOn)} className={classes.controlContainer} role="button" tabIndex={0}>
      {children}
      <div className={classes.chevronWrapper}>
        <IconChevronsUp className={`${classes.chevronIcon}`} data-rotate={headerOn ? "true" : "false"} />
      </div>
    </div>
  );
};
