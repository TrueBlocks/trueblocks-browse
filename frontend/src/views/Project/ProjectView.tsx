import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, GroupDefinition } from "@components";
import { SetSessionVal } from "@gocode/app/App";
import { types, messages } from "@gocode/models";
import { Page } from "@hooks";
import { EventsEmit } from "@runtime";
import { useAppState, ViewStateProvider } from "@state";
import { tableColumns } from "./ProjectTable";

export function ProjectView() {
  const { project, fetchProject } = useAppState();

  const handleEnter = (page: Page) => {
    const record = page.selected - page.getOffset();
    const address = project.items[record].address;
    SetSessionVal("route", `/history/${address}`);
    EventsEmit(messages.Message.NAVIGATE, {
      route: `/history/${address}`,
    });
  };

  //   if (!address) {
  //     return <Text>Address not found</Text>;
  //   }

  const table = useReactTable({
    data: project.items || [],
    columns: tableColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ViewStateProvider route={""} nItems={project.myCount} fetchFn={fetchProject} onEnter={handleEnter}>
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
      colSpan: 5,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nMonitors" },
        { label: "nNames", type: "int", accessor: "nNames" },
        { label: "nAbis", type: "int", accessor: "nAbis" },
      ],
    },
    {
      title: "Data 2",
      colSpan: 5,
      fields: [
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "nManifests", type: "int", accessor: "nManifests" },
        { label: "nCaches", type: "int", accessor: "nCaches" },
        { label: "historySize", type: "bytes", accessor: "historySize" },
        { label: "nOpen", type: "int", accessor: "myCount" },
      ],
    },
  ];
}
