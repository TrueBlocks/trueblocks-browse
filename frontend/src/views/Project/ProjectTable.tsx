import { Group } from "@mantine/core";
import { createColumnHelper } from "@tanstack/react-table";
import { CustomColumnDef, Formatter, ExploreButton, ExportButton, ViewButton } from "@components";
import { types } from "@gocode/models";

const columnHelper = createColumnHelper<types.HistoryContainer>();

export const tableColumns: CustomColumnDef<types.HistoryContainer, any>[] = [
  columnHelper.accessor("address", {
    header: () => "Address",
    cell: (info) => <Formatter type="address-editor" value={info.renderValue()} />,
    meta: { className: "wide cell" },
  }),
  columnHelper.accessor("nItems", {
    header: () => "Loaded",
    cell: (info) => {
      const { nItems, nTotal } = info.row.original;
      return <Formatter type="progress" value={nItems} value2={nTotal} />;
    },
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("balance", {
    header: () => "Balance",
    cell: (info) => <Formatter type="ether" value={info.renderValue()} />,
    meta: { className: "medium cell" },
  }),
  columnHelper.accessor("address", {
    header: () => " ",
    cell: (info) => {
      const { address } = info.row.original;
      return (
        <Group wrap={"nowrap"}>
          <ExploreButton size="sm" noText endpoint="address" value={info.renderValue()} />
          <ViewButton size="sm" noText value={info.renderValue()} />
          <ExportButton size="sm" noText value={info.renderValue()} />
        </Group>
      );
    },
    meta: { className: "wide cell" },
  }),
];
