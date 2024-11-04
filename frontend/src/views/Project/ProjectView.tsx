import { useEffect } from "react";
import { Text } from "@mantine/core";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { GoToAddress, ModifyHistory } from "@gocode/app/App";
import { Page, useRenderCounter } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ProjectTableDefNoDelete, ProjectTableDef, ProjectFormDef } from ".";

export const ProjectView = () => {
  const { project, fetchProject, info } = useAppState();
  const renderCount = useRenderCounter();

  useEffect(() => {
    fetchProject(0, 100);
  }, [info.filename, fetchProject]);

  const handleEnter = (page: Page) => {
    if (project && project.items) {
      const history = project.items[page.getRecord()];
      if (history && history.address) {
        GoToAddress(history.address).then(() => {});
      }
    }
  };

  const projColumns = project?.nItems < 2 ? ProjectTableDefNoDelete : ProjectTableDef;
  const table = useReactTable({
    data: project.items ?? [],
    columns: projColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "";
  const tabs = ["project"];
  const forms: ViewForm = {
    project: <FormTable data={project} groups={ProjectFormDef(table)} />,
  };

  if (project?.items?.length === 0) {
    return <></>;
  }

  return (
    <ViewStateProvider
      route={route}
      nItems={project.nItems}
      fetchFn={fetchProject}
      onEnter={handleEnter}
      modifyFn={ModifyHistory}
    >
      <Text>Render count: {renderCount}</Text>
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
