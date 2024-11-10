import { useEffect } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { GoToAddress, ModifyProject } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ProjectTableDefNoDelete, ProjectTableDef, ProjectFormDef } from ".";

export const ProjectView = () => {
  const { info } = useAppState();
  const { project, fetchProject } = useAppState();

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
      modifyFn={ModifyProject}
    >
      <DebugState n={project.lastUpdate} />
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};
