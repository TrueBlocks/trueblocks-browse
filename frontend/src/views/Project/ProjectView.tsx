import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, FieldGroup } from "@components";
import { GoToHistory, ModifyProject } from "@gocode/app/App";
import { types } from "@gocode/models";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { withoutDelete, withDelete } from "./ProjectTable";

export const ProjectView = () => {
  const { project, fetchProject } = useAppState();
  // const [filtered, setFiltered] = useState<types.HistoryContainer[]>([]);

  const handleEnter = (page: Page) => {
    const address = project.items[page.getRecord()].address;
    GoToHistory(address).then(() => {});
  };

  const modColumns = project.nOpenFiles < 2 ? withoutDelete : withDelete;
  const table = useReactTable({
    data: project.items ?? [],
    columns: modColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "";
  return (
    <ViewStateProvider
      route={route}
      nItems={project.nOpenFiles}
      fetchFn={fetchProject}
      onEnter={handleEnter}
      modifyFn={ModifyProject}
    >
      <View>
        <FormTable data={project} groups={createProjectForm(table)} />
      </View>
    </ViewStateProvider>
  );
};

type theInstance = InstanceType<typeof types.ProjectContainer>;
const createProjectForm = (table: any): FieldGroup<theInstance>[] => {
  return [
    {
      legend: "Open Monitors",
      fields: [],
      collapsable: false,
      components: [
        {
          component: <DataTable<types.HistoryContainer> table={table} loading={false} />,
        },
      ],
    },
    {
      legend: "Data 1",
      colSpan: 6,
      fields: [
        { label: "fileName", type: "text", accessor: "filename" },
        { label: "nHistories", type: "int", accessor: "nOpenFiles" },
        { label: "historySize", type: "bytes", accessor: "historySize" },
        { label: "dirty", type: "boolean", accessor: "dirty" },
      ],
    },
    {
      legend: "Data 2",
      colSpan: 6,
      fields: [
        { label: "nMonitors", type: "int", accessor: "nMonitors" },
        { label: "nIndexes", type: "int", accessor: "nIndexes" },
        { label: "nManifests", type: "int", accessor: "nManifests" },
        { label: "nNames", type: "int", accessor: "nNames" },
        { label: "nAbis", type: "int", accessor: "nAbis" },
        { label: "nCaches", type: "int", accessor: "nCaches" },
      ],
    },
  ];
};
