import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, DataTable, FieldGroup, ViewForm, AddButton } from "@components";
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

  const modColumns = project.nItems < 2 ? withoutDelete : withDelete;
  const table = useReactTable({
    data: project.items ?? [],
    columns: modColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "";
  const tabs = ["project"];
  const forms: ViewForm = {
    project: <FormTable data={project} groups={createProjectForm(table)} />,
  };
  return (
    <ViewStateProvider
      route={route}
      nItems={project.nItems}
      fetchFn={fetchProject}
      onEnter={handleEnter}
      modifyFn={ModifyProject}
    >
      <View tabs={tabs} forms={forms} />
    </ViewStateProvider>
  );
};

const createProjectForm = (table: any): FieldGroup<types.ProjectContainer>[] => {
  return [
    {
      label: "Data 1",
      colSpan: 6,
      fields: [
        { label: "fileName", type: "text", accessor: "filename" },
        { label: "nHistories", type: "int", accessor: "nItems" },
        { label: "historySize", type: "bytes", accessor: "historySize" },
        { label: "dirty", type: "boolean", accessor: "dirty" },
      ],
    },
    {
      label: "Data 2",
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
    {
      label: "Buttons",
      buttons: [<AddButton key={"add"} value={"https://trueblocks.io"} />],
    },
    {
      label: "Histories",
      fields: [],
      collapsable: false,
      components: [<DataTable<types.HistoryContainer> key={"dataTable"} table={table} loading={false} />],
    },
  ];
};
