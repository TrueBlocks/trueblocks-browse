import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { View, FormTable, ViewForm } from "@components";
import { GoToHistory, ModifyProject } from "@gocode/app/App";
import { base, types } from "@gocode/models";
import { Page } from "@hooks";
import { useAppState, ViewStateProvider } from "@state";
import { withoutDelete, withDelete, ProjectFormDef } from ".";

export const ProjectView = () => {
  const { project, fetchProject } = useAppState();
  // const [filtered, setFiltered] = useState<types.HistoryContainer[]>([]);

  const handleEnter = (page: Page) => {
    const address = project.items[page.getRecord()];
    GoToHistory(address).then(() => {});
  };

  const modColumns = project.nItems < 2 ? withoutDelete : withDelete;
  const projectContainers = project.items?.map(toHistoryContainer);

  const table = useReactTable({
    data: projectContainers ?? [],
    columns: modColumns,
    getCoreRowModel: getCoreRowModel(),
  });

  const route = "";
  const tabs = ["project"];
  const forms: ViewForm = {
    project: <FormTable data={project} groups={ProjectFormDef(table)} />,
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

const toHistoryContainer = (address: base.Address): types.HistoryContainer => {
  // Perform any additional operations needed here

  return {
    address: address as unknown as base.Address,
    balance: "",
    nErrors: 1,
    nLogs: 2,
    nTokens: 3,
    nTotal: 4,
    name: "5",
    items: [],
    nItems: 0,
    chain: "mainnet",
    lastUpdate: 0,
    convertValues: () => {},
  };
};
