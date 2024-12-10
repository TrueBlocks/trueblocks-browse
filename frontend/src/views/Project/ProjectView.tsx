// This file is auto-generated. Edit only code inside
// of ExistingCode markers (if any).

// EXISTING_CODE
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, TabItem, ViewForm, DebugState } from "@components";
import { ModifyProject } from "@gocode/app/App";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { ProjectFormDef, ProjectTableDef } from ".";
// EXISTING_CODE

export const ProjectView = () => {
  const { project, fetchProject, loadAddress } = useAppState();
  const handleEnter = (page: Page) => {
    loadAddress(project.items[page.getRecord()].address);
  };
  const handleModify = ModifyProject;

  const table = useReactTable({
    data: project?.items || [],
    columns: ProjectTableDef,
    getCoreRowModel: getCoreRowModel(),
  });

  const tabItems: ViewForm = {
    project: <TabItem tabName="project" data={project} groups={ProjectFormDef(table)} />,
  };

  return (
    <ViewStateProvider
      // do not remove - delint
      nItems={project.nItems}
      fetchFn={fetchProject}
      onEnter={handleEnter}
      modifyFn={handleModify}
    >
      <DebugState u={[project.updater]} />
      <View tabItems={tabItems} searchable />
    </ViewStateProvider>
  );
};

// EXISTING_CODE
// EXISTING_CODE
