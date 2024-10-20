import { Container, Fieldset, Grid, Accordion } from "@mantine/core";
import { IconPlus } from "@tabler/icons-react";
import { FieldRenderer, FieldGroup, isCollapsable, isButton, ButtonTray } from "@components";
import { useViewState } from "@state";
import classes from "./FormTable.module.css";

type FormTableProps<T> = {
  data: Partial<T>;
  groups: FieldGroup<T>[];
};

export const FormTable = <T,>({ data, groups }: FormTableProps<T>) => {
  const { headerShows, handleCollapse } = useViewState();

  const collapsableGroups = groups.filter((group) => isCollapsable(group) && !isButton(group));
  const nonCollapsableGroups = groups.filter((group) => !isCollapsable(group));
  const buttonGroup = groups.find((group) => isButton(group)) || null;

  if (headerShows == null) {
    // avoids flashing
    return <></>;
  }

  return (
    <Container styles={{ root: { minWidth: "95%" } }}>
      <Accordion
        classNames={{ chevron: classes.chevron }}
        data-rotate={headerShows ? "true" : "false"}
        styles={{ root: { marginBottom: "-20px", marginTop: "-15px" } }}
        value={headerShows ? "header" : null}
        onChange={handleCollapse}
        chevron={<IconPlus className={classes.icon} />}
      >
        <Accordion.Item value="header">
          <Accordion.Control c={"black"} bg="white">
            <ButtonTray buttonGroup={buttonGroup} />
          </Accordion.Control>
          <Accordion.Panel>
            <Grid>
              {collapsableGroups.map((group) => {
                return (
                  <Grid.Col key={group.label} span={group.colSpan ?? 12}>
                    <Fieldset bg="white" className={classes.fieldSet}>
                      {group.fields?.map((fld) => <FieldRenderer key={String(fld.accessor)} field={fld} data={data} />)}
                      {group.components?.map((cmp) => cmp)}
                    </Fieldset>
                  </Grid.Col>
                );
              })}
            </Grid>
          </Accordion.Panel>
        </Accordion.Item>
      </Accordion>
      <Grid>
        {nonCollapsableGroups.map((group) => (
          <Grid.Col key={group.label} span={group.colSpan ?? 12}>
            <Fieldset legend={group.label} bg="white" className={classes.fieldSet}>
              {group.fields?.map((fld) => <FieldRenderer key={String(fld.accessor)} field={fld} data={data} />)}
              {group.components?.map((cmp) => cmp)}
            </Fieldset>
          </Grid.Col>
        ))}
      </Grid>
    </Container>
  );
};
