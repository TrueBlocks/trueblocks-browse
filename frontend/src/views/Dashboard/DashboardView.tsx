// import { useEffect } from "react";
// import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
// import { View, FormTable, ViewForm } from "@components";
// import { useAppState, ViewStateProvider } from "@state";
import { useAppState } from "@state";
// import { GoToHistory } from "../../../wailsjs/go/app/App";
// import { Page, useNoops } from "../../hooks";
// import { withoutDelete, withDelete, DashboardFormDef } from ".";

export const DashboardView = () => {
  // const { modifyNoop } = useNoops();
  const { dashboard, fetchDashboard } = useAppState();

  // useEffect(() => {
  //   fetchDashboard();
  // }, []);

  return (
    <>
      <pre>{JSON.stringify(dashboard, null, 2)}</pre>
      <>----------------------------------</>
      {/* <div>{JSON.stringify(project)}</div> */}
    </>
  );
  // const { project, fetchProject } = useAppState();
  // const [filtered, setFiltered] = useState<types.HistoryContainer[]>([]);

  // const handleEnter = (page: Page) => {
  //   const record = project.items[page.getRecord()];
  //   GoToHistory(record).then(() => {});
  // };

  // const modColumns = dashboard.nItems < 2 ? withoutDelete : withDelete;

  // const table = useReactTable({
  //   data: project.items ?? [],
  //   columns: modColumns,
  //   getCoreRowModel: getCoreRowModel(),
  // });

  // const route = "";
  // const tabs = ["dashboard"];
  // const forms: ViewForm = {
  //   dashboard: <FormTable data={dashboard} groups={DashboardFormDef(dashboard, table)} />,
  // };
  // return (
  //   <ViewStateProvider
  //     route={route}
  //     nItems={dashboard.nItems}
  //     fetchFn={fetchProject}
  //     onEnter={handleEnter}
  //     modifyFn={modifyNoop}
  //   >
  //     <View tabs={tabs} forms={forms} />
  //   </ViewStateProvider>
  // );
};
