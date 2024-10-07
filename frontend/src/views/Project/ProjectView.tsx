import { useState } from "react";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { SetSessionVal, ModifyProject } from "@gocode/app/App";
import { types, messages } from "@gocode/models";
import { Page } from "@hooks";
import { EventsEmit } from "@runtime";
import { useAppState, ViewStateProvider } from "@state";
import { withoutDelete, withDelete } from "./ProjectTable";

export function ProjectView() {
  const { project, fetchProject } = useAppState();
  // const [filtered, setFiltered] = useState<types.HistoryContainer[]>([]);

  const handleEnter = (page: Page) => {
    const record = page.selected - page.getOffset();
    const address = project.items[record].address;
    SetSessionVal("route", `/history/${address}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${address}`,
    });
  };

  const modColumns = project.items?.length < 2 ? withoutDelete : withDelete;
  const table = useReactTable({
    data: project.items ?? [],
    columns: modColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider
      route={""}
      nItems={project.nOpenFiles}
      fetchFn={fetchProject}
      onEnter={handleEnter}
      modifyFn={ModifyProject}
    >
      <View>
        <FormTable data={project} definition={createProjectForm(table)} />
      </View>
    </ViewStateProvider>
  );
}

type theInstance = InstanceType<typeof types.ProjectContainer>;
function createProjectForm(table: any): GroupDefinition<theInstance>[] {
  return [
    {
      title: "Open Monitors",
      fields: [],
      components: [
        {
          component: <DataTable<types.HistoryContainer> table={table} loading={false} />,
        },
      ],
    },
    {
      title: "Data 1",
      colSpan: 6,
      fields: [
        { label: "fileName", type: "text", accessor: "filename" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "nManifests", type: "int", accessor: "nManifests" },
        { label: "nCaches", type: "int", accessor: "nCaches" },
        { label: "historySize", type: "bytes", accessor: "historySize" },
        { label: "nOpen", type: "int", accessor: "nOpenFiles" },
      ],
    },
    {
      title: "Data 2",
      colSpan: 3,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nMonitors" },
        { label: "nNames", type: "int", accessor: "nNames" },
        { label: "nAbis", type: "int", accessor: "nAbis" },
      ],
    },
  ];
}
