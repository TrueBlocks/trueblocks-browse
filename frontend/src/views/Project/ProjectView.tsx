// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { useEffect } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm, DebugState } from "@components";
import { ModifyProject } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ProjectTableDefNoDelete, ProjectTableDef as tmpProjectTableDef, ProjectFormDef } from ".";
// EXISTING_CODE

export const ProjectView = () => {
  const { project, fetchProject, loadAddress } = useAppState();
  const handleEnter = (page: Page) => {
    loadAddress(project.items[page.getRecord()].address);
  };
  const handleModify = ModifyProject;

  // EXISTING_CODE
  const { info } = useAppState();
  useEffect(() => {
    fetchProject(0, 100);
  }, [info.filename, fetchProject]);
  let ProjectTableDef = tmpProjectTableDef;
  if (project?.nItems <= 2) {
    ProjectTableDef = ProjectTableDefNoDelete;
  }
  // EXISTING_CODE

  const table = useReactTable({
    data: project?.items || [],
    columns: ProjectTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "";
  const tabs = ["project"];
  const forms: ViewForm = {
    project: <FormTable data={project} groups={ProjectFormDef(table)} />,
  };

  // if (!(project?.items?.length > 0)) {
  //   return <>Loading...</>;
  // }

  return (
    <ViewStateProvider
      // do not remove - delint
      route={route}
      nItems={project.nItems}
      fetchFn={fetchProject}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState n={project.lastUpdate} />
      <View tabs={tabs} forms={forms} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
